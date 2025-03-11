// Code generated by mockery v2.53.2. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockStorageCredentialsInterface is an autogenerated mock type for the StorageCredentialsInterface type
type MockStorageCredentialsInterface struct {
	mock.Mock
}

type MockStorageCredentialsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorageCredentialsInterface) EXPECT() *MockStorageCredentialsInterface_Expecter {
	return &MockStorageCredentialsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) Create(ctx context.Context, request catalog.CreateStorageCredential) (*catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateStorageCredential) (*catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateStorageCredential) *catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateStorageCredential) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockStorageCredentialsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateStorageCredential
func (_e *MockStorageCredentialsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_Create_Call {
	return &MockStorageCredentialsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateStorageCredential)) *MockStorageCredentialsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateStorageCredential))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_Create_Call) Return(_a0 *catalog.StorageCredentialInfo, _a1 error) *MockStorageCredentialsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateStorageCredential) (*catalog.StorageCredentialInfo, error)) *MockStorageCredentialsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) Delete(ctx context.Context, request catalog.DeleteStorageCredentialRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteStorageCredentialRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorageCredentialsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockStorageCredentialsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteStorageCredentialRequest
func (_e *MockStorageCredentialsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_Delete_Call {
	return &MockStorageCredentialsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteStorageCredentialRequest)) *MockStorageCredentialsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteStorageCredentialRequest))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_Delete_Call) Return(_a0 error) *MockStorageCredentialsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorageCredentialsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteStorageCredentialRequest) error) *MockStorageCredentialsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *MockStorageCredentialsInterface) DeleteByName(ctx context.Context, name string) error {
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

// MockStorageCredentialsInterface_DeleteByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByName'
type MockStorageCredentialsInterface_DeleteByName_Call struct {
	*mock.Call
}

// DeleteByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockStorageCredentialsInterface_Expecter) DeleteByName(ctx interface{}, name interface{}) *MockStorageCredentialsInterface_DeleteByName_Call {
	return &MockStorageCredentialsInterface_DeleteByName_Call{Call: _e.mock.On("DeleteByName", ctx, name)}
}

func (_c *MockStorageCredentialsInterface_DeleteByName_Call) Run(run func(ctx context.Context, name string)) *MockStorageCredentialsInterface_DeleteByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_DeleteByName_Call) Return(_a0 error) *MockStorageCredentialsInterface_DeleteByName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorageCredentialsInterface_DeleteByName_Call) RunAndReturn(run func(context.Context, string) error) *MockStorageCredentialsInterface_DeleteByName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) Get(ctx context.Context, request catalog.GetStorageCredentialRequest) (*catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetStorageCredentialRequest) (*catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetStorageCredentialRequest) *catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetStorageCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockStorageCredentialsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetStorageCredentialRequest
func (_e *MockStorageCredentialsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_Get_Call {
	return &MockStorageCredentialsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetStorageCredentialRequest)) *MockStorageCredentialsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetStorageCredentialRequest))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_Get_Call) Return(_a0 *catalog.StorageCredentialInfo, _a1 error) *MockStorageCredentialsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetStorageCredentialRequest) (*catalog.StorageCredentialInfo, error)) *MockStorageCredentialsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockStorageCredentialsInterface) GetByName(ctx context.Context, name string) (*catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockStorageCredentialsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockStorageCredentialsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockStorageCredentialsInterface_GetByName_Call {
	return &MockStorageCredentialsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockStorageCredentialsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockStorageCredentialsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_GetByName_Call) Return(_a0 *catalog.StorageCredentialInfo, _a1 error) *MockStorageCredentialsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.StorageCredentialInfo, error)) *MockStorageCredentialsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) List(ctx context.Context, request catalog.ListStorageCredentialsRequest) listing.Iterator[catalog.StorageCredentialInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.StorageCredentialInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListStorageCredentialsRequest) listing.Iterator[catalog.StorageCredentialInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.StorageCredentialInfo])
		}
	}

	return r0
}

// MockStorageCredentialsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockStorageCredentialsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListStorageCredentialsRequest
func (_e *MockStorageCredentialsInterface_Expecter) List(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_List_Call {
	return &MockStorageCredentialsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListStorageCredentialsRequest)) *MockStorageCredentialsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListStorageCredentialsRequest))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_List_Call) Return(_a0 listing.Iterator[catalog.StorageCredentialInfo]) *MockStorageCredentialsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorageCredentialsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListStorageCredentialsRequest) listing.Iterator[catalog.StorageCredentialInfo]) *MockStorageCredentialsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) ListAll(ctx context.Context, request catalog.ListStorageCredentialsRequest) ([]catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListStorageCredentialsRequest) ([]catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListStorageCredentialsRequest) []catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListStorageCredentialsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockStorageCredentialsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListStorageCredentialsRequest
func (_e *MockStorageCredentialsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_ListAll_Call {
	return &MockStorageCredentialsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_ListAll_Call) Run(run func(ctx context.Context, request catalog.ListStorageCredentialsRequest)) *MockStorageCredentialsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListStorageCredentialsRequest))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_ListAll_Call) Return(_a0 []catalog.StorageCredentialInfo, _a1 error) *MockStorageCredentialsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_ListAll_Call) RunAndReturn(run func(context.Context, catalog.ListStorageCredentialsRequest) ([]catalog.StorageCredentialInfo, error)) *MockStorageCredentialsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// StorageCredentialInfoNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) StorageCredentialInfoNameToIdMap(ctx context.Context, request catalog.ListStorageCredentialsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for StorageCredentialInfoNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListStorageCredentialsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListStorageCredentialsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListStorageCredentialsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StorageCredentialInfoNameToIdMap'
type MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call struct {
	*mock.Call
}

// StorageCredentialInfoNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListStorageCredentialsRequest
func (_e *MockStorageCredentialsInterface_Expecter) StorageCredentialInfoNameToIdMap(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	return &MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call{Call: _e.mock.On("StorageCredentialInfoNameToIdMap", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call) Run(run func(ctx context.Context, request catalog.ListStorageCredentialsRequest)) *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListStorageCredentialsRequest))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call) RunAndReturn(run func(context.Context, catalog.ListStorageCredentialsRequest) (map[string]string, error)) *MockStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) Update(ctx context.Context, request catalog.UpdateStorageCredential) (*catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateStorageCredential) (*catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateStorageCredential) *catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateStorageCredential) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockStorageCredentialsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateStorageCredential
func (_e *MockStorageCredentialsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_Update_Call {
	return &MockStorageCredentialsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateStorageCredential)) *MockStorageCredentialsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateStorageCredential))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_Update_Call) Return(_a0 *catalog.StorageCredentialInfo, _a1 error) *MockStorageCredentialsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateStorageCredential) (*catalog.StorageCredentialInfo, error)) *MockStorageCredentialsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields: ctx, request
func (_m *MockStorageCredentialsInterface) Validate(ctx context.Context, request catalog.ValidateStorageCredential) (*catalog.ValidateStorageCredentialResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 *catalog.ValidateStorageCredentialResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ValidateStorageCredential) (*catalog.ValidateStorageCredentialResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ValidateStorageCredential) *catalog.ValidateStorageCredentialResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.ValidateStorageCredentialResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ValidateStorageCredential) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageCredentialsInterface_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockStorageCredentialsInterface_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ValidateStorageCredential
func (_e *MockStorageCredentialsInterface_Expecter) Validate(ctx interface{}, request interface{}) *MockStorageCredentialsInterface_Validate_Call {
	return &MockStorageCredentialsInterface_Validate_Call{Call: _e.mock.On("Validate", ctx, request)}
}

func (_c *MockStorageCredentialsInterface_Validate_Call) Run(run func(ctx context.Context, request catalog.ValidateStorageCredential)) *MockStorageCredentialsInterface_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ValidateStorageCredential))
	})
	return _c
}

func (_c *MockStorageCredentialsInterface_Validate_Call) Return(_a0 *catalog.ValidateStorageCredentialResponse, _a1 error) *MockStorageCredentialsInterface_Validate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageCredentialsInterface_Validate_Call) RunAndReturn(run func(context.Context, catalog.ValidateStorageCredential) (*catalog.ValidateStorageCredentialResponse, error)) *MockStorageCredentialsInterface_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorageCredentialsInterface creates a new instance of MockStorageCredentialsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorageCredentialsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorageCredentialsInterface {
	mock := &MockStorageCredentialsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
