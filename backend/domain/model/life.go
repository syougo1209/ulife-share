package model

import "time"

type Life struct {
	Id        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
