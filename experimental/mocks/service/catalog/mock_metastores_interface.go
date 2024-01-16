// Code generated by mockery v2.39.1. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockMetastoresInterface is an autogenerated mock type for the MetastoresInterface type
type MockMetastoresInterface struct {
	mock.Mock
}

type MockMetastoresInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetastoresInterface) EXPECT() *MockMetastoresInterface_Expecter {
	return &MockMetastoresInterface_Expecter{mock: &_m.Mock}
}

// Assign provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) Assign(ctx context.Context, request catalog.CreateMetastoreAssignment) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Assign")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateMetastoreAssignment) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetastoresInterface_Assign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Assign'
type MockMetastoresInterface_Assign_Call struct {
	*mock.Call
}

// Assign is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateMetastoreAssignment
func (_e *MockMetastoresInterface_Expecter) Assign(ctx interface{}, request interface{}) *MockMetastoresInterface_Assign_Call {
	return &MockMetastoresInterface_Assign_Call{Call: _e.mock.On("Assign", ctx, request)}
}

func (_c *MockMetastoresInterface_Assign_Call) Run(run func(ctx context.Context, request catalog.CreateMetastoreAssignment)) *MockMetastoresInterface_Assign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateMetastoreAssignment))
	})
	return _c
}

func (_c *MockMetastoresInterface_Assign_Call) Return(_a0 error) *MockMetastoresInterface_Assign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_Assign_Call) RunAndReturn(run func(context.Context, catalog.CreateMetastoreAssignment) error) *MockMetastoresInterface_Assign_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) Create(ctx context.Context, request catalog.CreateMetastore) (*catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateMetastore) (*catalog.MetastoreInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.CreateMetastore) *catalog.MetastoreInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.CreateMetastore) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockMetastoresInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.CreateMetastore
func (_e *MockMetastoresInterface_Expecter) Create(ctx interface{}, request interface{}) *MockMetastoresInterface_Create_Call {
	return &MockMetastoresInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockMetastoresInterface_Create_Call) Run(run func(ctx context.Context, request catalog.CreateMetastore)) *MockMetastoresInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.CreateMetastore))
	})
	return _c
}

func (_c *MockMetastoresInterface_Create_Call) Return(_a0 *catalog.MetastoreInfo, _a1 error) *MockMetastoresInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_Create_Call) RunAndReturn(run func(context.Context, catalog.CreateMetastore) (*catalog.MetastoreInfo, error)) *MockMetastoresInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Current provides a mock function with given fields: ctx
func (_m *MockMetastoresInterface) Current(ctx context.Context) (*catalog.MetastoreAssignment, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Current")
	}

	var r0 *catalog.MetastoreAssignment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*catalog.MetastoreAssignment, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *catalog.MetastoreAssignment); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.MetastoreAssignment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_Current_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Current'
type MockMetastoresInterface_Current_Call struct {
	*mock.Call
}

// Current is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMetastoresInterface_Expecter) Current(ctx interface{}) *MockMetastoresInterface_Current_Call {
	return &MockMetastoresInterface_Current_Call{Call: _e.mock.On("Current", ctx)}
}

func (_c *MockMetastoresInterface_Current_Call) Run(run func(ctx context.Context)) *MockMetastoresInterface_Current_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMetastoresInterface_Current_Call) Return(_a0 *catalog.MetastoreAssignment, _a1 error) *MockMetastoresInterface_Current_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_Current_Call) RunAndReturn(run func(context.Context) (*catalog.MetastoreAssignment, error)) *MockMetastoresInterface_Current_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) Delete(ctx context.Context, request catalog.DeleteMetastoreRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.DeleteMetastoreRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetastoresInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockMetastoresInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.DeleteMetastoreRequest
func (_e *MockMetastoresInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockMetastoresInterface_Delete_Call {
	return &MockMetastoresInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockMetastoresInterface_Delete_Call) Run(run func(ctx context.Context, request catalog.DeleteMetastoreRequest)) *MockMetastoresInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.DeleteMetastoreRequest))
	})
	return _c
}

