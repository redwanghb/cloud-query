// Code generated by MockGen. DO NOT EDIT.
// Source: hosts_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadogV1 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	gomock "github.com/golang/mock/gomock"
)

// MockHostsAPIClient is a mock of HostsAPIClient interface.
type MockHostsAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockHostsAPIClientMockRecorder
}

// MockHostsAPIClientMockRecorder is the mock recorder for MockHostsAPIClient.
type MockHostsAPIClientMockRecorder struct {
	mock *MockHostsAPIClient
}

// NewMockHostsAPIClient creates a new mock instance.
func NewMockHostsAPIClient(ctrl *gomock.Controller) *MockHostsAPIClient {
	mock := &MockHostsAPIClient{ctrl: ctrl}
	mock.recorder = &MockHostsAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostsAPIClient) EXPECT() *MockHostsAPIClientMockRecorder {
	return m.recorder
}

// ListHosts mocks base method.
func (m *MockHostsAPIClient) ListHosts(arg0 context.Context, arg1 ...datadogV1.ListHostsOptionalParameters) (datadogV1.HostListResponse, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHosts", varargs...)
	ret0, _ := ret[0].(datadogV1.HostListResponse)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListHosts indicates an expected call of ListHosts.
func (mr *MockHostsAPIClientMockRecorder) ListHosts(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHosts", reflect.TypeOf((*MockHostsAPIClient)(nil).ListHosts), varargs...)
}
