package clinics

import "github.com/google/uuid"

type Clinic struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Specialty string `json:"specialty"`
}

func NewClinic(req ClinicRequest) Clinic {
	return Clinic{
		ID:       uuid.New().String(),
		Name:     req.Name,
		Address:  req.Address,
		Phone:    req.Phone,
		Specialty: req.Specialty,
	}
}