package model

type Patient struct {
	Id        int    `json:"id"`
	FName     string `json:"fname"`
	LName     string `json:"lname"`
	Email     string `json:"email"`
	BloodType string `json:"blood"`
	Cpf       int    `json:"cpf"`
	Birth     int    `json:"birth"`
	Phone     int    `json:"phone"`
	Mobile    int    `json:"mobile"`
}
