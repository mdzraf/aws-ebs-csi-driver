// Copyright 2024 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the 'License');
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an 'AS IS' BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes (interfaces: Interface)

// Package driver is a generated GoMock package.
package driver

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	discovery "k8s.io/client-go/discovery"
	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	v1alpha10 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	v10 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1beta10 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	v11 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	v1alpha11 "k8s.io/client-go/kubernetes/typed/authentication/v1alpha1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	v12 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	v1beta12 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	v13 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	v2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	v2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	v14 "k8s.io/client-go/kubernetes/typed/batch/v1"
	v1beta13 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	v15 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	v1alpha12 "k8s.io/client-go/kubernetes/typed/certificates/v1alpha1"
	v1beta14 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	v16 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	v1alpha13 "k8s.io/client-go/kubernetes/typed/coordination/v1alpha1"
	v1beta15 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	v17 "k8s.io/client-go/kubernetes/typed/core/v1"
	v18 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	v1beta16 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	v19 "k8s.io/client-go/kubernetes/typed/events/v1"
	v1beta17 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	v1beta18 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	v110 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1"
	v1beta19 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	v1beta20 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	v1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
	v111 "k8s.io/client-go/kubernetes/typed/networking/v1"
	v1alpha14 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	v1beta110 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	v112 "k8s.io/client-go/kubernetes/typed/node/v1"
	v1alpha15 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	v1beta111 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	v113 "k8s.io/client-go/kubernetes/typed/policy/v1"
	v1beta112 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	v114 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	v1alpha16 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	v1beta113 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	v1alpha3 "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
	v115 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	v1alpha17 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	v1beta114 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	v116 "k8s.io/client-go/kubernetes/typed/storage/v1"
	v1alpha18 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	v1beta115 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	v1alpha19 "k8s.io/client-go/kubernetes/typed/storagemigration/v1alpha1"
)

// MockKubernetesClient is a mock of Interface interface.
type MockKubernetesClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesClientMockRecorder
}

// MockKubernetesClientMockRecorder is the mock recorder for MockKubernetesClient.
type MockKubernetesClientMockRecorder struct {
	mock *MockKubernetesClient
}

// NewMockKubernetesClient creates a new mock instance.
func NewMockKubernetesClient(ctrl *gomock.Controller) *MockKubernetesClient {
	mock := &MockKubernetesClient{ctrl: ctrl}
	mock.recorder = &MockKubernetesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubernetesClient) EXPECT() *MockKubernetesClientMockRecorder {
	return m.recorder
}

// AdmissionregistrationV1 mocks base method.
func (m *MockKubernetesClient) AdmissionregistrationV1() v1.AdmissionregistrationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmissionregistrationV1")
	ret0, _ := ret[0].(v1.AdmissionregistrationV1Interface)
	return ret0
}

// AdmissionregistrationV1 indicates an expected call of AdmissionregistrationV1.
func (mr *MockKubernetesClientMockRecorder) AdmissionregistrationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmissionregistrationV1", reflect.TypeOf((*MockKubernetesClient)(nil).AdmissionregistrationV1))
}

// AdmissionregistrationV1alpha1 mocks base method.
func (m *MockKubernetesClient) AdmissionregistrationV1alpha1() v1alpha1.AdmissionregistrationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmissionregistrationV1alpha1")
	ret0, _ := ret[0].(v1alpha1.AdmissionregistrationV1alpha1Interface)
	return ret0
}

// AdmissionregistrationV1alpha1 indicates an expected call of AdmissionregistrationV1alpha1.
func (mr *MockKubernetesClientMockRecorder) AdmissionregistrationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmissionregistrationV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).AdmissionregistrationV1alpha1))
}

// AdmissionregistrationV1beta1 mocks base method.
func (m *MockKubernetesClient) AdmissionregistrationV1beta1() v1beta1.AdmissionregistrationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmissionregistrationV1beta1")
	ret0, _ := ret[0].(v1beta1.AdmissionregistrationV1beta1Interface)
	return ret0
}

