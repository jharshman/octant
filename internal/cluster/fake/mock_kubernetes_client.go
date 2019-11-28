// Code generated by MockGen. DO NOT EDIT.
// Source: ../../vendor/k8s.io/client-go/kubernetes/clientset.go

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	discovery "k8s.io/client-go/discovery"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1beta10 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/auditregistration/v1alpha1"
	v10 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	v11 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	v1beta12 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	v12 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	v2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	v13 "k8s.io/client-go/kubernetes/typed/batch/v1"
	v1beta13 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	v2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	v1beta14 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	v14 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	v1beta15 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	v15 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1beta16 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	v1beta17 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	v16 "k8s.io/client-go/kubernetes/typed/networking/v1"
	v1beta18 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	v1alpha10 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	v1beta19 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	v1beta110 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	v17 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	v1alpha11 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	v1beta111 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	v18 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	v1alpha12 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	v1beta112 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	v1alpha13 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	v19 "k8s.io/client-go/kubernetes/typed/storage/v1"
	v1alpha14 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	v1beta113 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	reflect "reflect"
)

// MockKubernetesInterface is a mock of Interface interface
type MockKubernetesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesInterfaceMockRecorder
}

// MockKubernetesInterfaceMockRecorder is the mock recorder for MockKubernetesInterface
type MockKubernetesInterfaceMockRecorder struct {
	mock *MockKubernetesInterface
}

// NewMockKubernetesInterface creates a new mock instance
func NewMockKubernetesInterface(ctrl *gomock.Controller) *MockKubernetesInterface {
	mock := &MockKubernetesInterface{ctrl: ctrl}
	mock.recorder = &MockKubernetesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubernetesInterface) EXPECT() *MockKubernetesInterfaceMockRecorder {
	return m.recorder
}

// Discovery mocks base method
func (m *MockKubernetesInterface) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

// Discovery indicates an expected call of Discovery
func (mr *MockKubernetesInterfaceMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discovery", reflect.TypeOf((*MockKubernetesInterface)(nil).Discovery))
}

// AdmissionregistrationV1beta1 mocks base method
func (m *MockKubernetesInterface) AdmissionregistrationV1beta1() v1beta1.AdmissionregistrationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmissionregistrationV1beta1")
	ret0, _ := ret[0].(v1beta1.AdmissionregistrationV1beta1Interface)
	return ret0
}

// AdmissionregistrationV1beta1 indicates an expected call of AdmissionregistrationV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) AdmissionregistrationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmissionregistrationV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).AdmissionregistrationV1beta1))
}

// AppsV1 mocks base method
func (m *MockKubernetesInterface) AppsV1() v1.AppsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1")
	ret0, _ := ret[0].(v1.AppsV1Interface)
	return ret0
}

// AppsV1 indicates an expected call of AppsV1
func (mr *MockKubernetesInterfaceMockRecorder) AppsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1", reflect.TypeOf((*MockKubernetesInterface)(nil).AppsV1))
}

// AppsV1beta1 mocks base method
func (m *MockKubernetesInterface) AppsV1beta1() v1beta10.AppsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta1")
	ret0, _ := ret[0].(v1beta10.AppsV1beta1Interface)
	return ret0
}

// AppsV1beta1 indicates an expected call of AppsV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) AppsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).AppsV1beta1))
}

// AppsV1beta2 mocks base method
func (m *MockKubernetesInterface) AppsV1beta2() v1beta2.AppsV1beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta2")
	ret0, _ := ret[0].(v1beta2.AppsV1beta2Interface)
	return ret0
}

// AppsV1beta2 indicates an expected call of AppsV1beta2
func (mr *MockKubernetesInterfaceMockRecorder) AppsV1beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta2", reflect.TypeOf((*MockKubernetesInterface)(nil).AppsV1beta2))
}

// AuditregistrationV1alpha1 mocks base method
func (m *MockKubernetesInterface) AuditregistrationV1alpha1() v1alpha1.AuditregistrationV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuditregistrationV1alpha1")
	ret0, _ := ret[0].(v1alpha1.AuditregistrationV1alpha1Interface)
	return ret0
}

// AuditregistrationV1alpha1 indicates an expected call of AuditregistrationV1alpha1
func (mr *MockKubernetesInterfaceMockRecorder) AuditregistrationV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuditregistrationV1alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).AuditregistrationV1alpha1))
}

