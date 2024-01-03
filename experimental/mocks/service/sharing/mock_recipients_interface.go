// Code generated by mockery v2.39.1. DO NOT EDIT.

package sharing

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	sharing "github.com/databricks/databricks-sdk-go/service/sharing"
)

// MockRecipientsInterface is an autogenerated mock type for the RecipientsInterface type
type MockRecipientsInterface struct {
	mock.Mock
}

type MockRecipientsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRecipientsInterface) EXPECT() *MockRecipientsInterface_Expecter {
	return &MockRecipientsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) Create(ctx context.Context, request sharing.CreateRecipient) (*sharing.RecipientInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *sharing.RecipientInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.CreateRecipient) (*sharing.RecipientInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.CreateRecipient) *sharing.RecipientInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.RecipientInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.CreateRecipient) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRecipientsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.CreateRecipient
func (_e *MockRecipientsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockRecipientsInterface_Create_Call {
	return &MockRecipientsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockRecipientsInterface_Create_Call) Run(run func(ctx context.Context, request sharing.CreateRecipient)) *MockRecipientsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.CreateRecipient))
	})
	return _c
}

func (_c *MockRecipientsInterface_Create_Call) Return(_a0 *sharing.RecipientInfo, _a1 error) *MockRecipientsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_Create_Call) RunAndReturn(run func(context.Context, sharing.CreateRecipient) (*sharing.RecipientInfo, error)) *MockRecipientsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) Delete(ctx context.Context, request sharing.DeleteRecipientRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.DeleteRecipientRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRecipientsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRecipientsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.DeleteRecipientRequest
func (_e *MockRecipientsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockRecipientsInterface_Delete_Call {
	return &MockRecipientsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockRecipientsInterface_Delete_Call) Run(run func(ctx context.Context, request sharing.DeleteRecipientRequest)) *MockRecipientsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.DeleteRecipientRequest))
	})
	return _c
}

func (_c *MockRecipientsInterface_Delete_Call) Return(_a0 error) *MockRecipientsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientsInterface_Delete_Call) RunAndReturn(run func(context.Context, sharing.DeleteRecipientRequest) error) *MockRecipientsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockRecipientsInterface) DeleteByName(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRecipientsInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockRecipientsInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockRecipientsInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockRecipientsInterface_DeleteByName_Call {
	return &MockRecipientsInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockRecipientsInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockRecipientsInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRecipientsInterface_DeleteByName_Call) Return(_a0 error) *MockRecipientsInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientsInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockRecipientsInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) Get(ctx context.Context, request sharing.GetRecipientRequest) (*sharing.RecipientInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sharing.RecipientInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetRecipientRequest) (*sharing.RecipientInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.GetRecipientRequest) *sharing.RecipientInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.RecipientInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.GetRecipientRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRecipientsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.GetRecipientRequest
func (_e *MockRecipientsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockRecipientsInterface_Get_Call {
	return &MockRecipientsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockRecipientsInterface_Get_Call) Run(run func(ctx context.Context, request sharing.GetRecipientRequest)) *MockRecipientsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.GetRecipientRequest))
	})
	return _c
}

func (_c *MockRecipientsInterface_Get_Call) Return(_a0 *sharing.RecipientInfo, _a1 error) *MockRecipientsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_Get_Call) RunAndReturn(run func(context.Context, sharing.GetRecipientRequest) (*sharing.RecipientInfo, error)) *MockRecipientsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockRecipientsInterface) GetByName(ctx context.Context, name string) (*sharing.RecipientInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *sharing.RecipientInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sharing.RecipientInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sharing.RecipientInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.RecipientInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockRecipientsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockRecipientsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockRecipientsInterface_GetByName_Call {
	return &MockRecipientsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockRecipientsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockRecipientsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRecipientsInterface_GetByName_Call) Return(_a0 *sharing.RecipientInfo, _a1 error) *MockRecipientsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*sharing.RecipientInfo, error)) *MockRecipientsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockRecipientsInterface) Impl() sharing.RecipientsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 sharing.RecipientsService
	if rf, ok := ret.Get(0).(func() sharing.RecipientsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sharing.RecipientsService)
		}
	}

	return r0
}

// MockRecipientsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockRecipientsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockRecipientsInterface_Expecter) Impl() *MockRecipientsInterface_Impl_Call {
	return &MockRecipientsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockRecipientsInterface_Impl_Call) Run(run func()) *MockRecipientsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRecipientsInterface_Impl_Call) Return(_a0 sharing.RecipientsService) *MockRecipientsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientsInterface_Impl_Call) RunAndReturn(run func() sharing.RecipientsService) *MockRecipientsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) List(ctx context.Context, request sharing.ListRecipientsRequest) *listing.PaginatingIterator[sharing.ListRecipientsRequest, *sharing.ListRecipientsResponse, sharing.RecipientInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *listing.PaginatingIterator[sharing.ListRecipientsRequest, *sharing.ListRecipientsResponse, sharing.RecipientInfo]
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListRecipientsRequest) *listing.PaginatingIterator[sharing.ListRecipientsRequest, *sharing.ListRecipientsResponse, sharing.RecipientInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listing.PaginatingIterator[sharing.ListRecipientsRequest, *sharing.ListRecipientsResponse, sharing.RecipientInfo])
		}
	}

	return r0
}

// MockRecipientsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockRecipientsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.ListRecipientsRequest
func (_e *MockRecipientsInterface_Expecter) List(ctx interface{}, request interface{}) *MockRecipientsInterface_List_Call {
	return &MockRecipientsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockRecipientsInterface_List_Call) Run(run func(ctx context.Context, request sharing.ListRecipientsRequest)) *MockRecipientsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.ListRecipientsRequest))
	})
	return _c
}

func (_c *MockRecipientsInterface_List_Call) Return(_a0 *listing.PaginatingIterator[sharing.ListRecipientsRequest, *sharing.ListRecipientsResponse, sharing.RecipientInfo]) *MockRecipientsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientsInterface_List_Call) RunAndReturn(run func(context.Context, sharing.ListRecipientsRequest) *listing.PaginatingIterator[sharing.ListRecipientsRequest, *sharing.ListRecipientsResponse, sharing.RecipientInfo]) *MockRecipientsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) ListAll(ctx context.Context, request sharing.ListRecipientsRequest) ([]sharing.RecipientInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []sharing.RecipientInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListRecipientsRequest) ([]sharing.RecipientInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListRecipientsRequest) []sharing.RecipientInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sharing.RecipientInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.ListRecipientsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockRecipientsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.ListRecipientsRequest
func (_e *MockRecipientsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockRecipientsInterface_ListAll_Call {
	return &MockRecipientsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockRecipientsInterface_ListAll_Call) Run(run func(ctx context.Context, request sharing.ListRecipientsRequest)) *MockRecipientsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.ListRecipientsRequest))
	})
	return _c
}

func (_c *MockRecipientsInterface_ListAll_Call) Return(_a0 []sharing.RecipientInfo, _a1 error) *MockRecipientsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_ListAll_Call) RunAndReturn(run func(context.Context, sharing.ListRecipientsRequest) ([]sharing.RecipientInfo, error)) *MockRecipientsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// RecipientInfoNameToMetastoreIdMap provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) RecipientInfoNameToMetastoreIdMap(ctx context.Context, request sharing.ListRecipientsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for RecipientInfoNameToMetastoreIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListRecipientsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.ListRecipientsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.ListRecipientsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecipientInfoNameToMetastoreIdMap'
type MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call struct {
	*mock.Call
}

// RecipientInfoNameToMetastoreIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.ListRecipientsRequest
func (_e *MockRecipientsInterface_Expecter) RecipientInfoNameToMetastoreIdMap(ctx interface{}, request interface{}) *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call {
	return &MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call{Call: _e.mock.On("RecipientInfoNameToMetastoreIdMap", ctx, request)}
}

func (_c *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call) Run(run func(ctx context.Context, request sharing.ListRecipientsRequest)) *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.ListRecipientsRequest))
	})
	return _c
}

func (_c *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call) RunAndReturn(run func(context.Context, sharing.ListRecipientsRequest) (map[string]string, error)) *MockRecipientsInterface_RecipientInfoNameToMetastoreIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// RotateToken provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) RotateToken(ctx context.Context, request sharing.RotateRecipientToken) (*sharing.RecipientInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for RotateToken")
	}

	var r0 *sharing.RecipientInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.RotateRecipientToken) (*sharing.RecipientInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.RotateRecipientToken) *sharing.RecipientInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.RecipientInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.RotateRecipientToken) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_RotateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RotateToken'
type MockRecipientsInterface_RotateToken_Call struct {
	*mock.Call
}

// RotateToken is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.RotateRecipientToken
func (_e *MockRecipientsInterface_Expecter) RotateToken(ctx interface{}, request interface{}) *MockRecipientsInterface_RotateToken_Call {
	return &MockRecipientsInterface_RotateToken_Call{Call: _e.mock.On("RotateToken", ctx, request)}
}

func (_c *MockRecipientsInterface_RotateToken_Call) Run(run func(ctx context.Context, request sharing.RotateRecipientToken)) *MockRecipientsInterface_RotateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.RotateRecipientToken))
	})
	return _c
}

func (_c *MockRecipientsInterface_RotateToken_Call) Return(_a0 *sharing.RecipientInfo, _a1 error) *MockRecipientsInterface_RotateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_RotateToken_Call) RunAndReturn(run func(context.Context, sharing.RotateRecipientToken) (*sharing.RecipientInfo, error)) *MockRecipientsInterface_RotateToken_Call {
	_c.Call.Return(run)
	return _c
}

