// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "mygram/model"

	mock "github.com/stretchr/testify/mock"
)

// IPhotoRepository is an autogenerated mock type for the IPhotoRepository type
type IPhotoRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *IPhotoRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *IPhotoRepository) Get() ([]model.Photo, error) {
	ret := _m.Called()

	var r0 []model.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Photo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Photo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Photo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: id
func (_m *IPhotoRepository) GetOne(id string) (model.Photo, error) {
	ret := _m.Called(id)

	var r0 model.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Photo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) model.Photo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Photo)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: photo
func (_m *IPhotoRepository) Save(photo model.Photo) (model.Photo, error) {
	ret := _m.Called(photo)

	var r0 model.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Photo) (model.Photo, error)); ok {
		return rf(photo)
	}
	if rf, ok := ret.Get(0).(func(model.Photo) model.Photo); ok {
		r0 = rf(photo)
	} else {
		r0 = ret.Get(0).(model.Photo)
	}

	if rf, ok := ret.Get(1).(func(model.Photo) error); ok {
		r1 = rf(photo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: updatePhoto, id
func (_m *IPhotoRepository) Update(updatePhoto model.Photo, id string) (model.Photo, error) {
	ret := _m.Called(updatePhoto, id)

	var r0 model.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Photo, string) (model.Photo, error)); ok {
		return rf(updatePhoto, id)
	}
	if rf, ok := ret.Get(0).(func(model.Photo, string) model.Photo); ok {
		r0 = rf(updatePhoto, id)
	} else {
		r0 = ret.Get(0).(model.Photo)
	}

	if rf, ok := ret.Get(1).(func(model.Photo, string) error); ok {
		r1 = rf(updatePhoto, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIPhotoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPhotoRepository creates a new instance of IPhotoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPhotoRepository(t mockConstructorTestingTNewIPhotoRepository) *IPhotoRepository {
	mock := &IPhotoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
