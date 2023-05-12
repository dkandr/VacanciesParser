package repository

import (
	"errors"
	"gitlab.com/dvkgroup/vacancies-parser/app/model"
	"strconv"
)

type VacancyStorage struct {
	vacancies map[int64]model.Vacancy
}

func NewVacancyRepository() *VacancyStorage {
	return &VacancyStorage{vacancies: make(map[int64]model.Vacancy)}
}

func (r VacancyStorage) Create(vacancy model.Vacancy) (model.Vacancy, error) {
	id, err := strconv.ParseInt(vacancy.Identifier.Value, 10, 64) // получаем ID
	if err != nil {
		return model.Vacancy{}, err
	}

	// если id нет, то добавляем
	if _, ok := r.vacancies[id]; !ok {
		r.vacancies[id] = vacancy
		return vacancy, nil
	}

	// id уже существует
	return r.vacancies[id], errors.New("vacancy id exists in db")

}

func (r VacancyStorage) GetByID(id int64) (model.Vacancy, error) {
	// если id есть
	if v, ok := r.vacancies[id]; ok {
		return v, nil
	}

	// id если нет
	return model.Vacancy{}, errors.New("id not found")

}

func (r VacancyStorage) GetList() []model.Vacancy {
	res := make([]model.Vacancy, 0, len(r.vacancies))

	// map to slice
	for i := range r.vacancies {
		res = append(res, r.vacancies[i])
	}

	return res
}

func (r VacancyStorage) Delete(id int64) error {
	if _, ok := r.vacancies[id]; ok {
		delete(r.vacancies, id)
		return nil
	}

	return errors.New("id not found")
}
