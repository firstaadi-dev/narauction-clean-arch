package usecase

import "github.com/firstaadi-dev/narauction-clean-arch/domain"

type EventUsecase struct {
	EventRepository domain.EventRepository
}

// GetById implements domain.EventUsecase
func (e *EventUsecase) GetById(id uint) (*domain.Event, error) {
	res, err := e.EventRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete implements domain.EventUsecase
func (e *EventUsecase) Delete(id uint) error {
	err := e.EventRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Store implements domain.EventUsecase
func (e *EventUsecase) Store(event domain.Event) error {
	err := e.EventRepository.Store(event)
	if err != nil {
		return err
	}
	return nil
}

// Update implements domain.EventUsecase
func (e *EventUsecase) Update(id uint, event domain.Event) error {
	err := e.EventRepository.Update(id, event)
	if err != nil {
		return err
	}
	return nil
}

// GetUpcoming implements domain.EventUsecase
func (e *EventUsecase) GetUpcoming() (*domain.UpcomingEvent, error) {
	res, err := e.EventRepository.GetUpcoming()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Fetch implements domain.EventUsecase
func (e *EventUsecase) Fetch() ([]domain.Event, error) {
	res, err := e.EventRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewEventUsecase(a domain.EventRepository) domain.EventUsecase {
	return &EventUsecase{
		EventRepository: a,
	}
}
