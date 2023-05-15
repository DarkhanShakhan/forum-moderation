package entity

import "github.com/DarkhanShakhan/forum-moderation/internal/domain/enum"

type User struct {
	ID       string
	Username string
	Email    string
	Role     enum.UserRole
}
