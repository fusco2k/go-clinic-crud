package control

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/fusco2k/go-clinic-crud/model"

	"github.com/fusco2k/go-clinic-crud/db"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("view/template/*.gohtml"))
}

//PatientController is a basic struct to get pointed to
type PatientController struct {
}

//NewPatientController gives a pointed PC
func NewPatientController() *PatientController {
	return &PatientController{}
}

//Patients retrieves all patients from the DB and shows as parsed template
func (pc PatientController) Patients(w http.ResponseWriter, r *http.Request) {
	patients := db.GetAll()
	err := tpl.ExecuteTemplate(w, "user.gohtml", patients)
	if err != nil {
		log.Fatalln(err)
	}
}

//GetPatient returns a patient from the requested ID
func (pc PatientController) GetPatient(w http.ResponseWriter, r *http.Request) {
	//parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//separate the request user id from url
	urlID := strings.SplitAfter(r.URL.String(), "/user/")
	stringID := urlID[1]
	id, _ := strconv.Atoi(stringID)

	patient, _ := db.GetOne(id)

	err = tpl.ExecuteTemplate(w, "userview.gohtml", patient)
	if err != nil {
		log.Fatalln(err)
	}

	//write the tcp 200 response with the JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

//CreatePatient creates a new patient from the JSON request
func (pc PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	//parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//initialize a patient model to populate with the json request
	patient := model.Patient{}

	//populate the patient using the json
	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Fatalln(err)
	}

	db.CreateOne(patient)

	//write the tcp 200 response with the JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

//DeletePatient deletes the request id related patient
func (pc PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {
	//parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//initialize a patient model to populate with the json request
	patient := model.Patient{}

	//populate the patient using the json
	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Fatalln(err)
	}

	db.DeleteOne(patient.ID)

	//write the tcp 200 response with the JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}
