package main

import (
	"github.com/fusco2k/go-clinic-crud/control"
	"log"
	"net/http"
)

func main() {
	pc := control.NewPatientController()

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/user/", func(writer http.ResponseWriter, request *http.Request) {
		// pc.GetPatient(writer, request)
	})
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		// case "POST":
		// 	pc.CreatePatient(writer, request)
		// case "DELETE":
		// 	pc.DeletePatient(writer, request)
		default:
			pc.Patients(writer, request)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
