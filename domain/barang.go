package domain

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Lot            string
	NamaBarang     string
	Foto           pq.StringArray `gorm:"type:varchar[]"`
	TahunPembuatan int
	NamaPembuat    string
	DescId         string
	DescEn         string
	PriceRange     pq.Int64Array `gorm:"type:integer[]"`
	Size           pq.Int64Array `gorm:"type:integer[]"`
	Tipe           string
	EventID        uint
	AsalDaerah     string
	DyeType        string
	HargaAwal      int
	UrlThumbnail   pq.StringArray `gorm:"type:varchar[]"`
	IsAvailable    bool
	//Event          Event
}

type BarangMin struct {
	gorm.Model
	NamaBarang     string
	TahunPembuatan int
}

type BarangUsecase interface {
	Fetch() ([]Barang, error)
	GetById(id uint) (*Barang, error)
	GetByEventId(eventId uint) ([]Barang, error)
	Store(barang Barang) error
	Update(id uint, barang Barang) error
	Delete(id uint) error
}

type BarangRepository interface {
	Fetch() (res []Barang, err error)
	GetById(id uint) (res *Barang, err error)
	GetByEventId(eventId uint) (res []Barang, err error)
	Store(barang Barang) (err error)
	Update(id uint, barang Barang) (err error)
	Delete(id uint) (err error)
}
