package control

import (
	"encoding/json"
	"fmt"
	"github.com/fusco2k/go-clinic-crud/model"
	"strings"
	"log"
	"net/http"
)

type PatientController struct {
}

func NewPatientController() *PatientController {
	return &PatientController{}
}

func (pc PatientController) GetPatient(w http.ResponseWriter, r *http.Request) {
	//Generate a fake patient for response
	p := model.Patient{
		Id:    001,
		FName: "James",
		LName: "Bond",
	}
	//Parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	//separate the request user id from url
	id := strings.SplitAfter(r.URL.String(), "/user/")
	fmt.Println(id[1])
	//marshal the patient to generate the response
	js, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	//write the tcp 200 response with the JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	_, err = fmt.Fprintf(w, "%s\n", js)
	if err != nil {
		log.Println(err)
	}
}

func (pc PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	err:=r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(r.Body)
}

func (pc PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {

}

func (pc PatientController) ServeHttp(w http.ResponseWriter, r *http.Request) {

}