// AuthenticationV1 mocks base method
func (m *MockKubernetesInterface) AuthenticationV1() v10.AuthenticationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1")
	ret0, _ := ret[0].(v10.AuthenticationV1Interface)
	return ret0
}

// AuthenticationV1 indicates an expected call of AuthenticationV1
func (mr *MockKubernetesInterfaceMockRecorder) AuthenticationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1", reflect.TypeOf((*MockKubernetesInterface)(nil).AuthenticationV1))
}

// AuthenticationV1beta1 mocks base method
func (m *MockKubernetesInterface) AuthenticationV1beta1() v1beta11.AuthenticationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1beta1")
	ret0, _ := ret[0].(v1beta11.AuthenticationV1beta1Interface)
	return ret0
}

// AuthenticationV1beta1 indicates an expected call of AuthenticationV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) AuthenticationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).AuthenticationV1beta1))
}

// AuthorizationV1 mocks base method
func (m *MockKubernetesInterface) AuthorizationV1() v11.AuthorizationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV1")
	ret0, _ := ret[0].(v11.AuthorizationV1Interface)
	return ret0
}

// AuthorizationV1 indicates an expected call of AuthorizationV1
func (mr *MockKubernetesInterfaceMockRecorder) AuthorizationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV1", reflect.TypeOf((*MockKubernetesInterface)(nil).AuthorizationV1))
}

// AuthorizationV1beta1 mocks base method
func (m *MockKubernetesInterface) AuthorizationV1beta1() v1beta12.AuthorizationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV1beta1")
	ret0, _ := ret[0].(v1beta12.AuthorizationV1beta1Interface)
	return ret0
}

// AuthorizationV1beta1 indicates an expected call of AuthorizationV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) AuthorizationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).AuthorizationV1beta1))
}

// AutoscalingV1 mocks base method
func (m *MockKubernetesInterface) AutoscalingV1() v12.AutoscalingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV1")
	ret0, _ := ret[0].(v12.AutoscalingV1Interface)
	return ret0
}

// AutoscalingV1 indicates an expected call of AutoscalingV1
func (mr *MockKubernetesInterfaceMockRecorder) AutoscalingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV1", reflect.TypeOf((*MockKubernetesInterface)(nil).AutoscalingV1))
}

// AutoscalingV2beta1 mocks base method
func (m *MockKubernetesInterface) AutoscalingV2beta1() v2beta1.AutoscalingV2beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta1")
	ret0, _ := ret[0].(v2beta1.AutoscalingV2beta1Interface)
	return ret0
}

// AutoscalingV2beta1 indicates an expected call of AutoscalingV2beta1
func (mr *MockKubernetesInterfaceMockRecorder) AutoscalingV2beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).AutoscalingV2beta1))
}

// AutoscalingV2beta2 mocks base method
func (m *MockKubernetesInterface) AutoscalingV2beta2() v2beta2.AutoscalingV2beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta2")
	ret0, _ := ret[0].(v2beta2.AutoscalingV2beta2Interface)
	return ret0
}

// AutoscalingV2beta2 indicates an expected call of AutoscalingV2beta2
func (mr *MockKubernetesInterfaceMockRecorder) AutoscalingV2beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta2", reflect.TypeOf((*MockKubernetesInterface)(nil).AutoscalingV2beta2))
}

// BatchV1 mocks base method
func (m *MockKubernetesInterface) BatchV1() v13.BatchV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV1")
	ret0, _ := ret[0].(v13.BatchV1Interface)
	return ret0
}

// BatchV1 indicates an expected call of BatchV1
func (mr *MockKubernetesInterfaceMockRecorder) BatchV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV1", reflect.TypeOf((*MockKubernetesInterface)(nil).BatchV1))
}

// BatchV1beta1 mocks base method
func (m *MockKubernetesInterface) BatchV1beta1() v1beta13.BatchV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV1beta1")
	ret0, _ := ret[0].(v1beta13.BatchV1beta1Interface)
	return ret0
}

// BatchV1beta1 indicates an expected call of BatchV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) BatchV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).BatchV1beta1))
}

// BatchV2alpha1 mocks base method
func (m *MockKubernetesInterface) BatchV2alpha1() v2alpha1.BatchV2alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV2alpha1")
	ret0, _ := ret[0].(v2alpha1.BatchV2alpha1Interface)
	return ret0
}

// BatchV2alpha1 indicates an expected call of BatchV2alpha1
func (mr *MockKubernetesInterfaceMockRecorder) BatchV2alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV2alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).BatchV2alpha1))
}

