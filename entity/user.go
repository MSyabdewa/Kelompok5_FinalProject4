package entity

import "time"

type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
