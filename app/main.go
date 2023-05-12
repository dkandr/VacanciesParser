package main

import (
	"gitlab.com/dvkgroup/vacancies-parser/app/cmd/api"
	"gitlab.com/dvkgroup/vacancies-parser/app/handler"
	"gitlab.com/dvkgroup/vacancies-parser/app/parser"
	"gitlab.com/dvkgroup/vacancies-parser/app/repository"
	"gitlab.com/dvkgroup/vacancies-parser/app/service"
)

func main() {
	p := parser.NewVacancyParser()

	r := repository.NewVacancyRepository()
	s := service.NewVacancyService(r, p)
	c := handler.NewVacancyController(s)

	api.Run(c)
}
