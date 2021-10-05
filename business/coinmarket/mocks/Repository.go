// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	coinmarket "aprian1337/thukul-service/business/coinmarket"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetBySymbol provides a mock function with given fields: ctx, symbol
func (_m *Repository) GetBySymbol(ctx context.Context, symbol string) (coinmarket.Domain, error) {
	ret := _m.Called(ctx, symbol)

	var r0 coinmarket.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) coinmarket.Domain); ok {
		r0 = rf(ctx, symbol)
	} else {
		r0 = ret.Get(0).(coinmarket.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrice provides a mock function with given fields: ctx, symbol, amount
func (_m *Repository) GetPrice(ctx context.Context, symbol string, amount float64) (float64, error) {
	ret := _m.Called(ctx, symbol, amount)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string, float64) float64); ok {
		r0 = rf(ctx, symbol, amount)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, float64) error); ok {
		r1 = rf(ctx, symbol, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
