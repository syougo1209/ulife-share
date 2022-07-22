package repository

import (
	model "github.com/syougo1209/ulife-share/domain/model"
)

type LifeRepository interface {
	Fetch() ([]model.Life, error)
}
