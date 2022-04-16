package usecase

import "github.com/firstaadi-dev/narauction-clean-arch/domain"

type EventUsecase struct {
	EventRepository domain.EventRepository
}

// Delete implements domain.EventUsecase
func (*EventUsecase) Delete(id uint) error {
	panic("unimplemented")
}

// Store implements domain.EventUsecase
func (*EventUsecase) Store(event domain.Event) error {
	panic("unimplemented")
}

// Update implements domain.EventUsecase
func (*EventUsecase) Update(id uint, event domain.Event) error {
	panic("unimplemented")
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
