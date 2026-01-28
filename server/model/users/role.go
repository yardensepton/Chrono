package users

type Role string

const (
	RoleTherapist Role = "therapist"
	RoleAssistant Role = "assistant"
	RolePatient   Role = "patient"
)