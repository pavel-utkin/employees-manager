package rest

import (
	"employees-manager/internal/server/models"
	"encoding/json"
	"net/http"
)

func (s Handler) Hire(w http.ResponseWriter, r *http.Request) {
	registeredEmployee := &models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&registeredEmployee)
	if err != nil {
		s.log.Println("json.NewDecoder err : ", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = s.employee.Hire(registeredEmployee)
	if err != nil {
		s.log.Println("Employee Fire err : ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
