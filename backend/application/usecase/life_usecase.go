package application

import (
	"github.com/syougo1209/ulife-share/model/entity"
	domain "github.com/syougo1209/ulife-share/model/repository"
)

type LifeUseCase interface {
	Fetch() ([]Output, error)
}
type lifeUseCase struct {
	lifeRepo domain.LifeRepository
}

func NewLifeUseCase(lr domain.LifeRepository) LifeUseCase {
	return &lifeUseCase{lr}
}

type Output struct {
	Life entity.Life
}

func (l *lifeUseCase) Fetch() ([]Output, error) {
	lifes, err := l.lifeRepo.Fetch()
	if err != nil {
		return nil, err
	}
	var lifesOutput []Output
	for i, v := range lifes {
		lifesOutput[i].Life = v
	}
	return lifesOutput, nil
}
