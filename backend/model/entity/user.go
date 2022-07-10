package entity

import "time"

type UserEntity struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string
	Introduction string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
