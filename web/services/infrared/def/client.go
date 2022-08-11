// Code generated by jrpc. DO NOT EDIT.

package infrareddef

import (
	context "context"
	"testing"

	taxi "github.com/jakewright/home-automation/libraries/go/taxi"
)

// InfraredService is the public interface of this service
type InfraredService interface {
	GetDevice(ctx context.Context, body *GetDeviceRequest) *GetDeviceFuture
	UpdateDevice(ctx context.Context, body *UpdateDeviceRequest) *UpdateDeviceFuture
}

// GetDeviceFuture represents an in-flight GetDevice request
type GetDeviceFuture struct {
	done <-chan struct{}
	rsp  *GetDeviceResponse
	err  error
}

// Wait blocks until the response is ready
func (f *GetDeviceFuture) Wait() (*GetDeviceResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// UpdateDeviceFuture represents an in-flight UpdateDevice request
type UpdateDeviceFuture struct {
	done <-chan struct{}
	rsp  *UpdateDeviceResponse
	err  error
}

// Wait blocks until the response is ready
func (f *UpdateDeviceFuture) Wait() (*UpdateDeviceResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// Client makes requests to this service
type Client struct {
	dispatcher taxi.Dispatcher
}

// Compile-time assertion that the client implements the interface
var _ InfraredService = (*Client)(nil)

// NewClient returns a new client
func NewClient(dispatcher taxi.Dispatcher) *Client {
	return &Client{
		dispatcher: dispatcher,
	}
}

// GetDevice dispatches an RPC to the service
func (c *Client) GetDevice(ctx context.Context, body *GetDeviceRequest) *GetDeviceFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "http://infrared/device",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &GetDeviceFuture{
		done: done,
		rsp:  &GetDeviceResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(ftr.rsp)
	}()

	return ftr
}

// UpdateDevice dispatches an RPC to the service
func (c *Client) UpdateDevice(ctx context.Context, body *UpdateDeviceRequest) *UpdateDeviceFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "PATCH",
		URL:    "http://infrared/device",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &UpdateDeviceFuture{
		done: done,
		rsp:  &UpdateDeviceResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(ftr.rsp)
	}()

	return ftr
}

// MockClient can be used in tests
type MockClient struct {
	dispatcher *taxi.MockClient
}

// Compile-time assertion that the mock client implements the interface
var _ InfraredService = (*MockClient)(nil)

// NewMockClient returns a new mock client
func NewMockClient(ctx context.Context, t *testing.T) *MockClient {
	f := taxi.NewTestFixture(t)

	return &MockClient{
		dispatcher: &taxi.MockClient{Handler: f},
	}
}

// GetDevice dispatches an RPC to the mock client
func (c *MockClient) GetDevice(ctx context.Context, body *GetDeviceRequest) *GetDeviceFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "GET",
		URL:    "http://infrared/device",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &GetDeviceFuture{
		done: done,
		rsp:  &GetDeviceResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// UpdateDevice dispatches an RPC to the mock client
func (c *MockClient) UpdateDevice(ctx context.Context, body *UpdateDeviceRequest) *UpdateDeviceFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "PATCH",
		URL:    "http://infrared/device",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &UpdateDeviceFuture{
		done: done,
		rsp:  &UpdateDeviceResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}
