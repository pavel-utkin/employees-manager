package rest

import (
	"employees-manager/internal/server/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime"
	"net/http"
)

func (s Handler) Hire(w http.ResponseWriter, r *http.Request) {
	registeredEmployee := &models.Employee{}

	contentType := r.Header.Get("Content-Type")
	mt, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
		return
	}

	switch mt {
	case "application/xml":
		byteValue, _ := io.ReadAll(r.Body)
		err := xml.Unmarshal(byteValue, registeredEmployee)
		if err != nil {
			s.log.Error("json.NewDecoder err : ", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&registeredEmployee)
		if err != nil {
			s.log.Error("json.NewDecoder err : ", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
	}
	err = s.employee.Hire(registeredEmployee)
	if err != nil {
		s.log.Error("Employee Fire err : ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintln("record was inserted successfully")))
	if err != nil {
		s.log.Error("Write err : ", err)
	}
	w.WriteHeader(http.StatusOK)
}
