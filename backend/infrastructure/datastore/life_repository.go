package datastore

import (
	"time"

	"github.com/syougo1209/ulife-share/domain/model"
	repository "github.com/syougo1209/ulife-share/domain/repository"
)

type LifeRepository struct{}

func NewLifeRepository() repository.LifeRepository {
	return &LifeRepository{}
}

func (l *LifeRepository) Fetch() ([]model.Life, error) {
	var dto []lifeDTO
	if _, err := dbmap.Select(&dto, "select * from life"); err != nil {
		return nil, model.ErrNotFound
	}

	lifes := make([]model.Life, len(dto))
	for i, v := range dto {
		lifes[i] = model.Life{
			Id:        v.Id,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return lifes, nil
}

type lifeDTO struct {
	Id        int       `db:"id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
