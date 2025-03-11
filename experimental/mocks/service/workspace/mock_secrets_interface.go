// Code generated by mockery v2.53.2. DO NOT EDIT.

package workspace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	workspace "github.com/databricks/databricks-sdk-go/service/workspace"
)

// MockSecretsInterface is an autogenerated mock type for the SecretsInterface type
type MockSecretsInterface struct {
	mock.Mock
}

type MockSecretsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSecretsInterface) EXPECT() *MockSecretsInterface_Expecter {
	return &MockSecretsInterface_Expecter{mock: &_m.Mock}
}

// CreateScope provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) CreateScope(ctx context.Context, request workspace.CreateScope) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.CreateScope) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_CreateScope_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateScope'
type MockSecretsInterface_CreateScope_Call struct {
	*mock.Call
}

// CreateScope is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.CreateScope
func (_e *MockSecretsInterface_Expecter) CreateScope(ctx interface{}, request interface{}) *MockSecretsInterface_CreateScope_Call {
	return &MockSecretsInterface_CreateScope_Call{Call: _e.mock.On("CreateScope", ctx, request)}
}

func (_c *MockSecretsInterface_CreateScope_Call) Run(run func(ctx context.Context, request workspace.CreateScope)) *MockSecretsInterface_CreateScope_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.CreateScope))
	})
	return _c
}

func (_c *MockSecretsInterface_CreateScope_Call) Return(_a0 error) *MockSecretsInterface_CreateScope_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_CreateScope_Call) RunAndReturn(run func(context.Context, workspace.CreateScope) error) *MockSecretsInterface_CreateScope_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAcl provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) DeleteAcl(ctx context.Context, request workspace.DeleteAcl) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAcl")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.DeleteAcl) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_DeleteAcl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAcl'
type MockSecretsInterface_DeleteAcl_Call struct {
	*mock.Call
}

// DeleteAcl is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.DeleteAcl
func (_e *MockSecretsInterface_Expecter) DeleteAcl(ctx interface{}, request interface{}) *MockSecretsInterface_DeleteAcl_Call {
	return &MockSecretsInterface_DeleteAcl_Call{Call: _e.mock.On("DeleteAcl", ctx, request)}
}

func (_c *MockSecretsInterface_DeleteAcl_Call) Run(run func(ctx context.Context, request workspace.DeleteAcl)) *MockSecretsInterface_DeleteAcl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.DeleteAcl))
	})
	return _c
}

func (_c *MockSecretsInterface_DeleteAcl_Call) Return(_a0 error) *MockSecretsInterface_DeleteAcl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_DeleteAcl_Call) RunAndReturn(run func(context.Context, workspace.DeleteAcl) error) *MockSecretsInterface_DeleteAcl_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteScope provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) DeleteScope(ctx context.Context, request workspace.DeleteScope) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.DeleteScope) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_DeleteScope_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteScope'
type MockSecretsInterface_DeleteScope_Call struct {
	*mock.Call
}

// DeleteScope is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.DeleteScope
func (_e *MockSecretsInterface_Expecter) DeleteScope(ctx interface{}, request interface{}) *MockSecretsInterface_DeleteScope_Call {
	return &MockSecretsInterface_DeleteScope_Call{Call: _e.mock.On("DeleteScope", ctx, request)}
}

func (_c *MockSecretsInterface_DeleteScope_Call) Run(run func(ctx context.Context, request workspace.DeleteScope)) *MockSecretsInterface_DeleteScope_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.DeleteScope))
	})
	return _c
}

