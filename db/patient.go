package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/fusco2k/go-clinic-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var patientC *mongo.Collection
//initialize the session and the collection handler
func init() {
	client := NewSession()
	patientC = client.Database("clinic").Collection("patient")
}

//GetAll retrieves all patients on the patient collection
func GetAll() []model.Patient {
	//initialize a slice model to get data
	var Patients []model.Patient
	//gets the cursos with data
	cursor, err := patientC.Find(nil, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(nil)
	// loop throght the cursor decoding the data and append it to the slice model
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
	//returns the slice model
	return Patients
}

//GetOne returns the patient from requested ID
func GetOne(id primitive.ObjectID) model.Patient {
	//initialize the model to decoded the mongo data
	var Patient model.Patient
	//gets the patient related to id and decode to the pointe patient model
	err := patientC.FindOne(nil, bson.M{"_id": id}).Decode(&Patient)
	if err != nil {
		log.Println(err)
		return Patient
	}
	//returns the patient
	return Patient
}

//CreateOne calls GetOne to check if there is already a patient with that id, if not, it creates the patient
func CreateOne(patient model.Patient) {
	//creates a new patient on the collection
	_, err := patientC.InsertOne(nil, patient)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteOne deletes the patient with the related id
func DeleteOne(id primitive.ObjectID) {
	//delete the patient from the collection
	_, err := patientC.DeleteOne(nil, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateOne modifies the relate patient
func UpdateOne(patient []model.Patient) {
	//Replace the data on the collection
	_, err := patientC.ReplaceOne(nil, bson.M{"_id": patient[0].ID}, patient[1])
	if err != nil {
		log.Fatal(err)
	}
}
