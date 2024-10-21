package web

import "hospital-information-system/model/domain"

type MedicalRecordResponse struct {
	PatientID  int                  `json:"patient_id"`
	Diagnosis  string               `json:"diagnosis"`
	Treatment  string               `json:"treatment"`
	Medication string               `json:"medication"`
	Patient    MedicalRecordPatient `json:"patient"`
}
type MedicalRecordPatient struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

func ToMedicalRecordResponses(medicalRecord []domain.MedicalRecord) []MedicalRecordResponse {
	var medicalRecordResponses []MedicalRecordResponse
	for _, medical := range medicalRecord {
		medicalRecordResponse := MedicalRecordResponse{
			PatientID:  medical.PatientID,
			Diagnosis:  medical.Diagnosis,
			Treatment:  medical.Treatment,
			Medication: medical.Medication,
			Patient: MedicalRecordPatient{
				ID:     medical.Patient.ID,
				Name:   medical.Patient.Name,
				UserID: medical.Patient.UserID,
			},
		}
		medicalRecordResponses = append(medicalRecordResponses, medicalRecordResponse)
	}
	return medicalRecordResponses
}
