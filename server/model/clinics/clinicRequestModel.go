package clinics

type ClinicRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Specialty string `json:"specialty" binding:"required"`
}	