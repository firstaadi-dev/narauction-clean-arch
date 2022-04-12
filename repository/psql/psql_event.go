package repository

import (
	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"gorm.io/gorm"
)

type psqlEventRepository struct {
	DB *gorm.DB
}

// Fetch implements domain.EventRepository
func (p *psqlEventRepository) Fetch() (res []domain.Event, err error) {
	event := make([]domain.Event, 0)
	p.DB.Find(&event)
	return event, nil
}

func NewPsqlEventRepository(Db *gorm.DB) domain.EventRepository {
	return &psqlEventRepository{DB: Db}
}
