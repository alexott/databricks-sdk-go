// Code generated by mockery v2.39.1. DO NOT EDIT.

package oauth2

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	oauth2 "github.com/databricks/databricks-sdk-go/service/oauth2"
)

// MockPublishedAppIntegrationInterface is an autogenerated mock type for the PublishedAppIntegrationInterface type
type MockPublishedAppIntegrationInterface struct {
	mock.Mock
}

type MockPublishedAppIntegrationInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPublishedAppIntegrationInterface) EXPECT() *MockPublishedAppIntegrationInterface_Expecter {
	return &MockPublishedAppIntegrationInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationInterface) Create(ctx context.Context, request oauth2.CreatePublishedAppIntegration) (*oauth2.CreatePublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *oauth2.CreatePublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.CreatePublishedAppIntegration) (*oauth2.CreatePublishedAppIntegrationOutput, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.CreatePublishedAppIntegration) *oauth2.CreatePublishedAppIntegrationOutput); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.CreatePublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.CreatePublishedAppIntegration) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockPublishedAppIntegrationInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.CreatePublishedAppIntegration
func (_e *MockPublishedAppIntegrationInterface_Expecter) Create(ctx interface{}, request interface{}) *MockPublishedAppIntegrationInterface_Create_Call {
	return &MockPublishedAppIntegrationInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockPublishedAppIntegrationInterface_Create_Call) Run(run func(ctx context.Context, request oauth2.CreatePublishedAppIntegration)) *MockPublishedAppIntegrationInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.CreatePublishedAppIntegration))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Create_Call) Return(_a0 *oauth2.CreatePublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Create_Call) RunAndReturn(run func(context.Context, oauth2.CreatePublishedAppIntegration) (*oauth2.CreatePublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationInterface) Delete(ctx context.Context, request oauth2.DeletePublishedAppIntegrationRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.DeletePublishedAppIntegrationRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublishedAppIntegrationInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockPublishedAppIntegrationInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.DeletePublishedAppIntegrationRequest
func (_e *MockPublishedAppIntegrationInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockPublishedAppIntegrationInterface_Delete_Call {
	return &MockPublishedAppIntegrationInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockPublishedAppIntegrationInterface_Delete_Call) Run(run func(ctx context.Context, request oauth2.DeletePublishedAppIntegrationRequest)) *MockPublishedAppIntegrationInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.DeletePublishedAppIntegrationRequest))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Delete_Call) Return(_a0 error) *MockPublishedAppIntegrationInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Delete_Call) RunAndReturn(run func(context.Context, oauth2.DeletePublishedAppIntegrationRequest) error) *MockPublishedAppIntegrationInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByIntegrationId provides a mock function with given fields: ctx, integrationId
