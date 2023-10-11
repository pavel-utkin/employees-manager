package rest

import (
	"net/http"
)

func (s Handler) Find(w http.ResponseWriter, r *http.Request) {
	s.log.Println("Form Value : ", r.FormValue("searchString"))
	_, err := s.employee.Find(r.FormValue("searchString"))
	if err != nil {
		s.log.Println("Employee find err : ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
