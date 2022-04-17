package domain

import (
	"time"

	"github.com/lib/pq"
)

type Barang struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      *time.Time     `json:"deletedAt"`
	Lot            string         `json:"lot"`
	NamaBarang     string         `json:"namaBarang"`
	Foto           pq.StringArray `gorm:"type:varchar[]" json:"foto"`
	TahunPembuatan int            `json:"tahunPembuatan"`
	NamaPembuat    string         `json:"namaPembuat"`
	DescId         string         `json:"descId"`
	DescEn         string         `json:"descEn"`
	PriceRange     pq.Int64Array  `gorm:"type:integer[]" json:"priceRange"`
	Size           pq.Int64Array  `gorm:"type:integer[]" json:"size"`
	Tipe           string         `json:"tipe"`
	EventID        uint           `json:"eventID"`
	AsalDaerah     string         `json:"asalDaerah"`
	DyeType        string         `json:"dyeType"`
	HargaAwal      int            `json:"hargaAwal"`
	UrlThumbnail   pq.StringArray `gorm:"type:varchar[]" json:"urlThumbnail"`
	IsAvailable    bool           `json:"isAvailable"`
	//Event          Event
}

type BarangMin struct {
	ID             uint       `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
	NamaBarang     string     `json:"namaBarang"`
	TahunPembuatan int        `json:"tahunPembuatan"`
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