// AdmissionregistrationV1beta1 indicates an expected call of AdmissionregistrationV1beta1.
func (mr *MockKubernetesClientMockRecorder) AdmissionregistrationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmissionregistrationV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).AdmissionregistrationV1beta1))
}

// AppsV1 mocks base method.
func (m *MockKubernetesClient) AppsV1() v10.AppsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1")
	ret0, _ := ret[0].(v10.AppsV1Interface)
	return ret0
}

// AppsV1 indicates an expected call of AppsV1.
func (mr *MockKubernetesClientMockRecorder) AppsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1", reflect.TypeOf((*MockKubernetesClient)(nil).AppsV1))
}

// AppsV1beta1 mocks base method.
func (m *MockKubernetesClient) AppsV1beta1() v1beta10.AppsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta1")
	ret0, _ := ret[0].(v1beta10.AppsV1beta1Interface)
	return ret0
}

// AppsV1beta1 indicates an expected call of AppsV1beta1.
func (mr *MockKubernetesClientMockRecorder) AppsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).AppsV1beta1))
}

// AppsV1beta2 mocks base method.
func (m *MockKubernetesClient) AppsV1beta2() v1beta2.AppsV1beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta2")
	ret0, _ := ret[0].(v1beta2.AppsV1beta2Interface)
	return ret0
}

// AppsV1beta2 indicates an expected call of AppsV1beta2.
func (mr *MockKubernetesClientMockRecorder) AppsV1beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta2", reflect.TypeOf((*MockKubernetesClient)(nil).AppsV1beta2))
}

// AuthenticationV1 mocks base method.
func (m *MockKubernetesClient) AuthenticationV1() v11.AuthenticationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1")
	ret0, _ := ret[0].(v11.AuthenticationV1Interface)
	return ret0
}

// AuthenticationV1 indicates an expected call of AuthenticationV1.
func (mr *MockKubernetesClientMockRecorder) AuthenticationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1", reflect.TypeOf((*MockKubernetesClient)(nil).AuthenticationV1))
}

// AuthenticationV1alpha1 mocks base method.
func (m *MockKubernetesClient) AuthenticationV1alpha1() v1alpha11.AuthenticationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1alpha1")
	ret0, _ := ret[0].(v1alpha11.AuthenticationV1alpha1Interface)
	return ret0
}

// AuthenticationV1alpha1 indicates an expected call of AuthenticationV1alpha1.
func (mr *MockKubernetesClientMockRecorder) AuthenticationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).AuthenticationV1alpha1))
}

// AuthenticationV1beta1 mocks base method.
func (m *MockKubernetesClient) AuthenticationV1beta1() v1beta11.AuthenticationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1beta1")
	ret0, _ := ret[0].(v1beta11.AuthenticationV1beta1Interface)
	return ret0
}

// AuthenticationV1beta1 indicates an expected call of AuthenticationV1beta1.
func (mr *MockKubernetesClientMockRecorder) AuthenticationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).AuthenticationV1beta1))
}

// AuthorizationV1 mocks base method.
func (m *MockKubernetesClient) AuthorizationV1() v12.AuthorizationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV1")
	ret0, _ := ret[0].(v12.AuthorizationV1Interface)
	return ret0
}

// AuthorizationV1 indicates an expected call of AuthorizationV1.
func (mr *MockKubernetesClientMockRecorder) AuthorizationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV1", reflect.TypeOf((*MockKubernetesClient)(nil).AuthorizationV1))
}

// AuthorizationV1beta1 mocks base method.
func (m *MockKubernetesClient) AuthorizationV1beta1() v1beta12.AuthorizationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV1beta1")
	ret0, _ := ret[0].(v1beta12.AuthorizationV1beta1Interface)
	return ret0
}

// AuthorizationV1beta1 indicates an expected call of AuthorizationV1beta1.
func (mr *MockKubernetesClientMockRecorder) AuthorizationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).AuthorizationV1beta1))
}

