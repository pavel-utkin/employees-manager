package rest

import (
	"fmt"
	"net/http"
)

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