func (_c *MockSecretsInterface_DeleteScope_Call) Return(_a0 error) *MockSecretsInterface_DeleteScope_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_DeleteScope_Call) RunAndReturn(run func(context.Context, workspace.DeleteScope) error) *MockSecretsInterface_DeleteScope_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteScopeByScope provides a mock function with given fields: ctx, scope
func (_m *MockSecretsInterface) DeleteScopeByScope(ctx context.Context, scope string) error {
	ret := _m.Called(ctx, scope)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScopeByScope")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, scope)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_DeleteScopeByScope_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteScopeByScope'
type MockSecretsInterface_DeleteScopeByScope_Call struct {
	*mock.Call
}

// DeleteScopeByScope is a helper method to define mock.On call
//   - ctx context.Context
//   - scope string
func (_e *MockSecretsInterface_Expecter) DeleteScopeByScope(ctx interface{}, scope interface{}) *MockSecretsInterface_DeleteScopeByScope_Call {
	return &MockSecretsInterface_DeleteScopeByScope_Call{Call: _e.mock.On("DeleteScopeByScope", ctx, scope)}
}

func (_c *MockSecretsInterface_DeleteScopeByScope_Call) Run(run func(ctx context.Context, scope string)) *MockSecretsInterface_DeleteScopeByScope_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSecretsInterface_DeleteScopeByScope_Call) Return(_a0 error) *MockSecretsInterface_DeleteScopeByScope_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_DeleteScopeByScope_Call) RunAndReturn(run func(context.Context, string) error) *MockSecretsInterface_DeleteScopeByScope_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSecret provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) DeleteSecret(ctx context.Context, request workspace.DeleteSecret) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSecret")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.DeleteSecret) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_DeleteSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSecret'
type MockSecretsInterface_DeleteSecret_Call struct {
	*mock.Call
}

// DeleteSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.DeleteSecret
func (_e *MockSecretsInterface_Expecter) DeleteSecret(ctx interface{}, request interface{}) *MockSecretsInterface_DeleteSecret_Call {
	return &MockSecretsInterface_DeleteSecret_Call{Call: _e.mock.On("DeleteSecret", ctx, request)}
}

func (_c *MockSecretsInterface_DeleteSecret_Call) Run(run func(ctx context.Context, request workspace.DeleteSecret)) *MockSecretsInterface_DeleteSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.DeleteSecret))
	})
	return _c
}

func (_c *MockSecretsInterface_DeleteSecret_Call) Return(_a0 error) *MockSecretsInterface_DeleteSecret_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_DeleteSecret_Call) RunAndReturn(run func(context.Context, workspace.DeleteSecret) error) *MockSecretsInterface_DeleteSecret_Call {
	_c.Call.Return(run)
	return _c
}

// GetAcl provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) GetAcl(ctx context.Context, request workspace.GetAclRequest) (*workspace.AclItem, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetAcl")
	}

	var r0 *workspace.AclItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetAclRequest) (*workspace.AclItem, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetAclRequest) *workspace.AclItem); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.AclItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.GetAclRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_GetAcl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAcl'
type MockSecretsInterface_GetAcl_Call struct {
	*mock.Call
}

// GetAcl is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.GetAclRequest
func (_e *MockSecretsInterface_Expecter) GetAcl(ctx interface{}, request interface{}) *MockSecretsInterface_GetAcl_Call {
	return &MockSecretsInterface_GetAcl_Call{Call: _e.mock.On("GetAcl", ctx, request)}
}

func (_c *MockSecretsInterface_GetAcl_Call) Run(run func(ctx context.Context, request workspace.GetAclRequest)) *MockSecretsInterface_GetAcl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.GetAclRequest))
	})
	return _c
}

func (_c *MockSecretsInterface_GetAcl_Call) Return(_a0 *workspace.AclItem, _a1 error) *MockSecretsInterface_GetAcl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_GetAcl_Call) RunAndReturn(run func(context.Context, workspace.GetAclRequest) (*workspace.AclItem, error)) *MockSecretsInterface_GetAcl_Call {
	_c.Call.Return(run)
	return _c
}

// GetSecret provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) GetSecret(ctx context.Context, request workspace.GetSecretRequest) (*workspace.GetSecretResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetSecret")
	}

	var r0 *workspace.GetSecretResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetSecretRequest) (*workspace.GetSecretResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetSecretRequest) *workspace.GetSecretResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.GetSecretResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.GetSecretRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_GetSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecret'
