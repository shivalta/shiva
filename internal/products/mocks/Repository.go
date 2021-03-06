// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	products "shiva/shiva-auth/internal/products"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: class
func (_m *Repository) Create(class products.Domain) (products.Domain, error) {
	ret := _m.Called(class)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(products.Domain) products.Domain); ok {
		r0 = rf(class)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.Domain) error); ok {
		r1 = rf(class)
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

// GetAll provides a mock function with given fields: search, key
func (_m *Repository) GetAll(search string, key string) ([]products.Domain, error) {
	ret := _m.Called(search, key)

	var r0 []products.Domain
	if rf, ok := ret.Get(0).(func(string, string) []products.Domain); ok {
		r0 = rf(search, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(search, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id uint) (products.Domain, error) {
	ret := _m.Called(id)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(uint) products.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: class
func (_m *Repository) Update(class products.Domain) (products.Domain, error) {
	ret := _m.Called(class)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(products.Domain) products.Domain); ok {
		r0 = rf(class)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.Domain) error); ok {
		r1 = rf(class)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWithoutImage provides a mock function with given fields: class
func (_m *Repository) UpdateWithoutImage(class products.Domain) (products.Domain, error) {
	ret := _m.Called(class)

	var r0 products.Domain
	if rf, ok := ret.Get(0).(func(products.Domain) products.Domain); ok {
		r0 = rf(class)
	} else {
		r0 = ret.Get(0).(products.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.Domain) error); ok {
		r1 = rf(class)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
