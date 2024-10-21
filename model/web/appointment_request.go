package web

type AppointmentCreateScheduleRequest struct {
	DoctorID        int    `json:"doctor_id"`
	PatientID       int    `json:"patient_id"`
	AppointmentDate string `json:"appointment_date"`
	AppointmentTime string `json:"appointment_time"`
	Status          string `json:"status"`
}
