package class

type Domain struct {
	ID      uint
	Name    string
	IsPasca bool
	Image   string
	Slug    string
}

type Usecase interface {
	GetAll(search string, key string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	Create(class Domain) (Domain, error)
	Update(class Domain) (Domain, error)
	Delete(id uint) error
}

type Repository interface {
	GetAll(search string, key string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	Create(class Domain) (Domain, error)
	Update(class Domain) (Domain, error)
	Delete(id uint) error
}
