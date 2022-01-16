package accounts

type Domain struct {
	ID         uint
	Name       string
	Email      string
	Handphone  string
	Address    string
	Password   string
	Repassword string
}

type Usecase interface {
	Create(user Domain) (data Domain, err error)
	GetAll(search string) (data []Domain, err error)
}

type Repository interface {
	Create(user Domain) (data Domain, err error)
	GetAll(search string) (data []Domain, err error)
}
