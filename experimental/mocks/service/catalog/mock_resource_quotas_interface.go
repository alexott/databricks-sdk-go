// Code generated by mockery v2.43.0. DO NOT EDIT.

package catalog

import (
	context "context"

	catalog "github.com/databricks/databricks-sdk-go/service/catalog"

	listing "github.com/databricks/databricks-sdk-go/listing"

	mock "github.com/stretchr/testify/mock"
)

// MockResourceQuotasInterface is an autogenerated mock type for the ResourceQuotasInterface type
type MockResourceQuotasInterface struct {
	mock.Mock
}

type MockResourceQuotasInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResourceQuotasInterface) EXPECT() *MockResourceQuotasInterface_Expecter {
	return &MockResourceQuotasInterface_Expecter{mock: &_m.Mock}
}

// GetQuota provides a mock function with given fields: ctx, request
func (_m *MockResourceQuotasInterface) GetQuota(ctx context.Context, request catalog.GetQuotaRequest) (*catalog.GetQuotaResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetQuota")
	}

	var r0 *catalog.GetQuotaResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetQuotaRequest) (*catalog.GetQuotaResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.GetQuotaRequest) *catalog.GetQuotaResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.GetQuotaResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.GetQuotaRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResourceQuotasInterface_GetQuota_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetQuota'
type MockResourceQuotasInterface_GetQuota_Call struct {
	*mock.Call
}

// GetQuota is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.GetQuotaRequest
func (_e *MockResourceQuotasInterface_Expecter) GetQuota(ctx interface{}, request interface{}) *MockResourceQuotasInterface_GetQuota_Call {
	return &MockResourceQuotasInterface_GetQuota_Call{Call: _e.mock.On("GetQuota", ctx, request)}
}

func (_c *MockResourceQuotasInterface_GetQuota_Call) Run(run func(ctx context.Context, request catalog.GetQuotaRequest)) *MockResourceQuotasInterface_GetQuota_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.GetQuotaRequest))
	})
	return _c
}

func (_c *MockResourceQuotasInterface_GetQuota_Call) Return(_a0 *catalog.GetQuotaResponse, _a1 error) *MockResourceQuotasInterface_GetQuota_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResourceQuotasInterface_GetQuota_Call) RunAndReturn(run func(context.Context, catalog.GetQuotaRequest) (*catalog.GetQuotaResponse, error)) *MockResourceQuotasInterface_GetQuota_Call {
	_c.Call.Return(run)
	return _c
}

// GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName provides a mock function with given fields: ctx, parentSecurableType, parentFullName, quotaName
func (_m *MockResourceQuotasInterface) GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName(ctx context.Context, parentSecurableType string, parentFullName string, quotaName string) (*catalog.GetQuotaResponse, error) {
	ret := _m.Called(ctx, parentSecurableType, parentFullName, quotaName)

	if len(ret) == 0 {
		panic("no return value specified for GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName")
	}

	var r0 *catalog.GetQuotaResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*catalog.GetQuotaResponse, error)); ok {
		return rf(ctx, parentSecurableType, parentFullName, quotaName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *catalog.GetQuotaResponse); ok {
		r0 = rf(ctx, parentSecurableType, parentFullName, quotaName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*catalog.GetQuotaResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, parentSecurableType, parentFullName, quotaName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName'
type MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call struct {
	*mock.Call
}

// GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName is a helper method to define mock.On call
//   - ctx context.Context
//   - parentSecurableType string
//   - parentFullName string
//   - quotaName string
func (_e *MockResourceQuotasInterface_Expecter) GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName(ctx interface{}, parentSecurableType interface{}, parentFullName interface{}, quotaName interface{}) *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call {
	return &MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call{Call: _e.mock.On("GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName", ctx, parentSecurableType, parentFullName, quotaName)}
}

func (_c *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call) Run(run func(ctx context.Context, parentSecurableType string, parentFullName string, quotaName string)) *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call) Return(_a0 *catalog.GetQuotaResponse, _a1 error) *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call) RunAndReturn(run func(context.Context, string, string, string) (*catalog.GetQuotaResponse, error)) *MockResourceQuotasInterface_GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName_Call {
	_c.Call.Return(run)
	return _c
}

// ListQuotas provides a mock function with given fields: ctx, request
func (_m *MockResourceQuotasInterface) ListQuotas(ctx context.Context, request catalog.ListQuotasRequest) listing.Iterator[catalog.QuotaInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListQuotas")
	}

	var r0 listing.Iterator[catalog.QuotaInfo]
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListQuotasRequest) listing.Iterator[catalog.QuotaInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[catalog.QuotaInfo])
		}
	}

	return r0
}

