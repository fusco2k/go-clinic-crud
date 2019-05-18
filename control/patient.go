package control

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/fusco2k/go-clinic-crud/model"

	"github.com/fusco2k/go-clinic-crud/db"
)

//PatientController is a basic struct to get pointed to
type PatientController struct{}

//NewPatientController gives a pointed PC
func NewPatientController() *PatientController {
	return &PatientController{}
}

//Patients retrieves all patients from the DB and shows as parsed template
func (pc PatientController) Patients(w http.ResponseWriter, r *http.Request) {
	patients := db.GetAll()
	json.NewEncoder(w).Encode(patients)
}

//GetPatient returns a patient from the requested ID
func (pc PatientController) GetPatient(w http.ResponseWriter, r *http.Request) {
	//parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	//separate the request user id from url
	urlID := strings.SplitAfter(r.URL.String(), "/patients/")
	//convert the id to ObjectID
	id, _ := primitive.ObjectIDFromHex(urlID[1])
	//get the patient relate to the url id
	patient := db.GetOne(id)
	//respond with the patient data by json
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
	//populate the patient decoding the json
	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Fatalln(err)
	}
	//create the patient using the populated model
	db.CreateOne(patient)
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
	//populate the patient decoding the json
	err = json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Fatalln(err)
	}
	//delete the related patient
	db.DeleteOne(patient.ID)
}

//UpdatePatient updates the requested patient with data related
func (pc PatientController) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	//parse the request form
	r.ParseForm()
	//initialize the decode model
	patient := []model.Patient{}
	//decode the json
	json.NewDecoder(r.Body).Decode(&patient)
	//update the data on the collection
	db.UpdateOne(patient)
}
