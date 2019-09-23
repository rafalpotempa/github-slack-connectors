// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import k8scomponents "github.com/kyma-incubator/hack-showcase/scenario/azure-comments-analytics/internal/k8scomponents"
import mock "github.com/stretchr/testify/mock"

// EventbusWrapper is an autogenerated mock type for the EventbusWrapper type
type EventbusWrapper struct {
	mock.Mock
}

// Subscription provides a mock function with given fields: namespace
func (_m *EventbusWrapper) Subscription(namespace string) k8scomponents.Subscription {
	ret := _m.Called(namespace)

	var r0 k8scomponents.Subscription
	if rf, ok := ret.Get(0).(func(string) k8scomponents.Subscription); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(k8scomponents.Subscription)
		}
	}

	return r0
}