package db

import (
	"fmt"
	"log"

	"github.com/fusco2k/go-clinic-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var patientDB *mongo.Collection

func init() {
	clientDB := NewSession()
	patientDB = clientDB.Database("clinic").Collection("patient")
}

//GetAll retrieves all patients on the patient collection
func GetAll() []model.Patient {

	var Patients []model.Patient

	cursor, err := patientDB.Find(nil, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(nil)

	for cursor.Next(nil) {
		var patient model.Patient
		err = cursor.Decode(&patient)
		if err != nil {
			log.Fatal(err)
		}
		Patients = append(Patients, patient)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(Patients)

	return Patients
}

//GetOne returns the patient from requested ID
func GetOne(id int) (model.Patient, bool){
	var Patient model.Patient

	err := patientDB.FindOne(nil, bson.M{"id": id}).Decode(&Patient)
	if err != nil {
		log.Println(err)
		return Patient, false
	}

	return Patient, true
}

//CreateOne calls GetOne to check if there is already a patient with that id, if not, it creates the patient
func CreateOne(patient model.Patient) {
	_, b := GetOne(patient.ID)
	if b {
		return 
	}
	result, err := patientDB.InsertOne(nil, patient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted patient ID: ", result.InsertedID)
}

//DeleteOne deletes the patient with the related id
func DeleteOne(id int) {
	_, b := GetOne(id)
	if !b {
		return 
	}
	result, err := patientDB.DeleteOne(nil, bson.M{"id":id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Patients deleted: ", result.DeletedCount)
}