func (_c *MockMetastoresInterface_Delete_Call) Return(_a0 error) *MockMetastoresInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_Delete_Call) RunAndReturn(run func(context.Context, catalog.DeleteMetastoreRequest) error) *MockMetastoresInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockMetastoresInterface) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetastoresInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockMetastoresInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockMetastoresInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockMetastoresInterface_DeleteById_Call {
	return &MockMetastoresInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockMetastoresInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockMetastoresInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockMetastoresInterface_DeleteById_Call) Return(_a0 error) *MockMetastoresInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockMetastoresInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) Get(ctx context.Context, request catalog.GetMetastoreRequest) (*catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetMetastoreRequest) (*catalog.MetastoreInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetMetastoreRequest) *catalog.MetastoreInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetMetastoreRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockMetastoresInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetMetastoreRequest
func (_e *MockMetastoresInterface_Expecter) Get(ctx interface{}, request interface{}) *MockMetastoresInterface_Get_Call {
	return &MockMetastoresInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockMetastoresInterface_Get_Call) Run(run func(ctx context.Context, request catalog.GetMetastoreRequest)) *MockMetastoresInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetMetastoreRequest))
	})
	return _c
}

func (_c *MockMetastoresInterface_Get_Call) Return(_a0 *catalog.MetastoreInfo, _a1 error) *MockMetastoresInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_Get_Call) RunAndReturn(run func(context.Context, catalog.GetMetastoreRequest) (*catalog.MetastoreInfo, error)) *MockMetastoresInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockMetastoresInterface) GetById(ctx context.Context, id string) (*catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.MetastoreInfo, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.MetastoreInfo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockMetastoresInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockMetastoresInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockMetastoresInterface_GetById_Call {
	return &MockMetastoresInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockMetastoresInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockMetastoresInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockMetastoresInterface_GetById_Call) Return(_a0 *catalog.MetastoreInfo, _a1 error) *MockMetastoresInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*catalog.MetastoreInfo, error)) *MockMetastoresInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *MockMetastoresInterface) GetByName(ctx context.Context, name string) (*catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*catalog.MetastoreInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *catalog.MetastoreInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type MockMetastoresInterface_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockMetastoresInterface_Expecter) GetByName(ctx interface{}, name interface{}) *MockMetastoresInterface_GetByName_Call {
	return &MockMetastoresInterface_GetByName_Call{Call: _e.mock.On("GetByName", ctx, name)}
}

func (_c *MockMetastoresInterface_GetByName_Call) Run(run func(ctx context.Context, name string)) *MockMetastoresInterface_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockMetastoresInterface_GetByName_Call) Return(_a0 *catalog.MetastoreInfo, _a1 error) *MockMetastoresInterface_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_GetByName_Call) RunAndReturn(run func(context.Context, string) (*catalog.MetastoreInfo, error)) *MockMetastoresInterface_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockMetastoresInterface) Impl() catalog.MetastoresService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 catalog.MetastoresService
	if rf, ok := ret.Get(0).(func() catalog.MetastoresService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.MetastoresService)
		}
	}

	return r0
}

// MockMetastoresInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockMetastoresInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockMetastoresInterface_Expecter) Impl() *MockMetastoresInterface_Impl_Call {
	return &MockMetastoresInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockMetastoresInterface_Impl_Call) Run(run func()) *MockMetastoresInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockMetastoresInterface_Impl_Call) Return(_a0 catalog.MetastoresService) *MockMetastoresInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_Impl_Call) RunAndReturn(run func() catalog.MetastoresService) *MockMetastoresInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockMetastoresInterface) List(ctx context.Context) listing.Iterator[catalog.MetastoreInfo] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[catalog.MetastoreInfo]
	if rf, ok := ret.Get(0).(func(context.Context) listing.Iterator[catalog.MetastoreInfo]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.MetastoreInfo])
		}
	}

	return r0
}

// MockMetastoresInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockMetastoresInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMetastoresInterface_Expecter) List(ctx interface{}) *MockMetastoresInterface_List_Call {
	return &MockMetastoresInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockMetastoresInterface_List_Call) Run(run func(ctx context.Context)) *MockMetastoresInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMetastoresInterface_List_Call) Return(_a0 listing.Iterator[catalog.MetastoreInfo]) *MockMetastoresInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_List_Call) RunAndReturn(run func(context.Context) listing.Iterator[catalog.MetastoreInfo]) *MockMetastoresInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx
func (_m *MockMetastoresInterface) ListAll(ctx context.Context) ([]catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]catalog.MetastoreInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []catalog.MetastoreInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockMetastoresInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMetastoresInterface_Expecter) ListAll(ctx interface{}) *MockMetastoresInterface_ListAll_Call {
	return &MockMetastoresInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx)}
}

