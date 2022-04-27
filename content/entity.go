package content

import "time"

type Content struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BeckerCount      int
	GoalAmount       string
	CurrentAmount    string
	Slug             string
	CreatedAt        time.Time
	UpdateAt         time.Time
}

type ContentImage struct {
	Id        int
	ContentId int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdateAt  time.Time
}