// AutoscalingV1 mocks base method.
func (m *MockKubernetesClient) AutoscalingV1() v13.AutoscalingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV1")
	ret0, _ := ret[0].(v13.AutoscalingV1Interface)
	return ret0
}

// AutoscalingV1 indicates an expected call of AutoscalingV1.
func (mr *MockKubernetesClientMockRecorder) AutoscalingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV1", reflect.TypeOf((*MockKubernetesClient)(nil).AutoscalingV1))
}

// AutoscalingV2 mocks base method.
func (m *MockKubernetesClient) AutoscalingV2() v2.AutoscalingV2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2")
	ret0, _ := ret[0].(v2.AutoscalingV2Interface)
	return ret0
}

// AutoscalingV2 indicates an expected call of AutoscalingV2.
func (mr *MockKubernetesClientMockRecorder) AutoscalingV2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2", reflect.TypeOf((*MockKubernetesClient)(nil).AutoscalingV2))
}

// AutoscalingV2beta1 mocks base method.
func (m *MockKubernetesClient) AutoscalingV2beta1() v2beta1.AutoscalingV2beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta1")
	ret0, _ := ret[0].(v2beta1.AutoscalingV2beta1Interface)
	return ret0
}

// AutoscalingV2beta1 indicates an expected call of AutoscalingV2beta1.
func (mr *MockKubernetesClientMockRecorder) AutoscalingV2beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta1", reflect.TypeOf((*MockKubernetesClient)(nil).AutoscalingV2beta1))
}

// AutoscalingV2beta2 mocks base method.
func (m *MockKubernetesClient) AutoscalingV2beta2() v2beta2.AutoscalingV2beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta2")
	ret0, _ := ret[0].(v2beta2.AutoscalingV2beta2Interface)
	return ret0
}

// AutoscalingV2beta2 indicates an expected call of AutoscalingV2beta2.
func (mr *MockKubernetesClientMockRecorder) AutoscalingV2beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta2", reflect.TypeOf((*MockKubernetesClient)(nil).AutoscalingV2beta2))
}

// BatchV1 mocks base method.
func (m *MockKubernetesClient) BatchV1() v14.BatchV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV1")
	ret0, _ := ret[0].(v14.BatchV1Interface)
	return ret0
}

// BatchV1 indicates an expected call of BatchV1.
func (mr *MockKubernetesClientMockRecorder) BatchV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV1", reflect.TypeOf((*MockKubernetesClient)(nil).BatchV1))
}

// BatchV1beta1 mocks base method.
func (m *MockKubernetesClient) BatchV1beta1() v1beta13.BatchV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV1beta1")
	ret0, _ := ret[0].(v1beta13.BatchV1beta1Interface)
	return ret0
}

// BatchV1beta1 indicates an expected call of BatchV1beta1.
func (mr *MockKubernetesClientMockRecorder) BatchV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).BatchV1beta1))
}

// CertificatesV1 mocks base method.
func (m *MockKubernetesClient) CertificatesV1() v15.CertificatesV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificatesV1")
	ret0, _ := ret[0].(v15.CertificatesV1Interface)
	return ret0
}

// CertificatesV1 indicates an expected call of CertificatesV1.
func (mr *MockKubernetesClientMockRecorder) CertificatesV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificatesV1", reflect.TypeOf((*MockKubernetesClient)(nil).CertificatesV1))
}

// CertificatesV1alpha1 mocks base method.
func (m *MockKubernetesClient) CertificatesV1alpha1() v1alpha12.CertificatesV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificatesV1alpha1")
	ret0, _ := ret[0].(v1alpha12.CertificatesV1alpha1Interface)
	return ret0
}

// CertificatesV1alpha1 indicates an expected call of CertificatesV1alpha1.
func (mr *MockKubernetesClientMockRecorder) CertificatesV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificatesV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).CertificatesV1alpha1))
}

// CertificatesV1beta1 mocks base method.
func (m *MockKubernetesClient) CertificatesV1beta1() v1beta14.CertificatesV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificatesV1beta1")
	ret0, _ := ret[0].(v1beta14.CertificatesV1beta1Interface)
	return ret0
}

