package router

import (
	_ "employees-manager/docs"
	"employees-manager/internal/server/api/rest"
	"employees-manager/internal/server/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Route(s *rest.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.ContentTypeHandler)
	r.Route("/api", func(r chi.Router) {
		r.Route("/employee", func(r chi.Router) {
			r.Post("/hire", middleware.BasicAuth(s.Hire))
			r.Post("/fire", middleware.BasicAuth(s.Fire))
			r.Get("/get-number-of-vacation-days", middleware.BasicAuth(s.GetNumberOfVacationDays))
			r.Get("/find", middleware.BasicAuth(s.Find))
		})
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8088/swagger/doc.json"), //The url pointing to API definition
	))
	return r
}