type MockSecretsInterface_GetSecret_Call struct {
	*mock.Call
}

// GetSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.GetSecretRequest
func (_e *MockSecretsInterface_Expecter) GetSecret(ctx interface{}, request interface{}) *MockSecretsInterface_GetSecret_Call {
	return &MockSecretsInterface_GetSecret_Call{Call: _e.mock.On("GetSecret", ctx, request)}
}

func (_c *MockSecretsInterface_GetSecret_Call) Run(run func(ctx context.Context, request workspace.GetSecretRequest)) *MockSecretsInterface_GetSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.GetSecretRequest))
	})
	return _c
}

func (_c *MockSecretsInterface_GetSecret_Call) Return(_a0 *workspace.GetSecretResponse, _a1 error) *MockSecretsInterface_GetSecret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_GetSecret_Call) RunAndReturn(run func(context.Context, workspace.GetSecretRequest) (*workspace.GetSecretResponse, error)) *MockSecretsInterface_GetSecret_Call {
	_c.Call.Return(run)
	return _c
}

// ListAcls provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) ListAcls(ctx context.Context, request workspace.ListAclsRequest) listing.Iterator[workspace.AclItem] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAcls")
	}

	var r0 listing.Iterator[workspace.AclItem]
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListAclsRequest) listing.Iterator[workspace.AclItem]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[workspace.AclItem])
		}
	}

	return r0
}

// MockSecretsInterface_ListAcls_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAcls'
type MockSecretsInterface_ListAcls_Call struct {
	*mock.Call
}

// ListAcls is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListAclsRequest
func (_e *MockSecretsInterface_Expecter) ListAcls(ctx interface{}, request interface{}) *MockSecretsInterface_ListAcls_Call {
	return &MockSecretsInterface_ListAcls_Call{Call: _e.mock.On("ListAcls", ctx, request)}
}

func (_c *MockSecretsInterface_ListAcls_Call) Run(run func(ctx context.Context, request workspace.ListAclsRequest)) *MockSecretsInterface_ListAcls_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListAclsRequest))
	})
	return _c
}

func (_c *MockSecretsInterface_ListAcls_Call) Return(_a0 listing.Iterator[workspace.AclItem]) *MockSecretsInterface_ListAcls_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_ListAcls_Call) RunAndReturn(run func(context.Context, workspace.ListAclsRequest) listing.Iterator[workspace.AclItem]) *MockSecretsInterface_ListAcls_Call {
	_c.Call.Return(run)
	return _c
}

// ListAclsAll provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) ListAclsAll(ctx context.Context, request workspace.ListAclsRequest) ([]workspace.AclItem, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAclsAll")
	}

	var r0 []workspace.AclItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListAclsRequest) ([]workspace.AclItem, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListAclsRequest) []workspace.AclItem); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]workspace.AclItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.ListAclsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_ListAclsAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAclsAll'
type MockSecretsInterface_ListAclsAll_Call struct {
	*mock.Call
}

// ListAclsAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListAclsRequest
func (_e *MockSecretsInterface_Expecter) ListAclsAll(ctx interface{}, request interface{}) *MockSecretsInterface_ListAclsAll_Call {
	return &MockSecretsInterface_ListAclsAll_Call{Call: _e.mock.On("ListAclsAll", ctx, request)}
}

func (_c *MockSecretsInterface_ListAclsAll_Call) Run(run func(ctx context.Context, request workspace.ListAclsRequest)) *MockSecretsInterface_ListAclsAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListAclsRequest))
	})
	return _c
}

