package rest

import (
	"log"
	"net/http"
)

func (s Handler) Find(w http.ResponseWriter, r *http.Request) {
	log.Println("Method find")
	log.Println(r.FormValue("Name"))
	_, err := s.employee.Find(r.FormValue("Name"))
	if err != nil {
		//logrus.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
