package usecase

import "github.com/firstaadi-dev/narauction-clean-arch/domain"

type BarangUsecase struct {
	BarangRepository domain.BarangRepository
}

func (b BarangUsecase) Fetch() ([]domain.Barang, error) {
	res, err := b.BarangRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b BarangUsecase) GetById(id uint) (*domain.Barang, error) {
	res, err := b.BarangRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b BarangUsecase) GetByEventId(eventId uint) ([]domain.Barang, error) {
	res, err := b.BarangRepository.GetByEventId(eventId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b BarangUsecase) Store(barang domain.Barang) error {
	err := b.BarangRepository.Store(barang)
	if err != nil {
		return err
	}
	return nil
}

func (b BarangUsecase) Update(id uint, barang domain.Barang) error {
	err := b.BarangRepository.Update(id, barang)
	if err != nil {
		return err
	}
	return nil
}

func (b BarangUsecase) Delete(id uint) error {
	err := b.BarangRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewBarangUsecase(a domain.BarangRepository) domain.BarangUsecase {
	return &BarangUsecase{
		BarangRepository: a,
	}
}
