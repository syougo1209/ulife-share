package application

import (
	"github.com/syougo1209/ulife-share/model/entity"
	"github.com/syougo1209/ulife-share/model/repository"
)

type LifeUseCase interface {
}
type lifeUseCase struct {
}

func NewLifeUseCase() LifeUseCase {
	return &lifeUseCase{}
}

type output struct {
	Life entity.Life
}

func (l *lifeUseCase) Fetch() (output, error) {
	repository.UserRepository
}
