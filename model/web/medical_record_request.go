package web

type MedicalRecordFindByPatientIDRequest struct {
	PatientID int `json:"patient_id" uri:"patient_id"`
}
