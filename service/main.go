package main

import (
	"gitlab.com/dvkgroup/vacancies-parser/service/cmd/api"
	"gitlab.com/dvkgroup/vacancies-parser/service/handler"
	"gitlab.com/dvkgroup/vacancies-parser/service/parser"
	"gitlab.com/dvkgroup/vacancies-parser/service/repository"
	"gitlab.com/dvkgroup/vacancies-parser/service/service"
)

func main() {
	p := parser.NewVacancyParser()

	r := repository.NewVacancyRepository()
	s := service.NewVacancyService(r, p)
	c := handler.NewVacancyController(s)

	api.Run(c)
}
