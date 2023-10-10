package api

import (
	"employees-manager/internal/server/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

// StartRESTService - starts the REST gophkeeper server
func StartRESTService(r *chi.Mux, config *config.Config, log *logrus.Logger) {
	log.Infof("Starting REST server %v\n", config.AddressREST)
	if lis := http.ListenAndServe(config.AddressREST, r); lis != nil {
		log.Fatalf("failed to listen: %v", lis)
	}
}
