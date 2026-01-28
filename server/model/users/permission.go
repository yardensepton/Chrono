package users

type Permission string

const (
	PermissionManageAvailability Permission = "manage_availability"
	PermissionWriteUsers         Permission = "write_users"
	PermissionManageAppointments Permission = "manage_appointments"
	PermissionViewPatients       Permission = "view_patients"
)