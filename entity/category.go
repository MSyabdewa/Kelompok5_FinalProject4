package entity

import "time"

type Category struct {
	ID                int
	Type              string
	SoldProductAmount int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
