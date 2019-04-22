package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Patient basic structure using ObjectID from Mongo as ID
type Patient struct {
	ID        string             `json:"id" `
	ObjectID  primitive.ObjectID ` bson:"_id"`
	FName     string             `json:"fname"`
	LName     string             `json:"lname"`
	Email     string             `json:"email"`
	BloodType string             `json:"blood"`
	Cpf       int                `json:"cpf"`
	Birth     int                `json:"birth"`
	Phone     int                `json:"phone"`
	Mobile    int                `json:"mobile"`
}
