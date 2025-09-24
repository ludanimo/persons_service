package person

import "persons_service/internal/entity"

// PersonUsecase интерфейс для бизнес-сценариев
type PersonUsecase interface {
	Save(person *entity.Person) error
	Get(id int) (*entity.Person, error)
}
