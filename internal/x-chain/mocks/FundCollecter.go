// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// FundCollecter is an autogenerated mock type for the FundCollecter type
type FundCollecter struct {
	mock.Mock
}

// CollectJobFundEvents provides a mock function with given fields: ctx
func (_m *FundCollecter) CollectJobFundEvents(ctx types.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFundCollecter interface {
	mock.TestingT
	Cleanup(func())
}

// NewFundCollecter creates a new instance of FundCollecter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFundCollecter(t mockConstructorTestingTNewFundCollecter) *FundCollecter {
	mock := &FundCollecter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
