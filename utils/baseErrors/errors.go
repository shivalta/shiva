package baseErrors

import "errors"

var (
	ErrNotFound    = errors.New("data not found")
	ErrInvalidAuth = errors.New("invalid authentication")
	//USERS MODEL
	ErrUserEmailRequired      = errors.New("email is required")
	ErrUserEmailUsed          = errors.New("email is has been used")
	ErrUsersNameRequired      = errors.New("name is required")
	ErrUsersHandphoneRequired = errors.New("handphone is required")
	ErrUsersReqNotValid       = errors.New("request not valid")
	ErrUsersAddressquired     = errors.New("address is required")
	ErrUsersPasswordRequired  = errors.New("password is required")
	ErrInvalidPassword        = errors.New("password invalid")
	ErrUsersEmailRequired     = errors.New("email is required")
	ErrInvalidTokenCredential = errors.New("token not found or expired")
	ErrEmailSmtp              = errors.New("email not send :(")

	ErrIDNotFound             = errors.New("id not found")
	ErrInvalidId              = errors.New("invalid id, id not numeric")
	ErrUserIdNotFound         = errors.New("user id not found")
	ErrEmailHasBeenRegister   = errors.New("email has been used")
	ErrInsufficientPermission = errors.New("insufficient permission")
)
