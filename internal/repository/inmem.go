package repository

import (
	"persons_service/internal/entity"
	"sync"
)

type inMemoryRepo struct {
	mu    sync.RWMutex
	names map[int]string
}

func NewInMemoryRepository() entity.PersonRepo {
	return &inMemoryRepo{
		names: make(map[int]string),
	}
}

func (r *inMemoryRepo) Create(person *entity.Person) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.names[person.ID] = person.Name
	return nil
}

func (r *inMemoryRepo) Get(id int) (*entity.Person, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	name, exists := r.names[id]
	if !exists {
		return nil, entity.ErrNotFound
	}
	return &entity.Person{ID: id, Name: name}, nil
}

func (r *inMemoryRepo) ExistsById(id int) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, exists := r.names[id]
	return exists, nil
}

func (r *inMemoryRepo) ExistsByName(name string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, n := range r.names {
		if n == name {
			return true, nil
		}
	}
	return false, nil
}

func (r *inMemoryRepo) GetAllNames() ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.names))
	for _, name := range r.names {
		names = append(names, name)
	}
	return names, nil
}