// SharePermissions provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) SharePermissions(ctx context.Context, request sharing.SharePermissionsRequest) (*sharing.GetRecipientSharePermissionsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SharePermissions")
	}

	var r0 *sharing.GetRecipientSharePermissionsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.SharePermissionsRequest) (*sharing.GetRecipientSharePermissionsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sharing.SharePermissionsRequest) *sharing.GetRecipientSharePermissionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.GetRecipientSharePermissionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sharing.SharePermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_SharePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SharePermissions'
type MockRecipientsInterface_SharePermissions_Call struct {
	*mock.Call
}

// SharePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.SharePermissionsRequest
func (_e *MockRecipientsInterface_Expecter) SharePermissions(ctx interface{}, request interface{}) *MockRecipientsInterface_SharePermissions_Call {
	return &MockRecipientsInterface_SharePermissions_Call{Call: _e.mock.On("SharePermissions", ctx, request)}
}

func (_c *MockRecipientsInterface_SharePermissions_Call) Run(run func(ctx context.Context, request sharing.SharePermissionsRequest)) *MockRecipientsInterface_SharePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.SharePermissionsRequest))
	})
	return _c
}

func (_c *MockRecipientsInterface_SharePermissions_Call) Return(_a0 *sharing.GetRecipientSharePermissionsResponse, _a1 error) *MockRecipientsInterface_SharePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_SharePermissions_Call) RunAndReturn(run func(context.Context, sharing.SharePermissionsRequest) (*sharing.GetRecipientSharePermissionsResponse, error)) *MockRecipientsInterface_SharePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// SharePermissionsByName provides a mock function with given fields: ctx, name
func (_m *MockRecipientsInterface) SharePermissionsByName(ctx context.Context, name string) (*sharing.GetRecipientSharePermissionsResponse, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for SharePermissionsByName")
	}

	var r0 *sharing.GetRecipientSharePermissionsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*sharing.GetRecipientSharePermissionsResponse, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *sharing.GetRecipientSharePermissionsResponse); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sharing.GetRecipientSharePermissionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRecipientsInterface_SharePermissionsByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SharePermissionsByName'
type MockRecipientsInterface_SharePermissionsByName_Call struct {
	*mock.Call
}

// SharePermissionsByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockRecipientsInterface_Expecter) SharePermissionsByName(ctx interface{}, name interface{}) *MockRecipientsInterface_SharePermissionsByName_Call {
	return &MockRecipientsInterface_SharePermissionsByName_Call{Call: _e.mock.On("SharePermissionsByName", ctx, name)}
}

func (_c *MockRecipientsInterface_SharePermissionsByName_Call) Run(run func(ctx context.Context, name string)) *MockRecipientsInterface_SharePermissionsByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRecipientsInterface_SharePermissionsByName_Call) Return(_a0 *sharing.GetRecipientSharePermissionsResponse, _a1 error) *MockRecipientsInterface_SharePermissionsByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRecipientsInterface_SharePermissionsByName_Call) RunAndReturn(run func(context.Context, string) (*sharing.GetRecipientSharePermissionsResponse, error)) *MockRecipientsInterface_SharePermissionsByName_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockRecipientsInterface) Update(ctx context.Context, request sharing.UpdateRecipient) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, sharing.UpdateRecipient) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRecipientsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRecipientsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request sharing.UpdateRecipient
func (_e *MockRecipientsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockRecipientsInterface_Update_Call {
	return &MockRecipientsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockRecipientsInterface_Update_Call) Run(run func(ctx context.Context, request sharing.UpdateRecipient)) *MockRecipientsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sharing.UpdateRecipient))
	})
	return _c
}

func (_c *MockRecipientsInterface_Update_Call) Return(_a0 error) *MockRecipientsInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientsInterface_Update_Call) RunAndReturn(run func(context.Context, sharing.UpdateRecipient) error) *MockRecipientsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockRecipientsInterface) WithImpl(impl sharing.RecipientsService) sharing.RecipientsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 sharing.RecipientsInterface
	if rf, ok := ret.Get(0).(func(sharing.RecipientsService) sharing.RecipientsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sharing.RecipientsInterface)
		}
	}

	return r0
}

// MockRecipientsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockRecipientsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl sharing.RecipientsService
func (_e *MockRecipientsInterface_Expecter) WithImpl(impl interface{}) *MockRecipientsInterface_WithImpl_Call {
	return &MockRecipientsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockRecipientsInterface_WithImpl_Call) Run(run func(impl sharing.RecipientsService)) *MockRecipientsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sharing.RecipientsService))
	})
	return _c
}

func (_c *MockRecipientsInterface_WithImpl_Call) Return(_a0 sharing.RecipientsInterface) *MockRecipientsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRecipientsInterface_WithImpl_Call) RunAndReturn(run func(sharing.RecipientsService) sharing.RecipientsInterface) *MockRecipientsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRecipientsInterface creates a new instance of MockRecipientsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRecipientsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRecipientsInterface {
	mock := &MockRecipientsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