// CertificatesV1beta1 indicates an expected call of CertificatesV1beta1.
func (mr *MockKubernetesClientMockRecorder) CertificatesV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificatesV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).CertificatesV1beta1))
}

// CoordinationV1 mocks base method.
func (m *MockKubernetesClient) CoordinationV1() v16.CoordinationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1")
	ret0, _ := ret[0].(v16.CoordinationV1Interface)
	return ret0
}

// CoordinationV1 indicates an expected call of CoordinationV1.
func (mr *MockKubernetesClientMockRecorder) CoordinationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1", reflect.TypeOf((*MockKubernetesClient)(nil).CoordinationV1))
}

// CoordinationV1alpha1 mocks base method.
func (m *MockKubernetesClient) CoordinationV1alpha1() v1alpha13.CoordinationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1alpha1")
	ret0, _ := ret[0].(v1alpha13.CoordinationV1alpha1Interface)
	return ret0
}

// CoordinationV1alpha1 indicates an expected call of CoordinationV1alpha1.
func (mr *MockKubernetesClientMockRecorder) CoordinationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).CoordinationV1alpha1))
}

// CoordinationV1beta1 mocks base method.
func (m *MockKubernetesClient) CoordinationV1beta1() v1beta15.CoordinationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1beta1")
	ret0, _ := ret[0].(v1beta15.CoordinationV1beta1Interface)
	return ret0
}

// CoordinationV1beta1 indicates an expected call of CoordinationV1beta1.
func (mr *MockKubernetesClientMockRecorder) CoordinationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).CoordinationV1beta1))
}

// CoreV1 mocks base method.
func (m *MockKubernetesClient) CoreV1() v17.CoreV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoreV1")
	ret0, _ := ret[0].(v17.CoreV1Interface)
	return ret0
}

// CoreV1 indicates an expected call of CoreV1.
func (mr *MockKubernetesClientMockRecorder) CoreV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoreV1", reflect.TypeOf((*MockKubernetesClient)(nil).CoreV1))
}

// Discovery mocks base method.
func (m *MockKubernetesClient) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

// Discovery indicates an expected call of Discovery.
func (mr *MockKubernetesClientMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discovery", reflect.TypeOf((*MockKubernetesClient)(nil).Discovery))
}

// DiscoveryV1 mocks base method.
func (m *MockKubernetesClient) DiscoveryV1() v18.DiscoveryV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoveryV1")
	ret0, _ := ret[0].(v18.DiscoveryV1Interface)
	return ret0
}

// DiscoveryV1 indicates an expected call of DiscoveryV1.
func (mr *MockKubernetesClientMockRecorder) DiscoveryV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryV1", reflect.TypeOf((*MockKubernetesClient)(nil).DiscoveryV1))
}

// DiscoveryV1beta1 mocks base method.
func (m *MockKubernetesClient) DiscoveryV1beta1() v1beta16.DiscoveryV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoveryV1beta1")
	ret0, _ := ret[0].(v1beta16.DiscoveryV1beta1Interface)
	return ret0
}

// DiscoveryV1beta1 indicates an expected call of DiscoveryV1beta1.
func (mr *MockKubernetesClientMockRecorder) DiscoveryV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).DiscoveryV1beta1))
}

// EventsV1 mocks base method.
func (m *MockKubernetesClient) EventsV1() v19.EventsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsV1")
	ret0, _ := ret[0].(v19.EventsV1Interface)
	return ret0
}

// EventsV1 indicates an expected call of EventsV1.
func (mr *MockKubernetesClientMockRecorder) EventsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsV1", reflect.TypeOf((*MockKubernetesClient)(nil).EventsV1))
}

// EventsV1beta1 mocks base method.
func (m *MockKubernetesClient) EventsV1beta1() v1beta17.EventsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsV1beta1")
	ret0, _ := ret[0].(v1beta17.EventsV1beta1Interface)
	return ret0
}

