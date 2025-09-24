package repository

import "persons_service/internal/entity"

// PersonRepo интерфейс хранилища
type PersonRepo interface {
	Save(person *entity.Person) error
	Get(id int) (*entity.Person, error)
	ExistsById(id int) (bool, error)
	ExistsByName(name string) (bool, error)
	GetAllNames() ([]string, error)
}
