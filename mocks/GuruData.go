// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	guru "Gurumu/features/guru"

	mock "github.com/stretchr/testify/mock"
)

// GuruData is an autogenerated mock type for the GuruData type
type GuruData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *GuruData) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBeranda provides a mock function with given fields: loc, subj, limit, offset
func (_m *GuruData) GetBeranda(loc string, subj string, limit int, offset int) (int, []guru.Core, error) {
	ret := _m.Called(loc, subj, limit, offset)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, int, int) int); ok {
		r0 = rf(loc, subj, limit, offset)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []guru.Core
	if rf, ok := ret.Get(1).(func(string, string, int, int) []guru.Core); ok {
		r1 = rf(loc, subj, limit, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]guru.Core)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, int, int) error); ok {
		r2 = rf(loc, subj, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *GuruData) GetByID(id uint) (interface{}, error) {
	ret := _m.Called(id)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(uint) interface{}); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newGuru
func (_m *GuruData) Register(newGuru guru.Core) (guru.Core, error) {
	ret := _m.Called(newGuru)

	var r0 guru.Core
	if rf, ok := ret.Get(0).(func(guru.Core) guru.Core); ok {
		r0 = rf(newGuru)
	} else {
		r0 = ret.Get(0).(guru.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(guru.Core) error); ok {
		r1 = rf(newGuru)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updateData
func (_m *GuruData) Update(id uint, updateData guru.Core) error {
	ret := _m.Called(id, updateData)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, guru.Core) error); ok {
		r0 = rf(id, updateData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verifikasi provides a mock function with given fields: cekdata
func (_m *GuruData) Verifikasi(cekdata guru.Core) bool {
	ret := _m.Called(cekdata)

	var r0 bool
	if rf, ok := ret.Get(0).(func(guru.Core) bool); ok {
		r0 = rf(cekdata)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewGuruData interface {
	mock.TestingT
	Cleanup(func())
}

// NewGuruData creates a new instance of GuruData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGuruData(t mockConstructorTestingTNewGuruData) *GuruData {
	mock := &GuruData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