// EventsV1beta1 indicates an expected call of EventsV1beta1.
func (mr *MockKubernetesClientMockRecorder) EventsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).EventsV1beta1))
}

// ExtensionsV1beta1 mocks base method.
func (m *MockKubernetesClient) ExtensionsV1beta1() v1beta18.ExtensionsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtensionsV1beta1")
	ret0, _ := ret[0].(v1beta18.ExtensionsV1beta1Interface)
	return ret0
}

// ExtensionsV1beta1 indicates an expected call of ExtensionsV1beta1.
func (mr *MockKubernetesClientMockRecorder) ExtensionsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionsV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).ExtensionsV1beta1))
}

// FlowcontrolV1 mocks base method.
func (m *MockKubernetesClient) FlowcontrolV1() v110.FlowcontrolV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1")
	ret0, _ := ret[0].(v110.FlowcontrolV1Interface)
	return ret0
}

// FlowcontrolV1 indicates an expected call of FlowcontrolV1.
func (mr *MockKubernetesClientMockRecorder) FlowcontrolV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1", reflect.TypeOf((*MockKubernetesClient)(nil).FlowcontrolV1))
}

// FlowcontrolV1beta1 mocks base method.
func (m *MockKubernetesClient) FlowcontrolV1beta1() v1beta19.FlowcontrolV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1beta1")
	ret0, _ := ret[0].(v1beta19.FlowcontrolV1beta1Interface)
	return ret0
}

// FlowcontrolV1beta1 indicates an expected call of FlowcontrolV1beta1.
func (mr *MockKubernetesClientMockRecorder) FlowcontrolV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).FlowcontrolV1beta1))
}

// FlowcontrolV1beta2 mocks base method.
func (m *MockKubernetesClient) FlowcontrolV1beta2() v1beta20.FlowcontrolV1beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1beta2")
	ret0, _ := ret[0].(v1beta20.FlowcontrolV1beta2Interface)
	return ret0
}

// FlowcontrolV1beta2 indicates an expected call of FlowcontrolV1beta2.
func (mr *MockKubernetesClientMockRecorder) FlowcontrolV1beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1beta2", reflect.TypeOf((*MockKubernetesClient)(nil).FlowcontrolV1beta2))
}

// FlowcontrolV1beta3 mocks base method.
func (m *MockKubernetesClient) FlowcontrolV1beta3() v1beta3.FlowcontrolV1beta3Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1beta3")
	ret0, _ := ret[0].(v1beta3.FlowcontrolV1beta3Interface)
	return ret0
}

// FlowcontrolV1beta3 indicates an expected call of FlowcontrolV1beta3.
func (mr *MockKubernetesClientMockRecorder) FlowcontrolV1beta3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1beta3", reflect.TypeOf((*MockKubernetesClient)(nil).FlowcontrolV1beta3))
}

// InternalV1alpha1 mocks base method.
func (m *MockKubernetesClient) InternalV1alpha1() v1alpha10.InternalV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InternalV1alpha1")
	ret0, _ := ret[0].(v1alpha10.InternalV1alpha1Interface)
	return ret0
}

// InternalV1alpha1 indicates an expected call of InternalV1alpha1.
func (mr *MockKubernetesClientMockRecorder) InternalV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).InternalV1alpha1))
}

// NetworkingV1 mocks base method.
func (m *MockKubernetesClient) NetworkingV1() v111.NetworkingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1")
	ret0, _ := ret[0].(v111.NetworkingV1Interface)
	return ret0
}

// NetworkingV1 indicates an expected call of NetworkingV1.
func (mr *MockKubernetesClientMockRecorder) NetworkingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1", reflect.TypeOf((*MockKubernetesClient)(nil).NetworkingV1))
}

// NetworkingV1alpha1 mocks base method.
func (m *MockKubernetesClient) NetworkingV1alpha1() v1alpha14.NetworkingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1alpha1")
	ret0, _ := ret[0].(v1alpha14.NetworkingV1alpha1Interface)
	return ret0
}