func (_c *MockSecretsInterface_ListAclsAll_Call) Return(_a0 []workspace.AclItem, _a1 error) *MockSecretsInterface_ListAclsAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_ListAclsAll_Call) RunAndReturn(run func(context.Context, workspace.ListAclsRequest) ([]workspace.AclItem, error)) *MockSecretsInterface_ListAclsAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListAclsByScope provides a mock function with given fields: ctx, scope
func (_m *MockSecretsInterface) ListAclsByScope(ctx context.Context, scope string) (*workspace.ListAclsResponse, error) {
	ret := _m.Called(ctx, scope)

	if len(ret) == 0 {
		panic("no return value specified for ListAclsByScope")
	}

	var r0 *workspace.ListAclsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*workspace.ListAclsResponse, error)); ok {
		return rf(ctx, scope)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *workspace.ListAclsResponse); ok {
		r0 = rf(ctx, scope)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.ListAclsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, scope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_ListAclsByScope_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAclsByScope'
type MockSecretsInterface_ListAclsByScope_Call struct {
	*mock.Call
}

// ListAclsByScope is a helper method to define mock.On call
//   - ctx context.Context
//   - scope string
func (_e *MockSecretsInterface_Expecter) ListAclsByScope(ctx interface{}, scope interface{}) *MockSecretsInterface_ListAclsByScope_Call {
	return &MockSecretsInterface_ListAclsByScope_Call{Call: _e.mock.On("ListAclsByScope", ctx, scope)}
}

func (_c *MockSecretsInterface_ListAclsByScope_Call) Run(run func(ctx context.Context, scope string)) *MockSecretsInterface_ListAclsByScope_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSecretsInterface_ListAclsByScope_Call) Return(_a0 *workspace.ListAclsResponse, _a1 error) *MockSecretsInterface_ListAclsByScope_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_ListAclsByScope_Call) RunAndReturn(run func(context.Context, string) (*workspace.ListAclsResponse, error)) *MockSecretsInterface_ListAclsByScope_Call {
	_c.Call.Return(run)
	return _c
}

// ListScopes provides a mock function with given fields: ctx
func (_m *MockSecretsInterface) ListScopes(ctx context.Context) listing.Iterator[workspace.SecretScope] {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListScopes")
	}

	var r0 listing.Iterator[workspace.SecretScope]
	if rf, ok := ret.Get(0).(func(context.Context) listing.Iterator[workspace.SecretScope]); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[workspace.SecretScope])
		}
	}

	return r0
}

// MockSecretsInterface_ListScopes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListScopes'
type MockSecretsInterface_ListScopes_Call struct {
	*mock.Call
}

// ListScopes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSecretsInterface_Expecter) ListScopes(ctx interface{}) *MockSecretsInterface_ListScopes_Call {
	return &MockSecretsInterface_ListScopes_Call{Call: _e.mock.On("ListScopes", ctx)}
}

func (_c *MockSecretsInterface_ListScopes_Call) Run(run func(ctx context.Context)) *MockSecretsInterface_ListScopes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSecretsInterface_ListScopes_Call) Return(_a0 listing.Iterator[workspace.SecretScope]) *MockSecretsInterface_ListScopes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_ListScopes_Call) RunAndReturn(run func(context.Context) listing.Iterator[workspace.SecretScope]) *MockSecretsInterface_ListScopes_Call {
	_c.Call.Return(run)
	return _c
}

// ListScopesAll provides a mock function with given fields: ctx
func (_m *MockSecretsInterface) ListScopesAll(ctx context.Context) ([]workspace.SecretScope, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListScopesAll")
	}

	var r0 []workspace.SecretScope
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]workspace.SecretScope, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []workspace.SecretScope); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]workspace.SecretScope)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_ListScopesAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListScopesAll'
type MockSecretsInterface_ListScopesAll_Call struct {
	*mock.Call
}

// ListScopesAll is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSecretsInterface_Expecter) ListScopesAll(ctx interface{}) *MockSecretsInterface_ListScopesAll_Call {
	return &MockSecretsInterface_ListScopesAll_Call{Call: _e.mock.On("ListScopesAll", ctx)}
}

