package repository

import (
	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"gorm.io/gorm"
)

type psqlBarangRepository struct {
	DB *gorm.DB
}

func (p psqlBarangRepository) Fetch() (res []domain.Barang, err error) {
	barang := make([]domain.Barang, 0)
	p.DB.Find(&barang)
	return barang, nil
}

func (p psqlBarangRepository) GetById(id uint) (res *domain.Barang, err error) {
	var barang domain.Barang
	result := p.DB.First(&barang, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &barang, nil
}

func (p psqlBarangRepository) GetByEventId(eventId uint) (res []domain.Barang, err error) {
	barang := make([]domain.Barang, 0)
	err = p.DB.Where("event_id = ?", eventId).Order("Lot").Find(&barang).Error
	if err != nil {
		return nil, err
	}
	return barang, nil
}

func (p psqlBarangRepository) Store(barang domain.Barang) (err error) {
	err = p.DB.Create(&barang).Error
	if err != nil {
		return err
	}
	return nil
}

func (p psqlBarangRepository) Update(id uint, barang domain.Barang) (err error) {
	err = p.DB.Model(&barang).Where("id = ?", id).Updates(&barang).Error
	if err != nil {
		return err
	}
	return nil
}

func (p psqlBarangRepository) Delete(id uint) (err error) {
	err = p.DB.Delete(&domain.Barang{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewPsqlBarangRepository(Db *gorm.DB) domain.BarangRepository {
	return &psqlBarangRepository{DB: Db}
}
