package model

type Patient struct {
	FName     string `json:"fname"`
	LName     string `json:"lname"`
	Email     string `json:"email"`
	BloodType string `json:"blood"`
	Id        int    `json:"id"`
	Cpf       int    `json:"cf"`
	Birth     int    `json:"birth"`
	Phone     int    `json:"phone"`
	Mobile    int    `json:"mobile"`
}
