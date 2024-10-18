package web

type PatientRegisterRequest struct {
	Name        string              `json:"name"`
	DateOfBirth string              `json:"date_of_birth"`
	Address     string              `json:"address"`
	Phone       string              `json:"phone"`
	User        UserRegisterRequest `json:"user"`
}
type PatientFindByIdRequest struct {
	Id int `json:"id" uri:"id"`
}
