package control

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

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

	patient := db.GetOne(stringID)

	err = tpl.ExecuteTemplate(w, "userview.gohtml", patient)
	if err != nil {
		log.Fatalln(err)
	}

	//write the tcp 200 response with the JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

// func (pc PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {

// 	//parse the request to get data
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//initialize a patient model to populate with the json request
// 	u := model.Patient{}

// 	//populate the patient using the json
// 	err = json.NewDecoder(r.Body).Decode(&u)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	//check if there is a existent patient with the same id, if true, throws an error
// 	if pc.storeDB[u.Id].Id == u.Id {
// 		_, err := fmt.Fprint(w, "user already exists")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return
// 	}

// 	//store in the map db
// 	pc.storeDB[u.Id] = u

// 	//write the tcp 200 response with the JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(u)
// }

// func (pc PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {

// 	//initialize a patient model to populate with the json request
// 	u := model.Patient{}

// 	//populate the patient using the json
// 	err := json.NewDecoder(r.Body).Decode(&u)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(u.Id)

// 	//check if there is a existent patient with the id
// 	if pc.storeDB[u.Id].Id != u.Id {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	//delete the patient from the map db
// 	delete(pc.storeDB, u.Id)

// 	//write the tcp 200 response with the message user deleted
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK) // 200
// 	_, err = fmt.Fprintln(w, "patient: "+strconv.Itoa(u.Id)+" deleted")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
