package repository

import (
	"time"

	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"gorm.io/gorm"
)

type psqlEventRepository struct {
	DB *gorm.DB
}

// GetById implements domain.EventRepository
func (p *psqlEventRepository) GetById(id uint) (res *domain.Event, err error) {
	var event domain.Event
	err = p.DB.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Fetch implements domain.EventRepository
func (p *psqlEventRepository) Fetch() (res []domain.Event, err error) {
	event := make([]domain.Event, 0)
	p.DB.Order("date desc").Find(&event)
	return event, nil
}

// GetUpcoming implements domain.EventRepository
func (p *psqlEventRepository) GetUpcoming() (res *domain.UpcomingEvent, err error) {
	var event domain.Event
	now := time.Now().Format("2006-01-02 15:04:05")
	err = p.DB.Where("date > ?", now).First(&event).Error
	if err != nil {
		return nil, err
	}
	ue := p.GenerateUpcoming(event)
	return ue, nil
}

// Store implements domain.EventRepository
func (p *psqlEventRepository) Store(event domain.Event) (err error) {
	err = p.DB.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}

// Update implements domain.EventRepository
func (p *psqlEventRepository) Update(id uint, event domain.Event) (err error) {
	err = p.DB.Model(&event).Where("id = ?", id).Updates(&event).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.EventRepository
func (p *psqlEventRepository) Delete(id uint) (err error) {
	err = p.DB.Delete(&domain.Event{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// GenerateUpcoming implements domain.EventRepository
func (p *psqlEventRepository) GenerateUpcoming(e domain.Event) (res *domain.UpcomingEvent) {
	var ue domain.UpcomingEvent
	now := time.Now()
	ue.Event = e
	ue.DayDelta = int(e.Date.Sub(now).Hours() / 24)
	if ue.DayDelta >= 0 {
		ue.Status = "upcoming event"
	} else {
		ue.Status = "past event"
	}
	getFoto := p.GetFotoAndCount(e.ID)
	ue.FotoItem = getFoto.FotoItem
	ue.ItemCount = getFoto.ItemCount

	return &ue
}

// GetFoto implements domain.EventRepository
func (p *psqlEventRepository) GetFotoAndCount(id uint) (res *domain.UpcomingEvent) {
	var ue domain.UpcomingEvent
	var fl []string
	barang := make([]domain.Barang, 0)
	result := p.DB.Where("event_id = ?", id).Find(&barang)
	ue.ItemCount = int(result.RowsAffected)
	for _, v := range barang {
		if len(fl) < 5 {
			fl = append(fl, v.Foto[0])
		}
	}
	ue.FotoItem = fl
	return &ue
}

func NewPsqlEventRepository(Db *gorm.DB) domain.EventRepository {
	return &psqlEventRepository{DB: Db}
}
