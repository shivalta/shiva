// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	accounts "shiva/shiva-auth/internal/accounts"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: id, password
func (_m *Repository) ChangePassword(id uint, password string) (accounts.Domain, error) {
	ret := _m.Called(id, password)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(uint, string) accounts.Domain); ok {
		r0 = rf(id, password)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: user
func (_m *Repository) Create(user accounts.Domain) (accounts.Domain, error) {
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
func (_m *Repository) Delete(id uint) error {
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
func (_m *Repository) GetAll(search string) ([]accounts.Domain, error) {
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

// GetByEmail provides a mock function with given fields: email
func (_m *Repository) GetByEmail(email string) (accounts.Domain, error) {
	ret := _m.Called(email)

	var r0 accounts.Domain
	if rf, ok := ret.Get(0).(func(string) accounts.Domain); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(accounts.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id uint) (accounts.Domain, error) {
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

// Update provides a mock function with given fields: user
func (_m *Repository) Update(user accounts.Domain) (accounts.Domain, error) {
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

// UpdateStatus provides a mock function with given fields: id, state
func (_m *Repository) UpdateStatus(id uint, state bool) error {
	ret := _m.Called(id, state)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, bool) error); ok {
		r0 = rf(id, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWithPassword provides a mock function with given fields: user
func (_m *Repository) UpdateWithPassword(user accounts.Domain) (accounts.Domain, error) {
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