func (_c *MockMetastoresInterface_ListAll_Call) Run(run func(ctx context.Context)) *MockMetastoresInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMetastoresInterface_ListAll_Call) Return(_a0 []catalog.MetastoreInfo, _a1 error) *MockMetastoresInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_ListAll_Call) RunAndReturn(run func(context.Context) ([]catalog.MetastoreInfo, error)) *MockMetastoresInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// MetastoreInfoNameToMetastoreIdMap provides a mock function with given fields: ctx
func (_m *MockMetastoresInterface) MetastoreInfoNameToMetastoreIdMap(ctx context.Context) (map[string]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for MetastoreInfoNameToMetastoreIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MetastoreInfoNameToMetastoreIdMap'
type MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call struct {
	*mock.Call
}

// MetastoreInfoNameToMetastoreIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMetastoresInterface_Expecter) MetastoreInfoNameToMetastoreIdMap(ctx interface{}) *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call {
	return &MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call{Call: _e.mock.On("MetastoreInfoNameToMetastoreIdMap", ctx)}
}

func (_c *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call) Run(run func(ctx context.Context)) *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call) RunAndReturn(run func(context.Context) (map[string]string, error)) *MockMetastoresInterface_MetastoreInfoNameToMetastoreIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// Summary provides a mock function with given fields: ctx
func (_m *MockMetastoresInterface) Summary(ctx context.Context) (*catalog.GetMetastoreSummaryResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Summary")
	}

	var r0 *catalog.GetMetastoreSummaryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*catalog.GetMetastoreSummaryResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *catalog.GetMetastoreSummaryResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.GetMetastoreSummaryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_Summary_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Summary'
type MockMetastoresInterface_Summary_Call struct {
	*mock.Call
}

// Summary is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockMetastoresInterface_Expecter) Summary(ctx interface{}) *MockMetastoresInterface_Summary_Call {
	return &MockMetastoresInterface_Summary_Call{Call: _e.mock.On("Summary", ctx)}
}

func (_c *MockMetastoresInterface_Summary_Call) Run(run func(ctx context.Context)) *MockMetastoresInterface_Summary_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockMetastoresInterface_Summary_Call) Return(_a0 *catalog.GetMetastoreSummaryResponse, _a1 error) *MockMetastoresInterface_Summary_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_Summary_Call) RunAndReturn(run func(context.Context) (*catalog.GetMetastoreSummaryResponse, error)) *MockMetastoresInterface_Summary_Call {
	_c.Call.Return(run)
	return _c
}

// Unassign provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) Unassign(ctx context.Context, request catalog.UnassignRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Unassign")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UnassignRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetastoresInterface_Unassign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unassign'
type MockMetastoresInterface_Unassign_Call struct {
	*mock.Call
}

// Unassign is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UnassignRequest
func (_e *MockMetastoresInterface_Expecter) Unassign(ctx interface{}, request interface{}) *MockMetastoresInterface_Unassign_Call {
	return &MockMetastoresInterface_Unassign_Call{Call: _e.mock.On("Unassign", ctx, request)}
}

func (_c *MockMetastoresInterface_Unassign_Call) Run(run func(ctx context.Context, request catalog.UnassignRequest)) *MockMetastoresInterface_Unassign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UnassignRequest))
	})
	return _c
}

func (_c *MockMetastoresInterface_Unassign_Call) Return(_a0 error) *MockMetastoresInterface_Unassign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_Unassign_Call) RunAndReturn(run func(context.Context, catalog.UnassignRequest) error) *MockMetastoresInterface_Unassign_Call {
	_c.Call.Return(run)
	return _c
}

