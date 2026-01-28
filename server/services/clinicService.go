package services

import (
	"my-go-server/model/clinics"
	"my-go-server/repositories"


)

type ClinicService struct {
	repo repositories.Repository[clinics.Clinic]
}

func NewClinicService(repo repositories.Repository[clinics.Clinic]) *ClinicService {
	return &ClinicService{repo: repo}
}

func (s *ClinicService) InsertClinic(clinicReq clinics.ClinicRequest) (clinics.Clinic, error) {
	clinicModel := clinics.NewClinic(clinicReq)
	return s.repo.Insert(clinicModel)
}

func (s *ClinicService) GetClinicByID(id string) (clinics.Clinic, error) {
	return s.repo.GetByID(id)
}

func (s *ClinicService) UpdateClinic(clinic clinics.Clinic) (clinics.Clinic, error) {
	return s.repo.Update(clinic)
}

func (s *ClinicService) DeleteClinic(id string) error {
	return s.repo.Delete(id)
}
