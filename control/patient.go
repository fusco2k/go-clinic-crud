package control

import (
	"encoding/json"
	"fmt"
	"github.com/fusco2k/go-clinic-crud/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var tpl *template.Template

//create the type controller to use outside the package
type PatientController struct {
	storeDB map[int]model.Patient
}

//function to call a new controller
func NewPatientController(db map[int]model.Patient) *PatientController {
	tpl = template.Must(template.New("").ParseGlob("view/template/*.gohtml"))
	return &PatientController{db}
}

func (pc PatientController) Users(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "user.gohtml", pc.storeDB)
	if err != nil {
		log.Fatalln(err)
	}
}

//function to call the patient by id
func (pc PatientController) GetPatient(w http.ResponseWriter, r *http.Request) {
	//parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//separate the request user id from url
	urlId := strings.SplitAfter(r.URL.String(), "/user/")
	stringId := urlId[1]

	//convert the id to int
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.Fatalln(err)
	}

	//check if there is a existent patient with the id
	if pc.storeDB[id].Id != id {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	//fetch user
	u := pc.storeDB[id]

	err = tpl.ExecuteTemplate(w, "userview.gohtml", u)
	if err != nil {
		log.Fatalln(err)
	}

	//marshal provided interface into JSON structure
	js, _ := json.Marshal(u)

	//write the tcp 200 response with the JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	_, err = fmt.Fprintf(w, "%s\n", js)
	if err != nil {
		log.Println(err)
	}
}

func (pc PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {

	//parse the request to get data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//initialize a patient model to populate with the json request
	u := model.Patient{}

	//populate the patient using the json
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatalln(err)
	}

	//check if there is a existent patient with the same id, if true, throws an error
	if pc.storeDB[u.Id].Id == u.Id {
		_, err := fmt.Fprint(w, "user already exists")
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	//store in the map db
	pc.storeDB[u.Id] = u

	//marshal provided interface into JSON structure
	js, _ := json.Marshal(u)

	//write the tcp 203 and redirect to the user page
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	_, err = fmt.Fprintf(w, "%s\n", js)
	if err != nil {
		log.Fatal(err)
	}
}

func (pc PatientController) DeletePatient(w http.ResponseWriter, r *http.Request) {

	//initialize a patient model to populate with the json request
	u := model.Patient{}

	//populate the patient using the json
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u.Id)

	//check if there is a existent patient with the id
	if pc.storeDB[u.Id].Id != u.Id {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//delete the patient from the map db
	delete(pc.storeDB, u.Id)

	//write the tcp 200 response with the message user deleted
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	_, err = fmt.Fprintln(w, "patient: "+strconv.Itoa(u.Id)+" deleted")
	if err != nil {
		log.Fatal(err)
	}
}
