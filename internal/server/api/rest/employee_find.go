package rest

import (
	"fmt"
	"net/http"
)

// Find
// @Tags GET
// @Summary Найти сотрудника по ФИО.
// @ID findGet
// @Produce plain
// @Param searchString query string false "ФИО"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employee/find{searchString} [GET]
func (s Handler) Find(w http.ResponseWriter, r *http.Request) {
	s.log.Info("Form Value : ", r.FormValue("searchString"))
	employee, err := s.employee.Find(r.FormValue("searchString"))
	if err != nil {
		s.log.Error("Employee find err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Employee find err : ", err)
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("%v", employee)))
	if err != nil {
		s.log.Error("Write err : ", err)
	}
	w.WriteHeader(http.StatusOK)
}
