package main

import (
	"github.com/fusco2k/go-clinic-crud/control"
	"github.com/fusco2k/go-clinic-crud/model"
	"log"
	"net/http"
)

func main() {
	pc := control.NewPatientController(getSession())

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/user/", func(writer http.ResponseWriter, request *http.Request) {
		pc.GetPatient(writer, request)
	})
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			pc.CreatePatient(writer, request)
		case "DELETE":
			pc.DeletePatient(writer, request)
		default:
			pc.Users(writer,request)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSession() map[int]model.Patient {
	// Connect to our local mongo
	storedb := map[int]model.Patient{}

	return storedb
}