// NetworkingV1alpha1 indicates an expected call of NetworkingV1alpha1.
func (mr *MockKubernetesClientMockRecorder) NetworkingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).NetworkingV1alpha1))
}

// NetworkingV1beta1 mocks base method.
func (m *MockKubernetesClient) NetworkingV1beta1() v1beta110.NetworkingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1beta1")
	ret0, _ := ret[0].(v1beta110.NetworkingV1beta1Interface)
	return ret0
}

// NetworkingV1beta1 indicates an expected call of NetworkingV1beta1.
func (mr *MockKubernetesClientMockRecorder) NetworkingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).NetworkingV1beta1))
}

// NodeV1 mocks base method.
func (m *MockKubernetesClient) NodeV1() v112.NodeV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1")
	ret0, _ := ret[0].(v112.NodeV1Interface)
	return ret0
}

// NodeV1 indicates an expected call of NodeV1.
func (mr *MockKubernetesClientMockRecorder) NodeV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1", reflect.TypeOf((*MockKubernetesClient)(nil).NodeV1))
}

// NodeV1alpha1 mocks base method.
func (m *MockKubernetesClient) NodeV1alpha1() v1alpha15.NodeV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1alpha1")
	ret0, _ := ret[0].(v1alpha15.NodeV1alpha1Interface)
	return ret0
}

// NodeV1alpha1 indicates an expected call of NodeV1alpha1.
func (mr *MockKubernetesClientMockRecorder) NodeV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).NodeV1alpha1))
}

// NodeV1beta1 mocks base method.
func (m *MockKubernetesClient) NodeV1beta1() v1beta111.NodeV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1beta1")
	ret0, _ := ret[0].(v1beta111.NodeV1beta1Interface)
	return ret0
}

// NodeV1beta1 indicates an expected call of NodeV1beta1.
func (mr *MockKubernetesClientMockRecorder) NodeV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).NodeV1beta1))
}

// PolicyV1 mocks base method.
func (m *MockKubernetesClient) PolicyV1() v113.PolicyV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyV1")
	ret0, _ := ret[0].(v113.PolicyV1Interface)
	return ret0
}

// PolicyV1 indicates an expected call of PolicyV1.
func (mr *MockKubernetesClientMockRecorder) PolicyV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyV1", reflect.TypeOf((*MockKubernetesClient)(nil).PolicyV1))
}

// PolicyV1beta1 mocks base method.
func (m *MockKubernetesClient) PolicyV1beta1() v1beta112.PolicyV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyV1beta1")
	ret0, _ := ret[0].(v1beta112.PolicyV1beta1Interface)
	return ret0
}

// PolicyV1beta1 indicates an expected call of PolicyV1beta1.
func (mr *MockKubernetesClientMockRecorder) PolicyV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).PolicyV1beta1))
}

// RbacV1 mocks base method.
func (m *MockKubernetesClient) RbacV1() v114.RbacV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1")
	ret0, _ := ret[0].(v114.RbacV1Interface)
	return ret0
}

// RbacV1 indicates an expected call of RbacV1.
func (mr *MockKubernetesClientMockRecorder) RbacV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1", reflect.TypeOf((*MockKubernetesClient)(nil).RbacV1))
}

// RbacV1alpha1 mocks base method.
func (m *MockKubernetesClient) RbacV1alpha1() v1alpha16.RbacV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1alpha1")
	ret0, _ := ret[0].(v1alpha16.RbacV1alpha1Interface)
	return ret0
}

// RbacV1alpha1 indicates an expected call of RbacV1alpha1.
func (mr *MockKubernetesClientMockRecorder) RbacV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).RbacV1alpha1))
}

// RbacV1beta1 mocks base method.
func (m *MockKubernetesClient) RbacV1beta1() v1beta113.RbacV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1beta1")
	ret0, _ := ret[0].(v1beta113.RbacV1beta1Interface)
	return ret0
}

// RbacV1beta1 indicates an expected call of RbacV1beta1.
func (mr *MockKubernetesClientMockRecorder) RbacV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).RbacV1beta1))
}

