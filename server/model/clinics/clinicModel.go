package clinics

import "github.com/google/uuid"

type Clinic struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Name      string `json:"name" bson:"name"`
	Address   string `json:"address" bson:"address"`
	Phone     string `json:"phone" bson:"phone"`
	Specialty string `json:"specialty" bson:"specialty"`
}

func NewClinic(req ClinicRequest) Clinic {
	return Clinic{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Specialty: req.Specialty,
	}
}
