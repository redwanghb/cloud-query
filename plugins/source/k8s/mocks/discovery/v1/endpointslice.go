// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/discovery/v1 (interfaces: EndpointSlicesGetter,EndpointSliceInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/discovery/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/discovery/v1"
	v12 "k8s.io/client-go/kubernetes/typed/discovery/v1"
)

// MockEndpointSlicesGetter is a mock of EndpointSlicesGetter interface.
type MockEndpointSlicesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointSlicesGetterMockRecorder
}

// MockEndpointSlicesGetterMockRecorder is the mock recorder for MockEndpointSlicesGetter.
type MockEndpointSlicesGetterMockRecorder struct {
	mock *MockEndpointSlicesGetter
}

// NewMockEndpointSlicesGetter creates a new mock instance.
func NewMockEndpointSlicesGetter(ctrl *gomock.Controller) *MockEndpointSlicesGetter {
	mock := &MockEndpointSlicesGetter{ctrl: ctrl}
	mock.recorder = &MockEndpointSlicesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpointSlicesGetter) EXPECT() *MockEndpointSlicesGetterMockRecorder {
	return m.recorder
}

// EndpointSlices mocks base method.
func (m *MockEndpointSlicesGetter) EndpointSlices(arg0 string) v12.EndpointSliceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointSlices", arg0)
	ret0, _ := ret[0].(v12.EndpointSliceInterface)
	return ret0
}

// EndpointSlices indicates an expected call of EndpointSlices.
func (mr *MockEndpointSlicesGetterMockRecorder) EndpointSlices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointSlices", reflect.TypeOf((*MockEndpointSlicesGetter)(nil).EndpointSlices), arg0)
}

// MockEndpointSliceInterface is a mock of EndpointSliceInterface interface.
type MockEndpointSliceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointSliceInterfaceMockRecorder
}

// MockEndpointSliceInterfaceMockRecorder is the mock recorder for MockEndpointSliceInterface.
type MockEndpointSliceInterfaceMockRecorder struct {
	mock *MockEndpointSliceInterface
}

// NewMockEndpointSliceInterface creates a new mock instance.
func NewMockEndpointSliceInterface(ctrl *gomock.Controller) *MockEndpointSliceInterface {
	mock := &MockEndpointSliceInterface{ctrl: ctrl}
	mock.recorder = &MockEndpointSliceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpointSliceInterface) EXPECT() *MockEndpointSliceInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockEndpointSliceInterface) Apply(arg0 context.Context, arg1 *v11.EndpointSliceApplyConfiguration, arg2 v10.ApplyOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockEndpointSliceInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockEndpointSliceInterface) Create(arg0 context.Context, arg1 *v1.EndpointSlice, arg2 v10.CreateOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEndpointSliceInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockEndpointSliceInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEndpointSliceInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockEndpointSliceInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockEndpointSliceInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockEndpointSliceInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockEndpointSliceInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEndpointSliceInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockEndpointSliceInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.EndpointSliceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.EndpointSliceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockEndpointSliceInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEndpointSliceInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockEndpointSliceInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockEndpointSliceInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockEndpointSliceInterface) Update(arg0 context.Context, arg1 *v1.EndpointSlice, arg2 v10.UpdateOptions) (*v1.EndpointSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.EndpointSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEndpointSliceInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockEndpointSliceInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockEndpointSliceInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockEndpointSliceInterface)(nil).Watch), arg0, arg1)
}
