package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gitlab.com/dvkgroup/vacancies-parser/app/model"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/tebeka/selenium"
)

type VacancyParser struct {
	driver selenium.WebDriver
}

func NewVacancyParser() *VacancyParser {
	return &VacancyParser{}
}

func (v VacancyParser) Parse(query string) ([]model.Vacancy, error) {
	// прописываем конфигурацию для драйвера
	caps := selenium.Capabilities{
		"browserName": "firefox",
	}

	// прописываем адрес нашего драйвера
	urlPrefix := "http://selenium:4444/wd/hub"

	// немного костылей чтобы драйвер не падал
	const maxTries = 5
	//var v.driver selenium.WebDriver
	var err error
	i := 1
	for i < maxTries {
		v.driver, err = selenium.NewRemote(caps, urlPrefix)
		if err != nil {
			log.Println(err)
			i++
			continue
		}
		break
	}

	defer v.driver.Quit()

	// start parse
	count, err := v.getVacancyCount(query)
	if err != nil {
		return nil, err
	}

	links, err := v.getVacancies(count, query)
	if err != nil {
		return nil, err
	}

	res := make([]model.Vacancy, 0, len(links))

	var m sync.Mutex
	var wg sync.WaitGroup

	limit := make(chan struct{}, 10)

	for _, l := range links {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()

			v, err := v.getVacancy(link, limit)
			if err != nil {
				log.Println("Parse:", err)
				return
			}

			m.Lock()
			res = append(res, v)
			m.Unlock()
		}(l)
	}

	wg.Wait()

	return res, nil
}

func (v VacancyParser) getVacancyCount(query string) (int, error) {
	// сразу обращаемся к странице с поиском вакансии по запросу
	page := 1 // номер страницы

	_ = v.driver.Get(fmt.Sprintf("https://career.habr.com/vacancies?page=%d&q=%s&type=all", page, query))

	elem, err := v.driver.FindElement(selenium.ByCSSSelector, ".search-total")
	if err != nil {
		return 0, err
	}

	vacancyCountRaw, err := elem.Text()
	if err != nil {
		return 0, err
	}

	vacancyCountText := strings.Fields(vacancyCountRaw)

	count, err := strconv.Atoi(vacancyCountText[1])
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (v VacancyParser) getVacancies(count int, query string) ([]string, error) {
	const (
		VacanciesPerPage = 25
		HabrCareerLink   = "https://career.habr.com"
	)

	links := make([]string, 0, count)

	for i := 1; i <= count/VacanciesPerPage+1; i++ {
		_ = v.driver.Get(fmt.Sprintf("https://career.habr.com/vacancies?page=%d&q=%s&type=all", i, query))

		elems, err := v.driver.FindElements(selenium.ByCSSSelector, ".vacancy-card__title-link")
		if err != nil {
			return nil, err
		}

		for n := range elems {
			link, err := elems[n].GetAttribute("href")
			if err != nil {
				continue
			}

			links = append(links, HabrCareerLink+link)
		}
	}

	return links, nil
}

func (v VacancyParser) getVacancy(link string, limit chan struct{}) (model.Vacancy, error) {
	limit <- struct{}{}

	defer func() {
		<-limit
	}()

	resp, err := http.Get(link)
	if err != nil {
		return model.Vacancy{}, err
	}

	if resp.Status != "200 OK" {
		return model.Vacancy{}, errors.New(link + ":" + resp.Status)
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil && doc != nil {
		return model.Vacancy{}, err
	}

	dd := doc.Find("script[type=\"application/ld+json\"]")
	if dd == nil {
		return model.Vacancy{}, errors.New("habr vacancy nodes not found")
	}

	ss := dd.First().Text()

	var vacancy model.Vacancy
	err = json.Unmarshal([]byte(ss), &vacancy)
	if err != nil {
		return model.Vacancy{}, err
	}

	return vacancy, nil
}
