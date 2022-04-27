package user

import "time"

type User struct {
	Id             int
	Name           string
	Occopution     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdateAt       time.Time
}
