// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	context "context"
	repo "srep/internal/repo"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type RepositoryInterface struct {
	mock.Mock
}

// CreateCharacter provides a mock function with given fields: ctx, character
func (_m *RepositoryInterface) CreateCharacter(ctx context.Context, character repo.Character) (int, error) {
	ret := _m.Called(ctx, character)

	if len(ret) == 0 {
		panic("no return value specified for CreateCharacter")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repo.Character) (int, error)); ok {
		return rf(ctx, character)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repo.Character) int); ok {
		r0 = rf(ctx, character)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repo.Character) error); ok {
		r1 = rf(ctx, character)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCharacter provides a mock function with given fields: ctx, id
func (_m *RepositoryInterface) DeleteCharacter(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCharacter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCharacters provides a mock function with given fields: ctx
func (_m *RepositoryInterface) GetAllCharacters(ctx context.Context) ([]repo.Character, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllCharacters")
	}

	var r0 []repo.Character
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]repo.Character, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []repo.Character); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repo.Character)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCharacter provides a mock function with given fields: ctx, id
func (_m *RepositoryInterface) GetCharacter(ctx context.Context, id int) (*repo.Character, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetCharacter")
	}

	var r0 *repo.Character
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*repo.Character, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *repo.Character); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repo.Character)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCharacter provides a mock function with given fields: ctx, id, updates
func (_m *RepositoryInterface) UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error {
	ret := _m.Called(ctx, id, updates)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCharacter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, map[string]interface{}) error); ok {
		r0 = rf(ctx, id, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositoryInterface creates a new instance of RepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryInterface {
	mock := &RepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}