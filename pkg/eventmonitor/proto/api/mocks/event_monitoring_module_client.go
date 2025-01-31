// Code generated by mockery v2.49.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/DataDog/datadog-agent/pkg/eventmonitor/proto/api"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// EventMonitoringModuleClient is an autogenerated mock type for the EventMonitoringModuleClient type
type EventMonitoringModuleClient struct {
	mock.Mock
}

// GetProcessEvents provides a mock function with given fields: ctx, in, opts
func (_m *EventMonitoringModuleClient) GetProcessEvents(ctx context.Context, in *api.GetProcessEventParams, opts ...grpc.CallOption) (grpc.ServerStreamingClient[api.ProcessEventMessage], error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProcessEvents")
	}

	var r0 grpc.ServerStreamingClient[api.ProcessEventMessage]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetProcessEventParams, ...grpc.CallOption) (grpc.ServerStreamingClient[api.ProcessEventMessage], error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.GetProcessEventParams, ...grpc.CallOption) grpc.ServerStreamingClient[api.ProcessEventMessage]); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.ServerStreamingClient[api.ProcessEventMessage])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.GetProcessEventParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEventMonitoringModuleClient creates a new instance of EventMonitoringModuleClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventMonitoringModuleClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventMonitoringModuleClient {
	mock := &EventMonitoringModuleClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
