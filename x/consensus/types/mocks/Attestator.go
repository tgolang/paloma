// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	consensustypes "github.com/palomachain/paloma/x/consensus/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Attestator is an autogenerated mock type for the Attestator type
type Attestator struct {
	mock.Mock
}

// ProcessAllEvidence provides a mock function with given fields: ctx, task, evidence
func (_m *Attestator) ProcessAllEvidence(ctx types.Context, task consensustypes.AttestTask, evidence []consensustypes.Evidence) (consensustypes.AttestResult, error) {
	ret := _m.Called(ctx, task, evidence)

	var r0 consensustypes.AttestResult
	if rf, ok := ret.Get(0).(func(types.Context, consensustypes.AttestTask, []consensustypes.Evidence) consensustypes.AttestResult); ok {
		r0 = rf(ctx, task, evidence)
	} else {
		r0 = ret.Get(0).(consensustypes.AttestResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context, consensustypes.AttestTask, []consensustypes.Evidence) error); ok {
		r1 = rf(ctx, task, evidence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateEvidence provides a mock function with given fields: ctx, task, evidence
func (_m *Attestator) ValidateEvidence(ctx types.Context, task consensustypes.AttestTask, evidence consensustypes.Evidence) error {
	ret := _m.Called(ctx, task, evidence)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, consensustypes.AttestTask, consensustypes.Evidence) error); ok {
		r0 = rf(ctx, task, evidence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAttestator interface {
	mock.TestingT
	Cleanup(func())
}

// NewAttestator creates a new instance of Attestator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAttestator(t mockConstructorTestingTNewAttestator) *Attestator {
	mock := &Attestator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