// CertificatesV1beta1 mocks base method
func (m *MockKubernetesInterface) CertificatesV1beta1() v1beta14.CertificatesV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificatesV1beta1")
	ret0, _ := ret[0].(v1beta14.CertificatesV1beta1Interface)
	return ret0
}

// CertificatesV1beta1 indicates an expected call of CertificatesV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) CertificatesV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificatesV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).CertificatesV1beta1))
}

// CoordinationV1beta1 mocks base method
func (m *MockKubernetesInterface) CoordinationV1beta1() v1beta15.CoordinationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1beta1")
	ret0, _ := ret[0].(v1beta15.CoordinationV1beta1Interface)
	return ret0
}

// CoordinationV1beta1 indicates an expected call of CoordinationV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) CoordinationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).CoordinationV1beta1))
}

// CoordinationV1 mocks base method
func (m *MockKubernetesInterface) CoordinationV1() v14.CoordinationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1")
	ret0, _ := ret[0].(v14.CoordinationV1Interface)
	return ret0
}

// CoordinationV1 indicates an expected call of CoordinationV1
func (mr *MockKubernetesInterfaceMockRecorder) CoordinationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1", reflect.TypeOf((*MockKubernetesInterface)(nil).CoordinationV1))
}

// CoreV1 mocks base method
func (m *MockKubernetesInterface) CoreV1() v15.CoreV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoreV1")
	ret0, _ := ret[0].(v15.CoreV1Interface)
	return ret0
}

// CoreV1 indicates an expected call of CoreV1
func (mr *MockKubernetesInterfaceMockRecorder) CoreV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoreV1", reflect.TypeOf((*MockKubernetesInterface)(nil).CoreV1))
}

// EventsV1beta1 mocks base method
func (m *MockKubernetesInterface) EventsV1beta1() v1beta16.EventsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsV1beta1")
	ret0, _ := ret[0].(v1beta16.EventsV1beta1Interface)
	return ret0
}

// EventsV1beta1 indicates an expected call of EventsV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) EventsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).EventsV1beta1))
}

// ExtensionsV1beta1 mocks base method
func (m *MockKubernetesInterface) ExtensionsV1beta1() v1beta17.ExtensionsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtensionsV1beta1")
	ret0, _ := ret[0].(v1beta17.ExtensionsV1beta1Interface)
	return ret0
}

// ExtensionsV1beta1 indicates an expected call of ExtensionsV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) ExtensionsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionsV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).ExtensionsV1beta1))
}

// NetworkingV1 mocks base method
func (m *MockKubernetesInterface) NetworkingV1() v16.NetworkingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1")
	ret0, _ := ret[0].(v16.NetworkingV1Interface)
	return ret0
}

// NetworkingV1 indicates an expected call of NetworkingV1
func (mr *MockKubernetesInterfaceMockRecorder) NetworkingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1", reflect.TypeOf((*MockKubernetesInterface)(nil).NetworkingV1))
}

// NetworkingV1beta1 mocks base method
func (m *MockKubernetesInterface) NetworkingV1beta1() v1beta18.NetworkingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1beta1")
	ret0, _ := ret[0].(v1beta18.NetworkingV1beta1Interface)
	return ret0
}

// NetworkingV1beta1 indicates an expected call of NetworkingV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) NetworkingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).NetworkingV1beta1))
}

// NodeV1alpha1 mocks base method
func (m *MockKubernetesInterface) NodeV1alpha1() v1alpha10.NodeV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1alpha1")
	ret0, _ := ret[0].(v1alpha10.NodeV1alpha1Interface)
	return ret0
}

// NodeV1alpha1 indicates an expected call of NodeV1alpha1
func (mr *MockKubernetesInterfaceMockRecorder) NodeV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).NodeV1alpha1))
}

// NodeV1beta1 mocks base method
func (m *MockKubernetesInterface) NodeV1beta1() v1beta19.NodeV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1beta1")
	ret0, _ := ret[0].(v1beta19.NodeV1beta1Interface)
	return ret0
}

// NodeV1beta1 indicates an expected call of NodeV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) NodeV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).NodeV1beta1))
}

// PolicyV1beta1 mocks base method
func (m *MockKubernetesInterface) PolicyV1beta1() v1beta110.PolicyV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyV1beta1")
	ret0, _ := ret[0].(v1beta110.PolicyV1beta1Interface)
	return ret0
}

// PolicyV1beta1 indicates an expected call of PolicyV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) PolicyV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).PolicyV1beta1))
}

