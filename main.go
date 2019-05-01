package main

import (
	"log"
	"net/http"

	"github.com/fusco2k/go-clinic-crud/control"
)

func main() {
	pc := control.NewPatientController()

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/patients/", func(writer http.ResponseWriter, request *http.Request) {
		pc.GetPatient(writer, request)
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/patients", 307)
	})
	http.HandleFunc("/patients", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			pc.CreatePatient(writer, request)
		case "DELETE":
			pc.DeletePatient(writer, request)
		case "GET":
			pc.Patients(writer, request)
		case "PUT":
			pc.UpdatePatient(writer, request)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