func (_m *MockPublishedAppIntegrationInterface) DeleteByIntegrationId(ctx context.Context, integrationId string) error {
	ret := _m.Called(ctx, integrationId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByIntegrationId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, integrationId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByIntegrationId'
type MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call struct {
	*mock.Call
}

// DeleteByIntegrationId is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationId string
func (_e *MockPublishedAppIntegrationInterface_Expecter) DeleteByIntegrationId(ctx interface{}, integrationId interface{}) *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call {
	return &MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call{Call: _e.mock.On("DeleteByIntegrationId", ctx, integrationId)}
}

func (_c *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call) Run(run func(ctx context.Context, integrationId string)) *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call) Return(_a0 error) *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call) RunAndReturn(run func(context.Context, string) error) *MockPublishedAppIntegrationInterface_DeleteByIntegrationId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationInterface) Get(ctx context.Context, request oauth2.GetPublishedAppIntegrationRequest) (*oauth2.GetPublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *oauth2.GetPublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.GetPublishedAppIntegrationRequest) (*oauth2.GetPublishedAppIntegrationOutput, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.GetPublishedAppIntegrationRequest) *oauth2.GetPublishedAppIntegrationOutput); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.GetPublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.GetPublishedAppIntegrationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockPublishedAppIntegrationInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.GetPublishedAppIntegrationRequest
func (_e *MockPublishedAppIntegrationInterface_Expecter) Get(ctx interface{}, request interface{}) *MockPublishedAppIntegrationInterface_Get_Call {
	return &MockPublishedAppIntegrationInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockPublishedAppIntegrationInterface_Get_Call) Run(run func(ctx context.Context, request oauth2.GetPublishedAppIntegrationRequest)) *MockPublishedAppIntegrationInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.GetPublishedAppIntegrationRequest))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Get_Call) Return(_a0 *oauth2.GetPublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Get_Call) RunAndReturn(run func(context.Context, oauth2.GetPublishedAppIntegrationRequest) (*oauth2.GetPublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIntegrationId provides a mock function with given fields: ctx, integrationId
func (_m *MockPublishedAppIntegrationInterface) GetByIntegrationId(ctx context.Context, integrationId string) (*oauth2.GetPublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx, integrationId)

	if len(ret) == 0 {
		panic("no return value specified for GetByIntegrationId")
	}

	var r0 *oauth2.GetPublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*oauth2.GetPublishedAppIntegrationOutput, error)); ok {
		return rf(ctx, integrationId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *oauth2.GetPublishedAppIntegrationOutput); ok {
		r0 = rf(ctx, integrationId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.GetPublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, integrationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationInterface_GetByIntegrationId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIntegrationId'
type MockPublishedAppIntegrationInterface_GetByIntegrationId_Call struct {
	*mock.Call
}

// GetByIntegrationId is a helper method to define mock.On call
//   - ctx context.Context
//   - integrationId string
func (_e *MockPublishedAppIntegrationInterface_Expecter) GetByIntegrationId(ctx interface{}, integrationId interface{}) *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call {
	return &MockPublishedAppIntegrationInterface_GetByIntegrationId_Call{Call: _e.mock.On("GetByIntegrationId", ctx, integrationId)}
}

func (_c *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call) Run(run func(ctx context.Context, integrationId string)) *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call) Return(_a0 *oauth2.GetPublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call) RunAndReturn(run func(context.Context, string) (*oauth2.GetPublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationInterface_GetByIntegrationId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockPublishedAppIntegrationInterface) Impl() oauth2.PublishedAppIntegrationService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 oauth2.PublishedAppIntegrationService
	if rf, ok := ret.Get(0).(func() oauth2.PublishedAppIntegrationService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oauth2.PublishedAppIntegrationService)
		}
	}

	return r0
}

// MockPublishedAppIntegrationInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockPublishedAppIntegrationInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockPublishedAppIntegrationInterface_Expecter) Impl() *MockPublishedAppIntegrationInterface_Impl_Call {
	return &MockPublishedAppIntegrationInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockPublishedAppIntegrationInterface_Impl_Call) Run(run func()) *MockPublishedAppIntegrationInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Impl_Call) Return(_a0 oauth2.PublishedAppIntegrationService) *MockPublishedAppIntegrationInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Impl_Call) RunAndReturn(run func() oauth2.PublishedAppIntegrationService) *MockPublishedAppIntegrationInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockPublishedAppIntegrationInterface) List(ctx context.Context) listing.Iterator[oauth2.GetPublishedAppIntegrationOutput] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[oauth2.GetPublishedAppIntegrationOutput]
	if rf, ok := ret.Get(0).(func(context.Context) listing.Iterator[oauth2.GetPublishedAppIntegrationOutput]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[oauth2.GetPublishedAppIntegrationOutput])
		}
	}

	return r0
}

// MockPublishedAppIntegrationInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockPublishedAppIntegrationInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPublishedAppIntegrationInterface_Expecter) List(ctx interface{}) *MockPublishedAppIntegrationInterface_List_Call {
	return &MockPublishedAppIntegrationInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockPublishedAppIntegrationInterface_List_Call) Run(run func(ctx context.Context)) *MockPublishedAppIntegrationInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_List_Call) Return(_a0 listing.Iterator[oauth2.GetPublishedAppIntegrationOutput]) *MockPublishedAppIntegrationInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_List_Call) RunAndReturn(run func(context.Context) listing.Iterator[oauth2.GetPublishedAppIntegrationOutput]) *MockPublishedAppIntegrationInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockPublishedAppIntegrationInterface) ListAll(ctx context.Context) ([]oauth2.GetPublishedAppIntegrationOutput, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []oauth2.GetPublishedAppIntegrationOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]oauth2.GetPublishedAppIntegrationOutput, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []oauth2.GetPublishedAppIntegrationOutput); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]oauth2.GetPublishedAppIntegrationOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublishedAppIntegrationInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockPublishedAppIntegrationInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPublishedAppIntegrationInterface_Expecter) ListAll(ctx interface{}) *MockPublishedAppIntegrationInterface_ListAll_Call {
	return &MockPublishedAppIntegrationInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockPublishedAppIntegrationInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockPublishedAppIntegrationInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_ListAll_Call) Return(_a0 []oauth2.GetPublishedAppIntegrationOutput, _a1 error) *MockPublishedAppIntegrationInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]oauth2.GetPublishedAppIntegrationOutput, error)) *MockPublishedAppIntegrationInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockPublishedAppIntegrationInterface) Update(ctx context.Context, request oauth2.UpdatePublishedAppIntegration) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.UpdatePublishedAppIntegration) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublishedAppIntegrationInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPublishedAppIntegrationInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request oauth2.UpdatePublishedAppIntegration
func (_e *MockPublishedAppIntegrationInterface_Expecter) Update(ctx interface{}, request interface{}) *MockPublishedAppIntegrationInterface_Update_Call {
	return &MockPublishedAppIntegrationInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockPublishedAppIntegrationInterface_Update_Call) Run(run func(ctx context.Context, request oauth2.UpdatePublishedAppIntegration)) *MockPublishedAppIntegrationInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.UpdatePublishedAppIntegration))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Update_Call) Return(_a0 error) *MockPublishedAppIntegrationInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_Update_Call) RunAndReturn(run func(context.Context, oauth2.UpdatePublishedAppIntegration) error) *MockPublishedAppIntegrationInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockPublishedAppIntegrationInterface) WithImpl(impl oauth2.PublishedAppIntegrationService) oauth2.PublishedAppIntegrationInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 oauth2.PublishedAppIntegrationInterface
	if rf, ok := ret.Get(0).(func(oauth2.PublishedAppIntegrationService) oauth2.PublishedAppIntegrationInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(oauth2.PublishedAppIntegrationInterface)
		}
	}

	return r0
}

// MockPublishedAppIntegrationInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockPublishedAppIntegrationInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl oauth2.PublishedAppIntegrationService
func (_e *MockPublishedAppIntegrationInterface_Expecter) WithImpl(impl interface{}) *MockPublishedAppIntegrationInterface_WithImpl_Call {
	return &MockPublishedAppIntegrationInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockPublishedAppIntegrationInterface_WithImpl_Call) Run(run func(impl oauth2.PublishedAppIntegrationService)) *MockPublishedAppIntegrationInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(oauth2.PublishedAppIntegrationService))
	})
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_WithImpl_Call) Return(_a0 oauth2.PublishedAppIntegrationInterface) *MockPublishedAppIntegrationInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublishedAppIntegrationInterface_WithImpl_Call) RunAndReturn(run func(oauth2.PublishedAppIntegrationService) oauth2.PublishedAppIntegrationInterface) *MockPublishedAppIntegrationInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPublishedAppIntegrationInterface creates a new instance of MockPublishedAppIntegrationInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPublishedAppIntegrationInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPublishedAppIntegrationInterface {
	mock := &MockPublishedAppIntegrationInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}