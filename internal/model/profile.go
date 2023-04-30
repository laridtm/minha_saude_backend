package model

type Profile struct {
	FullName       string `json:"fullName" bson:"fullName"`
	Gender         string `json:"gender" bson:"gender"`
	BirthDate      string `json:"birthDate" bson:"birthDate"`
	CPF            string `json:"cpf" bson:"cpf"`
	PhoneNumber    string `json:"phoneNumber" bson:"phoneNumber"`
	Address        string `json:"address" bson:"address"`
	MaritalStatus  string `json:"maritalStatus" bson:"maritalStatus"`
	BloodType      string `json:"bloodType" bson:"bloodType"`
	EmergencyPhone string `json:"emergencyPhone" bson:"emergencyPhone"`
	Allergies      string `json:"allergies" bson:"allergies"`
}