// MockResourceQuotasInterface_ListQuotas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListQuotas'
type MockResourceQuotasInterface_ListQuotas_Call struct {
	*mock.Call
}

// ListQuotas is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListQuotasRequest
func (_e *MockResourceQuotasInterface_Expecter) ListQuotas(ctx interface{}, request interface{}) *MockResourceQuotasInterface_ListQuotas_Call {
	return &MockResourceQuotasInterface_ListQuotas_Call{Call: _e.mock.On("ListQuotas", ctx, request)}
}

func (_c *MockResourceQuotasInterface_ListQuotas_Call) Run(run func(ctx context.Context, request catalog.ListQuotasRequest)) *MockResourceQuotasInterface_ListQuotas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListQuotasRequest))
	})
	return _c
}

func (_c *MockResourceQuotasInterface_ListQuotas_Call) Return(_a0 listing.Iterator[catalog.QuotaInfo]) *MockResourceQuotasInterface_ListQuotas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockResourceQuotasInterface_ListQuotas_Call) RunAndReturn(run func(context.Context, catalog.ListQuotasRequest) listing.Iterator[catalog.QuotaInfo]) *MockResourceQuotasInterface_ListQuotas_Call {
	_c.Call.Return(run)
	return _c
}

// ListQuotasAll provides a mock function with given fields: ctx, request
func (_m *MockResourceQuotasInterface) ListQuotasAll(ctx context.Context, request catalog.ListQuotasRequest) ([]catalog.QuotaInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListQuotasAll")
	}

	var r0 []catalog.QuotaInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListQuotasRequest) ([]catalog.QuotaInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, catalog.ListQuotasRequest) []catalog.QuotaInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]catalog.QuotaInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, catalog.ListQuotasRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResourceQuotasInterface_ListQuotasAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListQuotasAll'
type MockResourceQuotasInterface_ListQuotasAll_Call struct {
	*mock.Call
}

// ListQuotasAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request catalog.ListQuotasRequest
func (_e *MockResourceQuotasInterface_Expecter) ListQuotasAll(ctx interface{}, request interface{}) *MockResourceQuotasInterface_ListQuotasAll_Call {
	return &MockResourceQuotasInterface_ListQuotasAll_Call{Call: _e.mock.On("ListQuotasAll", ctx, request)}
}

func (_c *MockResourceQuotasInterface_ListQuotasAll_Call) Run(run func(ctx context.Context, request catalog.ListQuotasRequest)) *MockResourceQuotasInterface_ListQuotasAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(catalog.ListQuotasRequest))
	})
	return _c
}

func (_c *MockResourceQuotasInterface_ListQuotasAll_Call) Return(_a0 []catalog.QuotaInfo, _a1 error) *MockResourceQuotasInterface_ListQuotasAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResourceQuotasInterface_ListQuotasAll_Call) RunAndReturn(run func(context.Context, catalog.ListQuotasRequest) ([]catalog.QuotaInfo, error)) *MockResourceQuotasInterface_ListQuotasAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockResourceQuotasInterface creates a new instance of MockResourceQuotasInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResourceQuotasInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResourceQuotasInterface {
	mock := &MockResourceQuotasInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}