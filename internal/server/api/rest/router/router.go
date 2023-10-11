package router

import (
	"employees-manager/internal/server/api/rest"
	"github.com/go-chi/chi/v5"
)

func Route(s *rest.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/employee", func(r chi.Router) {
			r.Post("/hire", s.Hire)
			r.Post("/fire", s.Fire)
			r.Get("/get-number-of-vacation-days", s.GetNumberOfVacationDays)
			r.Get("/find", s.Find)
		})
	})
	return r
}
