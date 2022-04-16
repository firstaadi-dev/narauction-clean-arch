package domain

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name      string
	Date      time.Time
	DescId    string
	DescEn    string
	OpeningId string
	OpeningEn string
	Foto      pq.StringArray `gorm:"type:varchar[]"`
}
type UpcomingEvent struct {
	Event     Event
	DayDelta  int
	Status    string
	ItemCount int
	FotoItem  pq.StringArray `gorm:"type:varchar[]"`
}

type EventUsecase interface {
	Fetch() ([]Event, error)
	GetUpcoming() (*UpcomingEvent, error)
	Store(event Event) error
	Update(id uint, event Event) error
	Delete(id uint) error
}

type EventRepository interface {
	Fetch() (res []Event, err error)
	GetUpcoming() (res *UpcomingEvent, err error)
	GetFotoAndCount(id uint) (res *UpcomingEvent)
	GenerateUpcoming(e Event) (res *UpcomingEvent)
	Store(event Event) (err error)
	Update(id uint, event Event) (err error)
	Delete(id uint) (err error)
}