// UnassignByWorkspaceId provides a mock function with given fields: ctx, workspaceId
func (_m *MockMetastoresInterface) UnassignByWorkspaceId(ctx context.Context, workspaceId int64) error {
	ret := _m.Called(ctx, workspaceId)

	if len(ret) == 0 {
		panic("no return value specified for UnassignByWorkspaceId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, workspaceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetastoresInterface_UnassignByWorkspaceId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnassignByWorkspaceId'
type MockMetastoresInterface_UnassignByWorkspaceId_Call struct {
	*mock.Call
}

// UnassignByWorkspaceId is a helper method to define mock.On call
//   - ctx context.Context
//   - workspaceId int64
func (_e *MockMetastoresInterface_Expecter) UnassignByWorkspaceId(ctx interface{}, workspaceId interface{}) *MockMetastoresInterface_UnassignByWorkspaceId_Call {
	return &MockMetastoresInterface_UnassignByWorkspaceId_Call{Call: _e.mock.On("UnassignByWorkspaceId", ctx, workspaceId)}
}

func (_c *MockMetastoresInterface_UnassignByWorkspaceId_Call) Run(run func(ctx context.Context, workspaceId int64)) *MockMetastoresInterface_UnassignByWorkspaceId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockMetastoresInterface_UnassignByWorkspaceId_Call) Return(_a0 error) *MockMetastoresInterface_UnassignByWorkspaceId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_UnassignByWorkspaceId_Call) RunAndReturn(run func(context.Context, int64) error) *MockMetastoresInterface_UnassignByWorkspaceId_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) Update(ctx context.Context, request catalog.UpdateMetastore) (*catalog.MetastoreInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *catalog.MetastoreInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateMetastore) (*catalog.MetastoreInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateMetastore) *catalog.MetastoreInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.MetastoreInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.UpdateMetastore) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoresInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockMetastoresInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateMetastore
func (_e *MockMetastoresInterface_Expecter) Update(ctx interface{}, request interface{}) *MockMetastoresInterface_Update_Call {
	return &MockMetastoresInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockMetastoresInterface_Update_Call) Run(run func(ctx context.Context, request catalog.UpdateMetastore)) *MockMetastoresInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateMetastore))
	})
	return _c
}

func (_c *MockMetastoresInterface_Update_Call) Return(_a0 *catalog.MetastoreInfo, _a1 error) *MockMetastoresInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoresInterface_Update_Call) RunAndReturn(run func(context.Context, catalog.UpdateMetastore) (*catalog.MetastoreInfo, error)) *MockMetastoresInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAssignment provides a mock function with given fields: ctx, request
func (_m *MockMetastoresInterface) UpdateAssignment(ctx context.Context, request catalog.UpdateMetastoreAssignment) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAssignment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.UpdateMetastoreAssignment) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMetastoresInterface_UpdateAssignment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAssignment'
type MockMetastoresInterface_UpdateAssignment_Call struct {
	*mock.Call
}

// UpdateAssignment is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.UpdateMetastoreAssignment
func (_e *MockMetastoresInterface_Expecter) UpdateAssignment(ctx interface{}, request interface{}) *MockMetastoresInterface_UpdateAssignment_Call {
	return &MockMetastoresInterface_UpdateAssignment_Call{Call: _e.mock.On("UpdateAssignment", ctx, request)}
}

func (_c *MockMetastoresInterface_UpdateAssignment_Call) Run(run func(ctx context.Context, request catalog.UpdateMetastoreAssignment)) *MockMetastoresInterface_UpdateAssignment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.UpdateMetastoreAssignment))
	})
	return _c
}

func (_c *MockMetastoresInterface_UpdateAssignment_Call) Return(_a0 error) *MockMetastoresInterface_UpdateAssignment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_UpdateAssignment_Call) RunAndReturn(run func(context.Context, catalog.UpdateMetastoreAssignment) error) *MockMetastoresInterface_UpdateAssignment_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockMetastoresInterface) WithImpl(impl catalog.MetastoresService) catalog.MetastoresInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 catalog.MetastoresInterface
	if rf, ok := ret.Get(0).(func(catalog.MetastoresService) catalog.MetastoresInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.MetastoresInterface)
		}
	}

	return r0
}

// MockMetastoresInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockMetastoresInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl catalog.MetastoresService
func (_e *MockMetastoresInterface_Expecter) WithImpl(impl interface{}) *MockMetastoresInterface_WithImpl_Call {
	return &MockMetastoresInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockMetastoresInterface_WithImpl_Call) Run(run func(impl catalog.MetastoresService)) *MockMetastoresInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(catalog.MetastoresService))
	})
	return _c
}

func (_c *MockMetastoresInterface_WithImpl_Call) Return(_a0 catalog.MetastoresInterface) *MockMetastoresInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetastoresInterface_WithImpl_Call) RunAndReturn(run func(catalog.MetastoresService) catalog.MetastoresInterface) *MockMetastoresInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetastoresInterface creates a new instance of MockMetastoresInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetastoresInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetastoresInterface {
	mock := &MockMetastoresInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
