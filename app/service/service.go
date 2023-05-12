package service

import (
	"gitlab.com/dvkgroup/vacancies-parser/app/model"
	"log"
)

type VacancyParser interface {
	Parse(string) ([]model.Vacancy, error)
}

type VacancyRepository interface {
	Create(model.Vacancy) (model.Vacancy, error)
	GetByID(int64) (model.Vacancy, error)
	GetList() []model.Vacancy
	Delete(int64) error
}

type VacancyService struct {
	r VacancyRepository
	p VacancyParser
}

func NewVacancyService(r VacancyRepository, p VacancyParser) *VacancyService {
	return &VacancyService{r: r, p: p}
}

func (s *VacancyService) Search(query string) {
	//go func() {
	log.Println("parser start (query :", query, ")")
	vacancies, err := s.p.Parse(query)
	if err != nil {
		log.Println(err)
		return
	}

	for i := range vacancies {
		_, err := s.r.Create(vacancies[i])
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("parser.. ok!")
	//}()
}

func (s *VacancyService) Create(vacancy model.Vacancy) (model.Vacancy, error) {
	return s.r.Create(vacancy)
}

func (s *VacancyService) GetByID(id int64) (model.Vacancy, error) {
	return s.r.GetByID(id)
}

func (s *VacancyService) GetList() []model.Vacancy {
	return s.r.GetList()
}

func (s *VacancyService) Delete(id int64) error {
	return s.r.Delete(id)
}
