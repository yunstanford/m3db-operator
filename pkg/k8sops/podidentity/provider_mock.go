// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db-operator/pkg/k8sops/podidentity/provider.go

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package podidentity is a generated GoMock package.
package podidentity

import (
	"reflect"

	"github.com/m3db/m3db-operator/pkg/apis/m3dboperator/v1alpha1"

	"github.com/golang/mock/gomock"
	"k8s.io/api/core/v1"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// Identity mocks base method
func (m *MockProvider) Identity(pod *v1.Pod, cluster *v1alpha1.M3DBCluster) (*v1alpha1.PodIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identity", pod, cluster)
	ret0, _ := ret[0].(*v1alpha1.PodIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identity indicates an expected call of Identity
func (mr *MockProviderMockRecorder) Identity(pod, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockProvider)(nil).Identity), pod, cluster)
}
