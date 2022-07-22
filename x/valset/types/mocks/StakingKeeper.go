// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// IterateValidators provides a mock function with given fields: ctx, fn
func (_m *StakingKeeper) IterateValidators(ctx types.Context, fn func(int64, stakingtypes.ValidatorI) bool) {
	_m.Called(ctx, fn)
}

// Validator provides a mock function with given fields: ctx, addr
func (_m *StakingKeeper) Validator(ctx types.Context, addr types.ValAddress) stakingtypes.ValidatorI {
	ret := _m.Called(ctx, addr)

	var r0 stakingtypes.ValidatorI
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) stakingtypes.ValidatorI); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.ValidatorI)
		}
	}

	return r0
}

type mockConstructorTestingTNewStakingKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakingKeeper(t mockConstructorTestingTNewStakingKeeper) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
