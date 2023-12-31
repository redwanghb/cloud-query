// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/apps/v1 (interfaces: AppsV1Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	rest "k8s.io/client-go/rest"
)

// MockAppsV1Interface is a mock of AppsV1Interface interface.
type MockAppsV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAppsV1InterfaceMockRecorder
}

// MockAppsV1InterfaceMockRecorder is the mock recorder for MockAppsV1Interface.
type MockAppsV1InterfaceMockRecorder struct {
	mock *MockAppsV1Interface
}

// NewMockAppsV1Interface creates a new mock instance.
func NewMockAppsV1Interface(ctrl *gomock.Controller) *MockAppsV1Interface {
	mock := &MockAppsV1Interface{ctrl: ctrl}
	mock.recorder = &MockAppsV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppsV1Interface) EXPECT() *MockAppsV1InterfaceMockRecorder {
	return m.recorder
}

// ControllerRevisions mocks base method.
func (m *MockAppsV1Interface) ControllerRevisions(arg0 string) v1.ControllerRevisionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerRevisions", arg0)
	ret0, _ := ret[0].(v1.ControllerRevisionInterface)
	return ret0
}

// ControllerRevisions indicates an expected call of ControllerRevisions.
func (mr *MockAppsV1InterfaceMockRecorder) ControllerRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerRevisions", reflect.TypeOf((*MockAppsV1Interface)(nil).ControllerRevisions), arg0)
}

// DaemonSets mocks base method.
func (m *MockAppsV1Interface) DaemonSets(arg0 string) v1.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaemonSets", arg0)
	ret0, _ := ret[0].(v1.DaemonSetInterface)
	return ret0
}

// DaemonSets indicates an expected call of DaemonSets.
func (mr *MockAppsV1InterfaceMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSets", reflect.TypeOf((*MockAppsV1Interface)(nil).DaemonSets), arg0)
}

// Deployments mocks base method.
func (m *MockAppsV1Interface) Deployments(arg0 string) v1.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments", arg0)
	ret0, _ := ret[0].(v1.DeploymentInterface)
	return ret0
}

// Deployments indicates an expected call of Deployments.
func (mr *MockAppsV1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockAppsV1Interface)(nil).Deployments), arg0)
}

// RESTClient mocks base method.
func (m *MockAppsV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAppsV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAppsV1Interface)(nil).RESTClient))
}

// ReplicaSets mocks base method.
func (m *MockAppsV1Interface) ReplicaSets(arg0 string) v1.ReplicaSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplicaSets", arg0)
	ret0, _ := ret[0].(v1.ReplicaSetInterface)
	return ret0
}

// ReplicaSets indicates an expected call of ReplicaSets.
func (mr *MockAppsV1InterfaceMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSets", reflect.TypeOf((*MockAppsV1Interface)(nil).ReplicaSets), arg0)
}

// StatefulSets mocks base method.
func (m *MockAppsV1Interface) StatefulSets(arg0 string) v1.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatefulSets", arg0)
	ret0, _ := ret[0].(v1.StatefulSetInterface)
	return ret0
}

// StatefulSets indicates an expected call of StatefulSets.
func (mr *MockAppsV1InterfaceMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatefulSets", reflect.TypeOf((*MockAppsV1Interface)(nil).StatefulSets), arg0)
}
