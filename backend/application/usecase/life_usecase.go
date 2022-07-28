package application

import (
	model "github.com/syougo1209/ulife-share/domain/model"
	repository "github.com/syougo1209/ulife-share/domain/repository"
)

type LifeUseCase interface {
	Fetch() (FetchOutput, error)
}
type lifeUseCase struct {
	lifeRepo repository.LifeRepository
}

func NewLifeUseCase(lr repository.LifeRepository) LifeUseCase {
	return &lifeUseCase{lr}
}

type FetchOutput []model.Life

func (l *lifeUseCase) Fetch() (FetchOutput, error) {
	lifes, err := l.lifeRepo.Fetch()
	if err != nil {
		return nil, err
	}
	lifesOutput := make(FetchOutput, len(lifes))
	for i, v := range lifes {
		lifesOutput[i] = v
	}
	return lifesOutput, nil
}
