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

type Content struct {
	Id               int
	Userid           string
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	GoalAmount       string
	CurrentAmount    string
	BeckerCount      string
	Slug             string
	CreatedAt        time.Time
	UpdateAt         time.Time
}
