package accounts

type Domain struct {
	ID         uint
	Name       string
	Email      string
	Handphone  string
	Address    string
	IsAdmin    bool
	Password   string
	Repassword string
}

type Usecase interface {
	Update(user Domain) (Domain, error)
	Delete(id uint) error
	Create(user Domain) (Domain, error)
	GetAll(search string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	Login(email string, password string) (string, error)
}

type Repository interface {
	Update(user Domain) (Domain, error)
	Delete(id uint) error
	Create(user Domain) (Domain, error)
	GetByEmail(email string) (Domain, error)
	GetAll(search string) ([]Domain, error)
	GetById(id uint) (Domain, error)
	ChangePassword(id uint, password string) (Domain, error)
}
