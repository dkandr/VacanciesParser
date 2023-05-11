package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func vacancyRouter(c *handler.VacancyController) http.Handler {
	r := chi.NewRouter()

	r.Post("/search", c.Search)
	r.Post("/get", c.GetByID)
	r.Post("/delete", c.Delete)
	r.Get("/list", c.GetList)

	return r
}
