package rest

import (
	"net/http"
	"strconv"
)

func (s Handler) Fire(w http.ResponseWriter, r *http.Request) {
	s.log.Println("Form Value : ", r.FormValue("employeeID"))
	requestedEmployeeID, err := strconv.ParseInt(r.FormValue("employeeID"), 10, 64)
	if err != nil {
		s.log.Println("Employee Fire err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.employee.Fire(requestedEmployeeID)
	if err != nil {
		s.log.Println("Employee Fire err : ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
