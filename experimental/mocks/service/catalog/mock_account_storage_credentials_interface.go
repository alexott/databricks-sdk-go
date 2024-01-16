// Code generated by mockery v2.39.1. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	mock "github.com/stretchr/testify/mock"
)

// MockAccountStorageCredentialsInterface is an autogenerated mock type for the AccountStorageCredentialsInterface type
type MockAccountStorageCredentialsInterface struct {
	mock.Mock
}

type MockAccountStorageCredentialsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountStorageCredentialsInterface) EXPECT() *MockAccountStorageCredentialsInterface_Expecter {
	return &MockAccountStorageCredentialsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockAccountStorageCredentialsInterface) Create(ctx context.Context, request catalog.AccountsCreateStorageCredential) (*catalog.AccountsStorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.AccountsStorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsCreateStorageCredential) (*catalog.AccountsStorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsCreateStorageCredential) *catalog.AccountsStorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsStorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.AccountsCreateStorageCredential) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountStorageCredentialsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.AccountsCreateStorageCredential
func (_e *MockAccountStorageCredentialsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockAccountStorageCredentialsInterface_Create_Call {
	return &MockAccountStorageCredentialsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockAccountStorageCredentialsInterface_Create_Call) Run(run func(ctx context.Context, request catalog.AccountsCreateStorageCredential)) *MockAccountStorageCredentialsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.AccountsCreateStorageCredential))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Create_Call) Return(_a0 *catalog.AccountsStorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.AccountsCreateStorageCredential) (*catalog.AccountsStorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockAccountStorageCredentialsInterface) Delete(ctx context.Context, request catalog.DeleteAccountStorageCredentialRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteAccountStorageCredentialRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountStorageCredentialsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockAccountStorageCredentialsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteAccountStorageCredentialRequest
func (_e *MockAccountStorageCredentialsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockAccountStorageCredentialsInterface_Delete_Call {
	return &MockAccountStorageCredentialsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockAccountStorageCredentialsInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteAccountStorageCredentialRequest)) *MockAccountStorageCredentialsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteAccountStorageCredentialRequest))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Delete_Call) Return(_a0 error) *MockAccountStorageCredentialsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteAccountStorageCredentialRequest) error) *MockAccountStorageCredentialsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByMetastoreIdAndStorageCredentialName provides a mock function with given fields: ctx, metastoreId, storageCredentialName
func (_m *MockAccountStorageCredentialsInterface) DeleteByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) error {
	ret := _m.Called(ctx, metastoreId, storageCredentialName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByMetastoreIdAndStorageCredentialName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, metastoreId, storageCredentialName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByMetastoreIdAndStorageCredentialName'
type MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call struct {
	*mock.Call
}

// DeleteByMetastoreIdAndStorageCredentialName is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
//   - storageCredentialName string
func (_e *MockAccountStorageCredentialsInterface_Expecter) DeleteByMetastoreIdAndStorageCredentialName(ctx interface{}, metastoreId interface{}, storageCredentialName interface{}) *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call {
	return &MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call{Call: _e.mock.On("DeleteByMetastoreIdAndStorageCredentialName", ctx, metastoreId, storageCredentialName)}
}

func (_c *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call) Run(run func(ctx context.Context, metastoreId string, storageCredentialName string)) *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call) Return(_a0 error) *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call) RunAndReturn(run func(context.Context, string, string) error) *MockAccountStorageCredentialsInterface_DeleteByMetastoreIdAndStorageCredentialName_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockAccountStorageCredentialsInterface) Get(ctx context.Context, request catalog.GetAccountStorageCredentialRequest) (*catalog.AccountsStorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.AccountsStorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetAccountStorageCredentialRequest) (*catalog.AccountsStorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetAccountStorageCredentialRequest) *catalog.AccountsStorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsStorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetAccountStorageCredentialRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockAccountStorageCredentialsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetAccountStorageCredentialRequest
func (_e *MockAccountStorageCredentialsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockAccountStorageCredentialsInterface_Get_Call {
	return &MockAccountStorageCredentialsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockAccountStorageCredentialsInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetAccountStorageCredentialRequest)) *MockAccountStorageCredentialsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetAccountStorageCredentialRequest))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Get_Call) Return(_a0 *catalog.AccountsStorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetAccountStorageCredentialRequest) (*catalog.AccountsStorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByMetastoreIdAndStorageCredentialName provides a mock function with given fields: ctx, metastoreId, storageCredentialName
func (_m *MockAccountStorageCredentialsInterface) GetByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) (*catalog.AccountsStorageCredentialInfo, error) {
	ret := _m.Called(ctx, metastoreId, storageCredentialName)

	if len(ret) == 0 {
		panic("no return value specified for GetByMetastoreIdAndStorageCredentialName")
	}

	var r0 *catalog.AccountsStorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*catalog.AccountsStorageCredentialInfo, error)); ok {
		return rf(ctx, metastoreId, storageCredentialName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *catalog.AccountsStorageCredentialInfo); ok {
		r0 = rf(ctx, metastoreId, storageCredentialName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsStorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, metastoreId, storageCredentialName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByMetastoreIdAndStorageCredentialName'
type MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call struct {
	*mock.Call
}

// GetByMetastoreIdAndStorageCredentialName is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
//   - storageCredentialName string
func (_e *MockAccountStorageCredentialsInterface_Expecter) GetByMetastoreIdAndStorageCredentialName(ctx interface{}, metastoreId interface{}, storageCredentialName interface{}) *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call {
	return &MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call{Call: _e.mock.On("GetByMetastoreIdAndStorageCredentialName", ctx, metastoreId, storageCredentialName)}
}

func (_c *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call) Run(run func(ctx context.Context, metastoreId string, storageCredentialName string)) *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call) Return(_a0 *catalog.AccountsStorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call) RunAndReturn(run func(context.Context, string, string) (*catalog.AccountsStorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_GetByMetastoreIdAndStorageCredentialName_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockAccountStorageCredentialsInterface) GetByName(ctx context.Context, name string) (*catalog.StorageCredentialInfo, error) {
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

// MockAccountStorageCredentialsInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockAccountStorageCredentialsInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockAccountStorageCredentialsInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockAccountStorageCredentialsInterface_GetByName_Call {
	return &MockAccountStorageCredentialsInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockAccountStorageCredentialsInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockAccountStorageCredentialsInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_GetByName_Call) Return(_a0 *catalog.StorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.StorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockAccountStorageCredentialsInterface) Impl() catalog.AccountStorageCredentialsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.AccountStorageCredentialsService
	if rf, ok := ret.Get(0).(func() catalog.AccountStorageCredentialsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AccountStorageCredentialsService)
		}
	}

	return r0
}

// MockAccountStorageCredentialsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockAccountStorageCredentialsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockAccountStorageCredentialsInterface_Expecter) Impl() *MockAccountStorageCredentialsInterface_Impl_Call {
	return &MockAccountStorageCredentialsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockAccountStorageCredentialsInterface_Impl_Call) Run(run func()) *MockAccountStorageCredentialsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Impl_Call) Return(_a0 catalog.AccountStorageCredentialsService) *MockAccountStorageCredentialsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Impl_Call) RunAndReturn(run func() catalog.AccountStorageCredentialsService) *MockAccountStorageCredentialsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockAccountStorageCredentialsInterface) List(ctx context.Context, request catalog.ListAccountStorageCredentialsRequest) ([]catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountStorageCredentialsRequest) ([]catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountStorageCredentialsRequest) []catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListAccountStorageCredentialsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockAccountStorageCredentialsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListAccountStorageCredentialsRequest
func (_e *MockAccountStorageCredentialsInterface_Expecter) List(ctx interface{}, request interface{}) *MockAccountStorageCredentialsInterface_List_Call {
	return &MockAccountStorageCredentialsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockAccountStorageCredentialsInterface_List_Call) Run(run func(ctx context.Context, request catalog.ListAccountStorageCredentialsRequest)) *MockAccountStorageCredentialsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListAccountStorageCredentialsRequest))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_List_Call) Return(_a0 []catalog.StorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_List_Call) RunAndReturn(run func(context.Context, catalog.ListAccountStorageCredentialsRequest) ([]catalog.StorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListByMetastoreId provides a mock function with given fields: ctx, metastoreId
func (_m *MockAccountStorageCredentialsInterface) ListByMetastoreId(ctx context.Context, metastoreId string) ([]catalog.StorageCredentialInfo, error) {
	ret := _m.Called(ctx, metastoreId)

	if len(ret) == 0 {
		panic("no return value specified for ListByMetastoreId")
	}

	var r0 []catalog.StorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]catalog.StorageCredentialInfo, error)); ok {
		return rf(ctx, metastoreId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []catalog.StorageCredentialInfo); ok {
		r0 = rf(ctx, metastoreId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.StorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, metastoreId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_ListByMetastoreId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByMetastoreId'
type MockAccountStorageCredentialsInterface_ListByMetastoreId_Call struct {
	*mock.Call
}

// ListByMetastoreId is a helper method to define mock.On call
//   - ctx context.Context
//   - metastoreId string
func (_e *MockAccountStorageCredentialsInterface_Expecter) ListByMetastoreId(ctx interface{}, metastoreId interface{}) *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call {
	return &MockAccountStorageCredentialsInterface_ListByMetastoreId_Call{Call: _e.mock.On("ListByMetastoreId", ctx, metastoreId)}
}

func (_c *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call) Run(run func(ctx context.Context, metastoreId string)) *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call) Return(_a0 []catalog.StorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call) RunAndReturn(run func(context.Context, string) ([]catalog.StorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_ListByMetastoreId_Call {
	_c.Call.Return(run)
	return _c
}

// StorageCredentialInfoNameToIdMap provides a mock function with given fields: ctx, request
func (_m *MockAccountStorageCredentialsInterface) StorageCredentialInfoNameToIdMap(ctx context.Context, request catalog.ListAccountStorageCredentialsRequest) (map[string]string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for StorageCredentialInfoNameToIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountStorageCredentialsRequest) (map[string]string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListAccountStorageCredentialsRequest) map[string]string); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListAccountStorageCredentialsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StorageCredentialInfoNameToIdMap'
type MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call struct {
	*mock.Call
}

// StorageCredentialInfoNameToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListAccountStorageCredentialsRequest
func (_e *MockAccountStorageCredentialsInterface_Expecter) StorageCredentialInfoNameToIdMap(ctx interface{}, request interface{}) *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	return &MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call{Call: _e.mock.On("StorageCredentialInfoNameToIdMap", ctx, request)}
}

func (_c *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call) Run(run func(ctx context.Context, request catalog.ListAccountStorageCredentialsRequest)) *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListAccountStorageCredentialsRequest))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call) RunAndReturn(run func(context.Context, catalog.ListAccountStorageCredentialsRequest) (map[string]string, error)) *MockAccountStorageCredentialsInterface_StorageCredentialInfoNameToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockAccountStorageCredentialsInterface) Update(ctx context.Context, request catalog.AccountsUpdateStorageCredential) (*catalog.AccountsStorageCredentialInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.AccountsStorageCredentialInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsUpdateStorageCredential) (*catalog.AccountsStorageCredentialInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.AccountsUpdateStorageCredential) *catalog.AccountsStorageCredentialInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.AccountsStorageCredentialInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.AccountsUpdateStorageCredential) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountStorageCredentialsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountStorageCredentialsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.AccountsUpdateStorageCredential
func (_e *MockAccountStorageCredentialsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockAccountStorageCredentialsInterface_Update_Call {
	return &MockAccountStorageCredentialsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockAccountStorageCredentialsInterface_Update_Call) Run(run func(ctx context.Context, request catalog.AccountsUpdateStorageCredential)) *MockAccountStorageCredentialsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.AccountsUpdateStorageCredential))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Update_Call) Return(_a0 *catalog.AccountsStorageCredentialInfo, _a1 error) *MockAccountStorageCredentialsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.AccountsUpdateStorageCredential) (*catalog.AccountsStorageCredentialInfo, error)) *MockAccountStorageCredentialsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockAccountStorageCredentialsInterface) WithImpl(impl catalog.AccountStorageCredentialsService) catalog.AccountStorageCredentialsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.AccountStorageCredentialsInterface
	if rf, ok := ret.Get(0).(func(catalog.AccountStorageCredentialsService) catalog.AccountStorageCredentialsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AccountStorageCredentialsInterface)
		}
	}

	return r0
}

// MockAccountStorageCredentialsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockAccountStorageCredentialsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.AccountStorageCredentialsService
func (_e *MockAccountStorageCredentialsInterface_Expecter) WithImpl(impl interface{}) *MockAccountStorageCredentialsInterface_WithImpl_Call {
	return &MockAccountStorageCredentialsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockAccountStorageCredentialsInterface_WithImpl_Call) Run(run func(impl catalog.AccountStorageCredentialsService)) *MockAccountStorageCredentialsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.AccountStorageCredentialsService))
	})
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_WithImpl_Call) Return(_a0 catalog.AccountStorageCredentialsInterface) *MockAccountStorageCredentialsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountStorageCredentialsInterface_WithImpl_Call) RunAndReturn(run func(catalog.AccountStorageCredentialsService) catalog.AccountStorageCredentialsInterface) *MockAccountStorageCredentialsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountStorageCredentialsInterface creates a new instance of MockAccountStorageCredentialsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountStorageCredentialsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountStorageCredentialsInterface {
	mock := &MockAccountStorageCredentialsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}