package models

type Employee struct {
	ID           int64   `xml:"id" json:"id"`
	Name         string  `xml:"name" json:"name"`
	Gender       string  `xml:"gender" json:"gender"`
	Age          int64   `xml:"age" json:"age"`
	Phone        string  `xml:"phone,omitempty" json:"phone,omitempty"`
	Email        string  `xml:"email,omitempty" json:"email,omitempty"`
	Address      string  `xml:"address,omitempty" json:"address,omitempty"`
	VacationDays float64 `xml:"vacationDays,omitempty" json:"vacationDays,omitempty"`
}
