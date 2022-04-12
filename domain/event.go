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
	gorm.Model
	Name      string
	Date      time.Time
	DescId    string
	DescEn    string
	OpeningId string
	OpeningEn string
	Foto      pq.StringArray `gorm:"type:varchar[]"`
	DayDelta  int
	Status    string
	ItemCount int
	FotoItem  pq.StringArray `gorm:"type:varchar[]"`
}

type EventUsecase interface {
	Fetch() ([]Event, error)
}

type EventRepository interface {
	Fetch() (res []Event, err error)
}
