package main

import (
	"context"
	"employees-manager/internal/server/api"
	"employees-manager/internal/server/api/rest"
	"employees-manager/internal/server/config"
	"employees-manager/internal/server/database"
	"employees-manager/internal/server/repository/employee"
	"employees-manager/internal/server/router"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {

	logger := logrus.New()
	servConfig := config.NewConfig(logger)

	db, err := database.New(servConfig, logger)
	if err != nil {
		logger.Fatal(err)
	}

	employeeRepository := employee.New(db)

	handlerRest := rest.NewHandler(db, servConfig, employeeRepository, logger)
	routerService := router.Route(handlerRest)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go api.StartRESTService(rs, servConfig, logger)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
