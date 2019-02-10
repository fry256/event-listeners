package model

import "time"

type User struct {
	ID        string
	Email     string
	Password  string
	Wishes    []*Wish
	CreatedOn time.Time
	UpdatedOn time.Time
}
