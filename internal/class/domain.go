package class

type Domain struct {
	ID      uint
	Name    string
	IsPasca bool
	Image   string
	Slug    string
}

type Usecase interface {
	GetAll(search string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	Create(user Domain) (Domain, error)
	Update(user Domain) (Domain, error)
	Delete(id uint) error
}

type Repository interface {
	Update(user Domain) (Domain, error)
	UpdateStatus(id uint, state bool) error
	Delete(id uint) error
	Create(user Domain) (Domain, error)
	GetByEmail(email string) (Domain, error)
	GetAll(search string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	ChangePassword(id uint, password string) (Domain, error)
}
