package repository

import (
	"github.com/syougo1209/ulife-share/model/entity"
)

type LifeRepository interface {
	Fetch() ([]entity.Life, error)
}
