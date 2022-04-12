package usecase

import "github.com/firstaadi-dev/narauction-clean-arch/domain"

type EventUsecase struct {
	EventRepository domain.EventRepository
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
