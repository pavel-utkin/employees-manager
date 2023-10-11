package models

type Employee struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Gender       string  `json:"gender"`
	Phone        string  `json:"phone,omitempty"`
	Email        string  `json:"email,omitempty"`
	Address      string  `json:"address,omitempty"`
	VacationDays float64 `json:"vacationDays,omitempty"`
}
