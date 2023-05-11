package handler

import (
	"encoding/json"
	"fmt"
	"gitlab.com/dvkgroup/vacancies-parser/service/model"
	"net/http"
	"strconv"
)

type VacancyService interface {
	Search(string)
	GetByID(int64) (model.Vacancy, error)
	GetList() []model.Vacancy
	Delete(int64) error
}

type Parser interface {
	Parse()
}

type VacancyController struct {
	s VacancyService
}

func NewVacancyController(s VacancyService) *VacancyController {
	return &VacancyController{s: s}
}

func (c VacancyController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	c.s.Search(query)

	_, _ = fmt.Fprintf(w, "Ok.\n")
}

func (c VacancyController) GetByID(w http.ResponseWriter, r *http.Request) {
	vacancyID, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vacancy, err := c.s.GetByID(vacancyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeVacancyResponse(w, vacancy)
}

func (c VacancyController) GetList(w http.ResponseWriter, r *http.Request) {
	vacancies := c.s.GetList()
	writeVacanciesResponse(w, vacancies)
}

func (c VacancyController) Delete(w http.ResponseWriter, r *http.Request) {
	vacancyID, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.s.Delete(vacancyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func writeVacancyResponse(w http.ResponseWriter, vacancy model.Vacancy) {
	// выставляем заголовки, что отправляем json в utf8
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err := json.NewEncoder(w).Encode(vacancy)

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeVacanciesResponse(w http.ResponseWriter, vacancies []model.Vacancy) {
	// выставляем заголовки, что отправляем json в utf8
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err := json.NewEncoder(w).Encode(vacancies)

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
