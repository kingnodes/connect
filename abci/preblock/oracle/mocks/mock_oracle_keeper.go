// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	oracletypes "github.com/skip-mev/slinky/x/oracle/types"

	pkgtypes "github.com/skip-mev/slinky/pkg/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Keeper is an autogenerated mock type for the Keeper type
type Keeper struct {
	mock.Mock
}

// GetAllCurrencyPairs provides a mock function with given fields: ctx
func (_m *Keeper) GetAllCurrencyPairs(ctx types.Context) []pkgtypes.CurrencyPair {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCurrencyPairs")
	}

	var r0 []pkgtypes.CurrencyPair
	if rf, ok := ret.Get(0).(func(types.Context) []pkgtypes.CurrencyPair); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pkgtypes.CurrencyPair)
		}
	}

	return r0
}

// SetPriceForCurrencyPair provides a mock function with given fields: ctx, cp, qp
func (_m *Keeper) SetPriceForCurrencyPair(ctx types.Context, cp pkgtypes.CurrencyPair, qp oracletypes.QuotePrice) error {
	ret := _m.Called(ctx, cp, qp)

	if len(ret) == 0 {
		panic("no return value specified for SetPriceForCurrencyPair")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, pkgtypes.CurrencyPair, oracletypes.QuotePrice) error); ok {
		r0 = rf(ctx, cp, qp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKeeper creates a new instance of Keeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *Keeper {
	mock := &Keeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
