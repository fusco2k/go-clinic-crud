package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Patient basic structure using ObjectID from Mongo as ID
type Patient struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName  string             `json:"fname,omitempty" bson:"fname,omitempty"`
	LName  string             `json:"lname,omitempty" bson:"lname,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Cpf    int                `json:"cpf,omitempty" bson:"cpf,omitempty"`
	Birth  int                `json:"birth,omitempty" bson:"birth,omitempty"`
	Phone  int                `json:"phone,omitempty" bson:"phone,omitempty"`
	Mobile int                `json:"mobile,omitempty" bson:"mobile,omitempty"`
}