// ResourceV1alpha3 mocks base method.
func (m *MockKubernetesClient) ResourceV1alpha3() v1alpha3.ResourceV1alpha3Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceV1alpha3")
	ret0, _ := ret[0].(v1alpha3.ResourceV1alpha3Interface)
	return ret0
}

// ResourceV1alpha3 indicates an expected call of ResourceV1alpha3.
func (mr *MockKubernetesClientMockRecorder) ResourceV1alpha3() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceV1alpha3", reflect.TypeOf((*MockKubernetesClient)(nil).ResourceV1alpha3))
}

// SchedulingV1 mocks base method.
func (m *MockKubernetesClient) SchedulingV1() v115.SchedulingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1")
	ret0, _ := ret[0].(v115.SchedulingV1Interface)
	return ret0
}

// SchedulingV1 indicates an expected call of SchedulingV1.
func (mr *MockKubernetesClientMockRecorder) SchedulingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1", reflect.TypeOf((*MockKubernetesClient)(nil).SchedulingV1))
}

// SchedulingV1alpha1 mocks base method.
func (m *MockKubernetesClient) SchedulingV1alpha1() v1alpha17.SchedulingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1alpha1")
	ret0, _ := ret[0].(v1alpha17.SchedulingV1alpha1Interface)
	return ret0
}

// SchedulingV1alpha1 indicates an expected call of SchedulingV1alpha1.
func (mr *MockKubernetesClientMockRecorder) SchedulingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).SchedulingV1alpha1))
}

// SchedulingV1beta1 mocks base method.
func (m *MockKubernetesClient) SchedulingV1beta1() v1beta114.SchedulingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1beta1")
	ret0, _ := ret[0].(v1beta114.SchedulingV1beta1Interface)
	return ret0
}

// SchedulingV1beta1 indicates an expected call of SchedulingV1beta1.
func (mr *MockKubernetesClientMockRecorder) SchedulingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).SchedulingV1beta1))
}

// StorageV1 mocks base method.
func (m *MockKubernetesClient) StorageV1() v116.StorageV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1")
	ret0, _ := ret[0].(v116.StorageV1Interface)
	return ret0
}

// StorageV1 indicates an expected call of StorageV1.
func (mr *MockKubernetesClientMockRecorder) StorageV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1", reflect.TypeOf((*MockKubernetesClient)(nil).StorageV1))
}

// StorageV1alpha1 mocks base method.
func (m *MockKubernetesClient) StorageV1alpha1() v1alpha18.StorageV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1alpha1")
	ret0, _ := ret[0].(v1alpha18.StorageV1alpha1Interface)
	return ret0
}

// StorageV1alpha1 indicates an expected call of StorageV1alpha1.
func (mr *MockKubernetesClientMockRecorder) StorageV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).StorageV1alpha1))
}

// StorageV1beta1 mocks base method.
func (m *MockKubernetesClient) StorageV1beta1() v1beta115.StorageV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1beta1")
	ret0, _ := ret[0].(v1beta115.StorageV1beta1Interface)
	return ret0
}

// StorageV1beta1 indicates an expected call of StorageV1beta1.
func (mr *MockKubernetesClientMockRecorder) StorageV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1beta1", reflect.TypeOf((*MockKubernetesClient)(nil).StorageV1beta1))
}

// StoragemigrationV1alpha1 mocks base method.
func (m *MockKubernetesClient) StoragemigrationV1alpha1() v1alpha19.StoragemigrationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoragemigrationV1alpha1")
	ret0, _ := ret[0].(v1alpha19.StoragemigrationV1alpha1Interface)
	return ret0
}

// StoragemigrationV1alpha1 indicates an expected call of StoragemigrationV1alpha1.
func (mr *MockKubernetesClientMockRecorder) StoragemigrationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoragemigrationV1alpha1", reflect.TypeOf((*MockKubernetesClient)(nil).StoragemigrationV1alpha1))
}
