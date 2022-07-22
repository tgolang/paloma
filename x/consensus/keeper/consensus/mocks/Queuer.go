// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	consensustypes "github.com/palomachain/paloma/x/consensus/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Queuer is an autogenerated mock type for the Queuer type
type Queuer struct {
	mock.Mock
}

// AddEvidence provides a mock function with given fields: ctx, id, evidence
func (_m *Queuer) AddEvidence(ctx types.Context, id uint64, evidence *consensustypes.Evidence) error {
	ret := _m.Called(ctx, id, evidence)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, *consensustypes.Evidence) error); ok {
		r0 = rf(ctx, id, evidence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddSignature provides a mock function with given fields: ctx, id, signData
func (_m *Queuer) AddSignature(ctx types.Context, id uint64, signData *consensustypes.SignData) error {
	ret := _m.Called(ctx, id, signData)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, *consensustypes.SignData) error); ok {
		r0 = rf(ctx, id, signData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChainInfo provides a mock function with given fields:
func (_m *Queuer) ChainInfo() (string, string) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: _a0
func (_m *Queuer) GetAll(_a0 types.Context) ([]consensustypes.QueuedSignedMessageI, error) {
	ret := _m.Called(_a0)

	var r0 []consensustypes.QueuedSignedMessageI
	if rf, ok := ret.Get(0).(func(types.Context) []consensustypes.QueuedSignedMessageI); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]consensustypes.QueuedSignedMessageI)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMsgByID provides a mock function with given fields: ctx, id
func (_m *Queuer) GetMsgByID(ctx types.Context, id uint64) (consensustypes.QueuedSignedMessageI, error) {
	ret := _m.Called(ctx, id)

	var r0 consensustypes.QueuedSignedMessageI
	if rf, ok := ret.Get(0).(func(types.Context, uint64) consensustypes.QueuedSignedMessageI); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(consensustypes.QueuedSignedMessageI)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPublicAccessData provides a mock function with given fields: ctx, id
func (_m *Queuer) GetPublicAccessData(ctx types.Context, id uint64) (*consensustypes.PublicAccessData, error) {
	ret := _m.Called(ctx, id)

	var r0 *consensustypes.PublicAccessData
	if rf, ok := ret.Get(0).(func(types.Context, uint64) *consensustypes.PublicAccessData); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*consensustypes.PublicAccessData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: _a0, _a1
func (_m *Queuer) Put(_a0 types.Context, _a1 ...consensustypes.ConsensusMsg) error {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, ...consensustypes.ConsensusMsg) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: _a0, _a1
func (_m *Queuer) Remove(_a0 types.Context, _a1 uint64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPublicAccessData provides a mock function with given fields: ctx, id, data
func (_m *Queuer) SetPublicAccessData(ctx types.Context, id uint64, data *consensustypes.PublicAccessData) error {
	ret := _m.Called(ctx, id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64, *consensustypes.PublicAccessData) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewQueuer interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueuer creates a new instance of Queuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueuer(t mockConstructorTestingTNewQueuer) *Queuer {
	mock := &Queuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
