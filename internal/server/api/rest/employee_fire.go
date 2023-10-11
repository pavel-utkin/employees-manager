package rest

import (
	"fmt"
	"net/http"
	"strconv"
)

// Fire
// @Tags POST
// @Summary Уволить сотрудника.
// @ID firePost
// @Produce plain
// @Param employeeID query string false "идентификатор сотрудника"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employee/fire{employeeID} [POST]

func (s Handler) Fire(w http.ResponseWriter, r *http.Request) {
	s.log.Info("Form Value : ", r.FormValue("employeeID"))
	requestedEmployeeID, err := strconv.ParseInt(r.FormValue("employeeID"), 10, 64)
	if err != nil {
		s.log.Error("Employee fire err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Employee fire err : ", err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.employee.Fire(requestedEmployeeID)
	if err != nil {
		s.log.Error("Employee fire err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Employee fire err : ", err)
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
