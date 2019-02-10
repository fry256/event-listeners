package model

import "time"

type Wish struct {
	ID        string
	UserID    string
	Content   string
	CreatedOn time.Time
	UpdatedOn time.Time
}
