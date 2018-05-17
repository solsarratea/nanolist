// Code generated by mockery v1.0.0. DO NOT EDIT.
package main

import mock "github.com/stretchr/testify/mock"

// MockSender is an autogenerated mock type for the Sender type
type MockSender struct {
	mock.Mock
}

// Send provides a mock function with given fields: msg, recipients
func (_m *MockSender) Send(msg *Message, recipients []string) error {
	ret := _m.Called(msg, recipients)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Message, []string) error); ok {
		r0 = rf(msg, recipients)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
