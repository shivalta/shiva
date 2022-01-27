// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	accounts "shiva/shiva-auth/internal/accounts"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *Usecase) Create(user accounts.Domain) (accounts.Domain, error) {
	ret := _m.Called(user)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(accounts.Domain) accounts.Domain); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Domain) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Usecase) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: search
func (_m *Usecase) GetAll(search string) ([]accounts.Domain, error) {
	ret := _m.Called(search)

	var r0 []accounts.Domain
	if rf, ok := ret.Get(0).(func(string) []accounts.Domain); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Usecase) GetById(id uint) (accounts.Domain, error) {
	ret := _m.Called(id)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(uint) accounts.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *Usecase) Login(email string, password string) (accounts.Domain, string, error) {
	ret := _m.Called(email, password)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(string, string) accounts.Domain); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: user
func (_m *Usecase) Update(user accounts.Domain) (accounts.Domain, error) {
	ret := _m.Called(user)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(accounts.Domain) accounts.Domain); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Domain) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: emailBase64, encrypt
func (_m *Usecase) Verify(emailBase64 string, encrypt string) (accounts.Domain, error) {
	ret := _m.Called(emailBase64, encrypt)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(string, string) accounts.Domain); ok {
		r0 = rf(emailBase64, encrypt)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(emailBase64, encrypt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
