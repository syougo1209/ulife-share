package datastore

import (
	"fmt"
	"time"

	"github.com/syougo1209/ulife-share/model/entity"
	domain "github.com/syougo1209/ulife-share/model/repository"
)

type LifeRepository struct{}

func NewLifeRepository() domain.LifeRepository {
	return &LifeRepository{}
}

func (l *LifeRepository) Fetch() ([]entity.Life, error) {
	var dto []lifeDTO
	if _, err := dbmap.Select(&dto, "select * from life"); err != nil {
		return nil, fmt.Errorf("LifeのFetch中にエラー発生: %w", err)
	}

	lifes := make([]entity.Life, len(dto))
	for i, v := range dto {
		lifes[i] = entity.Life{
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
