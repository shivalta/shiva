package baseErrors

import "errors"

var (
	ErrNotFound          = errors.New("data tidak ditemukan :(")
	ErrInvalidPayload    = errors.New("payload tidak valid")
	ErrInvalidAuth       = errors.New("autentikasi tidak valid")
	ErrUserEmailNotFound = errors.New("email tidak ditemukan")
	ErrUserNotActive     = errors.New("akun belum terverifikasi")
	ErrRecordNotFound    = errors.New("record not found")

	//USERS
	ErrUserEmailRequired      = errors.New("email wajib diisi!")
	ErrUserEmailUsed          = errors.New("email is has been used")
	ErrUsersNameRequired      = errors.New("name wajib diisi!")
	ErrUsersHandphoneRequired = errors.New("handphone wajib diisi!")
	ErrUsersReqNotValid       = errors.New("request tidak valid")
	ErrUsersAddressquired     = errors.New("address wajib diisi!")
	ErrUsersPasswordRequired  = errors.New("password wajib diisi!")
	ErrInvalidPassword        = errors.New("password tidak valid")
	ErrUsersEmailRequired     = errors.New("email wajib diisi!")
	ErrInvalidTokenCredential = errors.New("token tidak ditemukan atau sudah kadaluarsa")
	ErrEmailSmtp              = errors.New("email tidak terkirim :(")

	//CATEGORIES
	ErrCategoriesIdNotFound = errors.New("product categories id not found!")

	//PRODUCTS
	ErrProductClassIdNotSync = errors.New("product_class_id tidak diterima, karena tidak sama dengan product_class_id pada product_category")

	//ORDERS
	ErrNoHpRequired      = errors.New("nomor hp harus diisi!")
	ErrNoPdamRequired    = errors.New("nomor pdam harus diisi!")
	ErrNoListrikRequired = errors.New("nomor listrik harus diisi!")
)
