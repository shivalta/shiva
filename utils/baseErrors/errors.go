package baseErrors

import "errors"

var (
	ErrNotFound          = errors.New("data tidak ditemukan :(")
	ErrInvalidPayload    = errors.New("payload tidak valid")
	ErrInvalidAuth       = errors.New("autentikasi tidak valid")
	ErrUserEmailNotFound = errors.New("email tidak ditemukan")
	ErrUserNotActive     = errors.New("akun belum terverifikasi")

	//USERS MODEL
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
)
