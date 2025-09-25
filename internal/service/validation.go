package service

import (
	"persons_service/internal/entity"
	"strings"
	"unicode"
)

//  интерфейс хранилища
type PersonRepo interface {
	Save(person *entity.Person) error
	Get(id int) (*entity.Person, error)
	ExistsById(id int) (bool, error)
	ExistsByName(name string) (bool, error)
	GetAllNames() ([]string, error)
}

// валидация
type ValidationService struct {
	repo PersonRepo
}

func NewValidationService(repo PersonRepo) *ValidationService {
	return &ValidationService{repo: repo}
}

func (s *ValidationService) Save(person *entity.Person) error {

	if person.ID <= 0 {
		return ErrInvalidID
	}

	if len(person.Name) == 0 {
		return ErrEmptyName
	}

	// Форматирование имени
	formattedName := formatName(person.Name)

	for _, symbol := range formattedName {
		if !(unicode.Is(unicode.Cyrillic, symbol) || unicode.Is(unicode.Latin, symbol)) {
			return ErrInvalidName
		}
	}

	if exists, _ := s.repo.ExistsById(person.ID); exists {
		return ErrDuplicateID
	}

	// Проверка уникальности имени
	if exists, _ := s.repo.ExistsByName(formattedName); exists {
		return ErrDuplicateName
	}

	return s.repo.Save(&entity.Person{
		ID:   person.ID,
		Name: formattedName,
	})
}

func (s *ValidationService) Get(id int) (*entity.Person, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.Get(id)
}

func formatName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	runes := []rune(name)
	if len(runes) == 0 {
		return ""
	}
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
