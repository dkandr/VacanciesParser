package api

import (
	"gitlab.com/dvkgroup/vacancies-parser/app/handler"
	"net/http"

	"github.com/go-chi/chi"
)

func swaggerRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", handler.SwaggerUI)

	return r
}
