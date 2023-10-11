package rest

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s Handler) Fire(w http.ResponseWriter, r *http.Request) {
	s.log.Info("Form Value : ", r.FormValue("employeeID"))
	requestedEmployeeID, err := strconv.ParseInt(r.FormValue("employeeID"), 10, 64)
	if err != nil {
		s.log.Error("Employee Fire err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Employee Fire err : ", err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.employee.Fire(requestedEmployeeID)
	if err != nil {
		s.log.Error("Employee Fire err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Employee Fire err : ", err)
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("%v", requestedEmployeeID)))
	if err != nil {
		s.log.Error("Write err : ", err)
	}
	w.WriteHeader(http.StatusOK)
}
