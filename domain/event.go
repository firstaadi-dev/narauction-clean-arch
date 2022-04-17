package domain

import (
	"time"

	"github.com/lib/pq"
)

type Event struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt *time.Time     `json:"deletedAt"`
	Name      string         `json:"name,omitempty"`
	Date      time.Time      `json:"date"`
	DescId    string         `json:"descId,omitempty"`
	DescEn    string         `json:"descEn,omitempty"`
	OpeningId string         `json:"openingId,omitempty" `
	OpeningEn string         `json:"openingEn,omitempty"`
	Foto      pq.StringArray `gorm:"type:varchar[]" json:"foto,omitempty"`
}
type UpcomingEvent struct {
	Event     Event          `json:"event"`
	DayDelta  int            `json:"dayDelta,omitempty"`
	Status    string         `json:"status,omitempty"`
	ItemCount int            `json:"itemCount,omitempty"`
	FotoItem  pq.StringArray `gorm:"type:varchar[]" json:"fotoItem,omitempty"`
}

type EventUsecase interface {
	Fetch() ([]Event, error)
	GetById(id uint) (*Event, error)
	GetUpcoming() (*UpcomingEvent, error)
	Store(event Event) error
	Update(id uint, event Event) error
	Delete(id uint) error
}

type EventRepository interface {
	Fetch() (res []Event, err error)
	GetById(id uint) (res *Event, err error)
	GetUpcoming() (res *UpcomingEvent, err error)
	GetFotoAndCount(id uint) (res *UpcomingEvent)
	GenerateUpcoming(e Event) (res *UpcomingEvent)
	Store(event Event) (err error)
	Update(id uint, event Event) (err error)
	Delete(id uint) (err error)
}
