package clinics

import "github.com/google/uuid"

type Clinic struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"not null" json:"name"`
	Address string `gorm:"not null" json:"address"`
	Phone   string `gorm:"not null" json:"phone"`
	Specialty string `gorm:"not null" json:"specialty"`
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