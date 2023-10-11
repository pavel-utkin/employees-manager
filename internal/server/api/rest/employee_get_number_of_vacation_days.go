package rest

import (
	"fmt"
	"net/http"
	"strconv"
)

// GetNumberOfVacationDays
// @Tags GET
// @Summary Получить количество дней отпуска.
// @ID getNumberOfVacationDaysGet
// @Produce plain
// @Param employeeID query string false "идентификатор сотрудника"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/employee/get-number-of-vacation-days{employeeID} [GET]
func (s Handler) GetNumberOfVacationDays(w http.ResponseWriter, r *http.Request) {
	s.log.Info("Form Value : ", r.FormValue("employeeID"))
	requestedEmployeeID, err := strconv.ParseInt(r.FormValue("employeeID"), 10, 64)
	if err != nil {
		s.log.Error("Employee GetNumberOfVacationDays err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Write err : ", err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vacationDays, err := s.employee.GetNumberOfVacationDays(requestedEmployeeID)
	if err != nil {
		s.log.Error("Employee GetNumberOfVacationDays err : ", err)
		_, err = w.Write([]byte(fmt.Sprintf("%v", err)))
		if err != nil {
			s.log.Error("Write err : ", err)
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("%v", vacationDays)))
	if err != nil {
		s.log.Error("Write err : ", err)
	}
	w.WriteHeader(http.StatusOK)
}
