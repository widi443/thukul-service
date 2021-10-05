// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	transactions "aprian1337/thukul-service/business/transactions"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// TransactionsById provides a mock function with given fields: ctx, id
func (_m *Repository) TransactionsById(ctx context.Context, id string) (transactions.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) transactions.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionsCreate provides a mock function with given fields: ctx, domain
func (_m *Repository) TransactionsCreate(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, transactions.Domain) transactions.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, transactions.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionsUpdaterCompleted provides a mock function with given fields: ctx, transactionId, status
func (_m *Repository) TransactionsUpdaterCompleted(ctx context.Context, transactionId string, status int) (transactions.Domain, error) {
	ret := _m.Called(ctx, transactionId, status)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, int) transactions.Domain); ok {
		r0 = rf(ctx, transactionId, status)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, transactionId, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionsUpdaterVerify provides a mock function with given fields: ctx, transactionId
func (_m *Repository) TransactionsUpdaterVerify(ctx context.Context, transactionId string) (transactions.Domain, error) {
	ret := _m.Called(ctx, transactionId)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) transactions.Domain); ok {
		r0 = rf(ctx, transactionId)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, transactionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
