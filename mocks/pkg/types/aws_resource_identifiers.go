// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

// AWSResourceIdentifiers is an autogenerated mock type for the AWSResourceIdentifiers type
type AWSResourceIdentifiers struct {
	mock.Mock
}

// ARN provides a mock function with given fields:
func (_m *AWSResourceIdentifiers) ARN() *v1alpha1.AWSResourceName {
	ret := _m.Called()

	var r0 *v1alpha1.AWSResourceName
	if rf, ok := ret.Get(0).(func() *v1alpha1.AWSResourceName); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.AWSResourceName)
		}
	}

	return r0
}

// OwnerAccountID provides a mock function with given fields:
func (_m *AWSResourceIdentifiers) OwnerAccountID() *v1alpha1.AWSAccountID {
	ret := _m.Called()

	var r0 *v1alpha1.AWSAccountID
	if rf, ok := ret.Get(0).(func() *v1alpha1.AWSAccountID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.AWSAccountID)
		}
	}

	return r0
}
