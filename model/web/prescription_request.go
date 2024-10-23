package web

type PrescriptionCreateRequest struct {
	DoctorID   int    `json:"doctor_id"`
	PatientID  int    `json:"patient_id"`
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
}
type PrescriptionUpdateRequest struct {
	ID         int    `json:"id"`
	DoctorID   int    `json:"doctor_id"`
	PatientID  int    `json:"patient_id"`
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
}
