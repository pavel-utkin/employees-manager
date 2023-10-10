package employee

import (
	"employees-manager/internal/server/database"
	"employees-manager/internal/server/models"
)

type Employee struct {
	db *database.DB
}

func New(db *database.DB) *Employee {
	return &Employee{
		db: db,
	}
}

/*
type Constructor interface {
	Hire(user *model.UserRequest) (*model.User, error)
	GetNumberOfVacationDays(userRequest *model.UserRequest) (*model.User, error)
	Fire(username string) (bool, error)
	Find(username string) (bool, error)
}*/

func (e Employee) Hire() (*models.Employee, error) {
	registeredEmployee := &models.Employee{}
	if err := e.db.Pool.QueryRow(
		"INSERT INTO Employee (Name) VALUES ($1) RETURNING ID",
		registeredEmployee.Name,
	).Scan(&registeredEmployee.ID); err != nil {
		return nil, err
	}
	return registeredEmployee, nil
}

func (e Employee) GetNumberOfVacationDays() {

}

func (e Employee) Fire() {

}

func (e Employee) Find(name string) (int64, error) {
	var id int64
	if err := e.db.Pool.QueryRow("SELECT ID FROM Employee where Name = ?",
		name,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
