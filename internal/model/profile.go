package model

type Profile struct {
	FullName       string `json:"fullName"`
	BirthDate      string `json:"birthDate"`
	CPF            string `json:"cpf"`
	PhoneNumber    string `json:"phoneNumber"`
	Address        string `json:"address"`
	MaritalStatus  string `json:"maritalStatus"`
	BloodType      string `json:"bloodType"`
	EmergencyPhone string `json:"emergencyPhone"`
	Allergies      string `json:"allergies"`
}
