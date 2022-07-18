package entity

import "time"

type LifeEntity struct {
	Id        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
