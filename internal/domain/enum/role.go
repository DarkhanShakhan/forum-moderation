package enum

type UserRole int

const (
	RoleGuest UserRole = iota
	RoleUser
	RoleModerator
	RoleAdmin
)
