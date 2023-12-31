// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/core/v1 (interfaces: ResourceQuotasGetter,ResourceQuotaInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// MockResourceQuotasGetter is a mock of ResourceQuotasGetter interface.
type MockResourceQuotasGetter struct {
	ctrl     *gomock.Controller
	recorder *MockResourceQuotasGetterMockRecorder
}

// MockResourceQuotasGetterMockRecorder is the mock recorder for MockResourceQuotasGetter.
type MockResourceQuotasGetterMockRecorder struct {
	mock *MockResourceQuotasGetter
}

// NewMockResourceQuotasGetter creates a new mock instance.
func NewMockResourceQuotasGetter(ctrl *gomock.Controller) *MockResourceQuotasGetter {
	mock := &MockResourceQuotasGetter{ctrl: ctrl}
	mock.recorder = &MockResourceQuotasGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceQuotasGetter) EXPECT() *MockResourceQuotasGetterMockRecorder {
	return m.recorder
}

// ResourceQuotas mocks base method.
func (m *MockResourceQuotasGetter) ResourceQuotas(arg0 string) v12.ResourceQuotaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceQuotas", arg0)
	ret0, _ := ret[0].(v12.ResourceQuotaInterface)
	return ret0
}

// ResourceQuotas indicates an expected call of ResourceQuotas.
func (mr *MockResourceQuotasGetterMockRecorder) ResourceQuotas(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceQuotas", reflect.TypeOf((*MockResourceQuotasGetter)(nil).ResourceQuotas), arg0)
}

// MockResourceQuotaInterface is a mock of ResourceQuotaInterface interface.
type MockResourceQuotaInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceQuotaInterfaceMockRecorder
}

// MockResourceQuotaInterfaceMockRecorder is the mock recorder for MockResourceQuotaInterface.
type MockResourceQuotaInterfaceMockRecorder struct {
	mock *MockResourceQuotaInterface
}

// NewMockResourceQuotaInterface creates a new mock instance.
func NewMockResourceQuotaInterface(ctrl *gomock.Controller) *MockResourceQuotaInterface {
	mock := &MockResourceQuotaInterface{ctrl: ctrl}
	mock.recorder = &MockResourceQuotaInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceQuotaInterface) EXPECT() *MockResourceQuotaInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockResourceQuotaInterface) Apply(arg0 context.Context, arg1 *v11.ResourceQuotaApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockResourceQuotaInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockResourceQuotaInterface) ApplyStatus(arg0 context.Context, arg1 *v11.ResourceQuotaApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockResourceQuotaInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockResourceQuotaInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockResourceQuotaInterface) Create(arg0 context.Context, arg1 *v1.ResourceQuota, arg2 v10.CreateOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockResourceQuotaInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockResourceQuotaInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockResourceQuotaInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockResourceQuotaInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockResourceQuotaInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockResourceQuotaInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockResourceQuotaInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockResourceQuotaInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockResourceQuotaInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ResourceQuotaList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ResourceQuotaList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResourceQuotaInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceQuotaInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockResourceQuotaInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockResourceQuotaInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockResourceQuotaInterface) Update(arg0 context.Context, arg1 *v1.ResourceQuota, arg2 v10.UpdateOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockResourceQuotaInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockResourceQuotaInterface) UpdateStatus(arg0 context.Context, arg1 *v1.ResourceQuota, arg2 v10.UpdateOptions) (*v1.ResourceQuota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ResourceQuota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockResourceQuotaInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockResourceQuotaInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockResourceQuotaInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockResourceQuotaInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockResourceQuotaInterface)(nil).Watch), arg0, arg1)
}
