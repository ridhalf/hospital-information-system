package policy

import (
	"hospital-information-system/model/constants"
	"hospital-information-system/model/domain"
)

func UserPolicy(user domain.User, patient domain.Patient) bool {
	if user.ID == patient.UserID || user.Role == constants.ADMIN || user.Role == constants.PATIENT || user.Role == constants.NURSE {
		return true
	}
	return false
}
