package rest

import (
	"employees-manager/internal/server/config"
	"employees-manager/internal/server/database"
	"employees-manager/internal/server/repository/employee"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	database *database.DB
	config   *config.Config
	employee *employee.Employee
	log      *logrus.Logger
}

// NewHandler - creates a new server instance
func NewHandler(db *database.DB, config *config.Config, employeeRepository *employee.Employee,
	log *logrus.Logger) *Handler {
	return &Handler{database: db, config: config, employee: employeeRepository, log: log}
}