func (_c *MockSecretsInterface_ListScopesAll_Call) Run(run func(ctx context.Context)) *MockSecretsInterface_ListScopesAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSecretsInterface_ListScopesAll_Call) Return(_a0 []workspace.SecretScope, _a1 error) *MockSecretsInterface_ListScopesAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_ListScopesAll_Call) RunAndReturn(run func(context.Context) ([]workspace.SecretScope, error)) *MockSecretsInterface_ListScopesAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListSecrets provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) ListSecrets(ctx context.Context, request workspace.ListSecretsRequest) listing.Iterator[workspace.SecretMetadata] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListSecrets")
	}

	var r0 listing.Iterator[workspace.SecretMetadata]
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListSecretsRequest) listing.Iterator[workspace.SecretMetadata]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[workspace.SecretMetadata])
		}
	}

	return r0
}

// MockSecretsInterface_ListSecrets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSecrets'
type MockSecretsInterface_ListSecrets_Call struct {
	*mock.Call
}

// ListSecrets is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListSecretsRequest
func (_e *MockSecretsInterface_Expecter) ListSecrets(ctx interface{}, request interface{}) *MockSecretsInterface_ListSecrets_Call {
	return &MockSecretsInterface_ListSecrets_Call{Call: _e.mock.On("ListSecrets", ctx, request)}
}

func (_c *MockSecretsInterface_ListSecrets_Call) Run(run func(ctx context.Context, request workspace.ListSecretsRequest)) *MockSecretsInterface_ListSecrets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListSecretsRequest))
	})
	return _c
}

func (_c *MockSecretsInterface_ListSecrets_Call) Return(_a0 listing.Iterator[workspace.SecretMetadata]) *MockSecretsInterface_ListSecrets_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_ListSecrets_Call) RunAndReturn(run func(context.Context, workspace.ListSecretsRequest) listing.Iterator[workspace.SecretMetadata]) *MockSecretsInterface_ListSecrets_Call {
	_c.Call.Return(run)
	return _c
}

// ListSecretsAll provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) ListSecretsAll(ctx context.Context, request workspace.ListSecretsRequest) ([]workspace.SecretMetadata, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListSecretsAll")
	}

	var r0 []workspace.SecretMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListSecretsRequest) ([]workspace.SecretMetadata, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListSecretsRequest) []workspace.SecretMetadata); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]workspace.SecretMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.ListSecretsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_ListSecretsAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSecretsAll'
type MockSecretsInterface_ListSecretsAll_Call struct {
	*mock.Call
}

// ListSecretsAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListSecretsRequest
func (_e *MockSecretsInterface_Expecter) ListSecretsAll(ctx interface{}, request interface{}) *MockSecretsInterface_ListSecretsAll_Call {
	return &MockSecretsInterface_ListSecretsAll_Call{Call: _e.mock.On("ListSecretsAll", ctx, request)}
}

func (_c *MockSecretsInterface_ListSecretsAll_Call) Run(run func(ctx context.Context, request workspace.ListSecretsRequest)) *MockSecretsInterface_ListSecretsAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListSecretsRequest))
	})
	return _c
}

func (_c *MockSecretsInterface_ListSecretsAll_Call) Return(_a0 []workspace.SecretMetadata, _a1 error) *MockSecretsInterface_ListSecretsAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_ListSecretsAll_Call) RunAndReturn(run func(context.Context, workspace.ListSecretsRequest) ([]workspace.SecretMetadata, error)) *MockSecretsInterface_ListSecretsAll_Call {
	_c.Call.Return(run)
	return _c
}

