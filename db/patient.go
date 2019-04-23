package db

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

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
		patient.ID = patient.ObjectID.Hex()
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
func GetOne(id string) model.Patient {
	var Patient model.Patient

	object, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	err = patientDB.FindOne(nil, bson.M{"_id": object}).Decode(&Patient)
	if err != nil {
		log.Fatal(err)
	}
	Patient.ID = Patient.ObjectID.Hex()

	return Patient
}
