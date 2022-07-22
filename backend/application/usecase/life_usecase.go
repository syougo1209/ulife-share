package application

import (
	model "github.com/syougo1209/ulife-share/domain/model"
	repository "github.com/syougo1209/ulife-share/domain/repository"
)

type LifeUseCase interface {
	Fetch() ([]Output, error)
}
type lifeUseCase struct {
	lifeRepo repository.LifeRepository
}

func NewLifeUseCase(lr repository.LifeRepository) LifeUseCase {
	return &lifeUseCase{lr}
}

type Output struct {
	Life model.Life
}

func (l *lifeUseCase) Fetch() ([]Output, error) {
	lifes, err := l.lifeRepo.Fetch()
	if err != nil {
		return nil, err
	}
	lifesOutput := make([]Output, len(lifes))
	for i, v := range lifes {
		lifesOutput[i].Life = v
	}
	return lifesOutput, nil
}
