package entity

// интерфейс хранилища
type PersonRepo interface {
	Create(person *Person) error
	Get(id int) (*Person, error)
	ExistsById(id int) (bool, error)
	ExistsByName(name string) (bool, error)
	GetAllNames() ([]string, error)
}

// бизнес-сущность
type Person struct {
	ID   int
	Name string
}
