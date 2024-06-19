// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	domain "Posts/internal/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// PostRepository is an autogenerated mock type for the PostRepository type
type PostRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, entity
func (_m *PostRepository) Create(ctx context.Context, entity *domain.Post) error {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *PostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, limit, offset
func (_m *PostRepository) GetAll(ctx context.Context, limit int, offset int) ([]*domain.Post, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]*domain.Post, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*domain.Post); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAuthorID provides a mock function with given fields: ctx, userID, limit, offset
func (_m *PostRepository) GetByAuthorID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*domain.Post, error) {
	ret := _m.Called(ctx, userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetByAuthorID")
	}

	var r0 []*domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int, int) ([]*domain.Post, error)); ok {
		return rf(ctx, userID, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int, int) []*domain.Post); ok {
		r0 = rf(ctx, userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, int, int) error); ok {
		r1 = rf(ctx, userID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *PostRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.Post, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIds provides a mock function with given fields: ctx, ids
func (_m *PostRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]*domain.Post, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetByIds")
	}

	var r0 []*domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]*domain.Post, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []*domain.Post); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, entity
func (_m *PostRepository) Update(ctx context.Context, entity *domain.Post) error {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPostRepository creates a new instance of PostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostRepository {
	mock := &PostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