// RbacV1 mocks base method
func (m *MockKubernetesInterface) RbacV1() v17.RbacV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1")
	ret0, _ := ret[0].(v17.RbacV1Interface)
	return ret0
}

// RbacV1 indicates an expected call of RbacV1
func (mr *MockKubernetesInterfaceMockRecorder) RbacV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1", reflect.TypeOf((*MockKubernetesInterface)(nil).RbacV1))
}

// RbacV1beta1 mocks base method
func (m *MockKubernetesInterface) RbacV1beta1() v1beta111.RbacV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1beta1")
	ret0, _ := ret[0].(v1beta111.RbacV1beta1Interface)
	return ret0
}

// RbacV1beta1 indicates an expected call of RbacV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) RbacV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).RbacV1beta1))
}

// RbacV1alpha1 mocks base method
func (m *MockKubernetesInterface) RbacV1alpha1() v1alpha11.RbacV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1alpha1")
	ret0, _ := ret[0].(v1alpha11.RbacV1alpha1Interface)
	return ret0
}

// RbacV1alpha1 indicates an expected call of RbacV1alpha1
func (mr *MockKubernetesInterfaceMockRecorder) RbacV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).RbacV1alpha1))
}

// SchedulingV1alpha1 mocks base method
func (m *MockKubernetesInterface) SchedulingV1alpha1() v1alpha12.SchedulingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1alpha1")
	ret0, _ := ret[0].(v1alpha12.SchedulingV1alpha1Interface)
	return ret0
}

// SchedulingV1alpha1 indicates an expected call of SchedulingV1alpha1
func (mr *MockKubernetesInterfaceMockRecorder) SchedulingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).SchedulingV1alpha1))
}

// SchedulingV1beta1 mocks base method
func (m *MockKubernetesInterface) SchedulingV1beta1() v1beta112.SchedulingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1beta1")
	ret0, _ := ret[0].(v1beta112.SchedulingV1beta1Interface)
	return ret0
}

// SchedulingV1beta1 indicates an expected call of SchedulingV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) SchedulingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).SchedulingV1beta1))
}

// SchedulingV1 mocks base method
func (m *MockKubernetesInterface) SchedulingV1() v18.SchedulingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1")
	ret0, _ := ret[0].(v18.SchedulingV1Interface)
	return ret0
}

// SchedulingV1 indicates an expected call of SchedulingV1
func (mr *MockKubernetesInterfaceMockRecorder) SchedulingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1", reflect.TypeOf((*MockKubernetesInterface)(nil).SchedulingV1))
}

// SettingsV1alpha1 mocks base method
func (m *MockKubernetesInterface) SettingsV1alpha1() v1alpha13.SettingsV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettingsV1alpha1")
	ret0, _ := ret[0].(v1alpha13.SettingsV1alpha1Interface)
	return ret0
}

// SettingsV1alpha1 indicates an expected call of SettingsV1alpha1
func (mr *MockKubernetesInterfaceMockRecorder) SettingsV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettingsV1alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).SettingsV1alpha1))
}

// StorageV1beta1 mocks base method
func (m *MockKubernetesInterface) StorageV1beta1() v1beta113.StorageV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1beta1")
	ret0, _ := ret[0].(v1beta113.StorageV1beta1Interface)
	return ret0
}

// StorageV1beta1 indicates an expected call of StorageV1beta1
func (mr *MockKubernetesInterfaceMockRecorder) StorageV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1beta1", reflect.TypeOf((*MockKubernetesInterface)(nil).StorageV1beta1))
}

// StorageV1 mocks base method
func (m *MockKubernetesInterface) StorageV1() v19.StorageV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1")
	ret0, _ := ret[0].(v19.StorageV1Interface)
	return ret0
}

// StorageV1 indicates an expected call of StorageV1
func (mr *MockKubernetesInterfaceMockRecorder) StorageV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1", reflect.TypeOf((*MockKubernetesInterface)(nil).StorageV1))
}

// StorageV1alpha1 mocks base method
func (m *MockKubernetesInterface) StorageV1alpha1() v1alpha14.StorageV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1alpha1")
	ret0, _ := ret[0].(v1alpha14.StorageV1alpha1Interface)
	return ret0
}

// StorageV1alpha1 indicates an expected call of StorageV1alpha1
func (mr *MockKubernetesInterfaceMockRecorder) StorageV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1alpha1", reflect.TypeOf((*MockKubernetesInterface)(nil).StorageV1alpha1))
}