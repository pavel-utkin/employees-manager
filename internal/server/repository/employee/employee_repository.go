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

func (e Employee) Hire(newEmployee *models.Employee) (*models.Employee, error) {
	registeredEmployee := &models.Employee{}
	if err := e.db.Pool.QueryRow(
		"INSERT INTO Employee (Name,Gender,Phone,Email,Address,VacationDays) VALUES (?,?,?,?,?,?)",
		newEmployee.Name,
		newEmployee.Gender,
		newEmployee.Phone,
		newEmployee.Email,
		newEmployee.Address,
		newEmployee.VacationDays,
	).Scan(&registeredEmployee.ID); err != nil {
		return nil, err
	}
	return registeredEmployee, nil
}

func (e Employee) GetNumberOfVacationDays(employeeId int64) (float64, error) {
	var vacationDays float64
	if err := e.db.Pool.QueryRow("SELECT VacationDays FROM Employee where ID = ?",
		employeeId,
	).Scan(&vacationDays); err != nil {
		return vacationDays, err
	}
	return vacationDays, nil
}

func (e Employee) Fire(employeeId int64) error {
	err := e.db.Pool.QueryRow("DELETE FROM Employee where ID = ?",
		employeeId,
	).Err()
	return err
}

func (e Employee) Find(pattern string) (int64, error) {
	var id int64
	if err := e.db.Pool.QueryRow("SELECT ID FROM Employee where Name LIKE ?",
		pattern,
	).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
