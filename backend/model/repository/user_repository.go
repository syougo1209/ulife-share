package repository

import (
	"errors"
	"fmt"

	"github.com/syougo1209/ulife-share/model/entity"
)

type UserRepository interface {
	InsertUser(user entity.UserEntity) error
}
type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) InsertUser(user entity.UserEntity) error {
	if err := dbmap.Insert(user); err != nil {
		return errors.New(fmt.Sprintf("Error: %s", "insert user error"))
	}
	return nil
}

type userDTO struct {
}