// ListSecretsByScope provides a mock function with given fields: ctx, scope
func (_m *MockSecretsInterface) ListSecretsByScope(ctx context.Context, scope string) (*workspace.ListSecretsResponse, error) {
	ret := _m.Called(ctx, scope)

	if len(ret) == 0 {
		panic("no return value specified for ListSecretsByScope")
	}

	var r0 *workspace.ListSecretsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*workspace.ListSecretsResponse, error)); ok {
		return rf(ctx, scope)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *workspace.ListSecretsResponse); ok {
		r0 = rf(ctx, scope)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.ListSecretsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, scope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSecretsInterface_ListSecretsByScope_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSecretsByScope'
type MockSecretsInterface_ListSecretsByScope_Call struct {
	*mock.Call
}

// ListSecretsByScope is a helper method to define mock.On call
//   - ctx context.Context
//   - scope string
func (_e *MockSecretsInterface_Expecter) ListSecretsByScope(ctx interface{}, scope interface{}) *MockSecretsInterface_ListSecretsByScope_Call {
	return &MockSecretsInterface_ListSecretsByScope_Call{Call: _e.mock.On("ListSecretsByScope", ctx, scope)}
}

func (_c *MockSecretsInterface_ListSecretsByScope_Call) Run(run func(ctx context.Context, scope string)) *MockSecretsInterface_ListSecretsByScope_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSecretsInterface_ListSecretsByScope_Call) Return(_a0 *workspace.ListSecretsResponse, _a1 error) *MockSecretsInterface_ListSecretsByScope_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSecretsInterface_ListSecretsByScope_Call) RunAndReturn(run func(context.Context, string) (*workspace.ListSecretsResponse, error)) *MockSecretsInterface_ListSecretsByScope_Call {
	_c.Call.Return(run)
	return _c
}

// PutAcl provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) PutAcl(ctx context.Context, request workspace.PutAcl) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for PutAcl")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.PutAcl) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_PutAcl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutAcl'
type MockSecretsInterface_PutAcl_Call struct {
	*mock.Call
}

// PutAcl is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.PutAcl
func (_e *MockSecretsInterface_Expecter) PutAcl(ctx interface{}, request interface{}) *MockSecretsInterface_PutAcl_Call {
	return &MockSecretsInterface_PutAcl_Call{Call: _e.mock.On("PutAcl", ctx, request)}
}

func (_c *MockSecretsInterface_PutAcl_Call) Run(run func(ctx context.Context, request workspace.PutAcl)) *MockSecretsInterface_PutAcl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.PutAcl))
	})
	return _c
}

func (_c *MockSecretsInterface_PutAcl_Call) Return(_a0 error) *MockSecretsInterface_PutAcl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_PutAcl_Call) RunAndReturn(run func(context.Context, workspace.PutAcl) error) *MockSecretsInterface_PutAcl_Call {
	_c.Call.Return(run)
	return _c
}

// PutSecret provides a mock function with given fields: ctx, request
func (_m *MockSecretsInterface) PutSecret(ctx context.Context, request workspace.PutSecret) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for PutSecret")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.PutSecret) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSecretsInterface_PutSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutSecret'
type MockSecretsInterface_PutSecret_Call struct {
	*mock.Call
}

// PutSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.PutSecret
func (_e *MockSecretsInterface_Expecter) PutSecret(ctx interface{}, request interface{}) *MockSecretsInterface_PutSecret_Call {
	return &MockSecretsInterface_PutSecret_Call{Call: _e.mock.On("PutSecret", ctx, request)}
}

func (_c *MockSecretsInterface_PutSecret_Call) Run(run func(ctx context.Context, request workspace.PutSecret)) *MockSecretsInterface_PutSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.PutSecret))
	})
	return _c
}

func (_c *MockSecretsInterface_PutSecret_Call) Return(_a0 error) *MockSecretsInterface_PutSecret_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSecretsInterface_PutSecret_Call) RunAndReturn(run func(context.Context, workspace.PutSecret) error) *MockSecretsInterface_PutSecret_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSecretsInterface creates a new instance of MockSecretsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSecretsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSecretsInterface {
	mock := &MockSecretsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
