package web

import "hospital-information-system/model/domain"

type PrescriptionCreateResponse struct {
	Id         int    `json:"id"`
	DoctorID   int    `json:"doctor_id"`
	PatientID  int    `json:"patient_id"`
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
}

func ToPrescriptionCreateResponse(prescription domain.Prescription) PrescriptionCreateResponse {
	return PrescriptionCreateResponse{
		Id:         prescription.ID,
		DoctorID:   prescription.DoctorID,
		PatientID:  prescription.PatientID,
		Medication: prescription.Medication,
		Dosage:     prescription.Dosage,
	}
}
