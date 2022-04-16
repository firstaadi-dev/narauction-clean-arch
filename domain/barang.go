package domain

import (
	"github.com/lib/pq"
	"time"
)

type Barang struct {
	ID             uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt      time.Time      `json:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      *time.Time     `json:"deletedAt"`
	Lot            string         `json:"lot,omitempty"`
	NamaBarang     string         `json:"namaBarang" json:"namaBarang,omitempty"`
	Foto           pq.StringArray `gorm:"type:varchar[]" json:"foto,omitempty"`
	TahunPembuatan int            `json:"tahunPembuatan,omitempty"`
	NamaPembuat    string         `json:"namaPembuat,omitempty"`
	DescId         string         `json:"descId,omitempty"`
	DescEn         string         `json:"descEn,omitempty"`
	PriceRange     pq.Int64Array  `gorm:"type:integer[]" json:"priceRange,omitempty"`
	Size           pq.Int64Array  `gorm:"type:integer[]" json:"size,omitempty"`
	Tipe           string         `json:"tipe,omitempty"`
	EventID        uint           `json:"eventID,omitempty"`
	AsalDaerah     string         `json:"asalDaerah,omitempty"`
	DyeType        string         `json:"dyeType,omitempty"`
	HargaAwal      int            `json:"hargaAwal,omitempty"`
	UrlThumbnail   pq.StringArray `gorm:"type:varchar[]" json:"urlThumbnail,omitempty"`
	IsAvailable    bool           `json:"isAvailable,omitempty"`
	//Event          Event
}

type BarangMin struct {
	ID             uint       `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt      time.Time  `json:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
	NamaBarang     string     `json:"namaBarang,omitempty"`
	TahunPembuatan int        `json:"tahunPembuatan,omitempty"`
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
