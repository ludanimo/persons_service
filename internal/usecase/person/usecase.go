package person

import (
	"errors"
	"persons_service/internal/entity"
	"persons_service/internal/interface/repository"
	"strings"
	"unicode"
)

type personUsecase struct {
	repo repository.PersonRepo
}

func NewPersonUsecase(repo repository.PersonRepo) PersonUsecase {
	return &personUsecase{repo: repo}
}

func (uc *personUsecase) Save(person *entity.Person) error {
	// Валидация ID
	if person.ID <= 0 {
		return ErrInvalidID
	}

	// Валидация имени
	if len(person.Name) == 0 {
		return ErrEmptyName
	}

	// Форматирование имени
	formattedName := formatName(person.Name)

	// Валидация символов
	for _, symbol := range formattedName {
		if !(unicode.Is(unicode.Cyrillic, symbol) || unicode.Is(unicode.Latin, symbol)) {
			return ErrInvalidName
		}
	}

	// Проверка уникальности ID
	if exists, _ := uc.repo.ExistsById(person.ID); exists {
		return ErrDuplicateID
	}

	// Проверка уникальности имени
	if exists, _ := uc.repo.ExistsByName(formattedName); exists {
		return ErrDuplicateName
	}

	// Сохранение
	return uc.repo.Save(&entity.Person{
		ID:   person.ID,
		Name: formattedName,
	})
}

func (uc *personUsecase) Get(id int) (*entity.Person, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	return uc.repo.Get(id)
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

// Ошибки
var (
	ErrInvalidID     = errors.New("invalid ID")
	ErrEmptyName     = errors.New("empty name")
	ErrInvalidName   = errors.New("invalid name characters")
	ErrDuplicateID   = errors.New("duplicate ID")
	ErrDuplicateName = errors.New("duplicate name")
)
