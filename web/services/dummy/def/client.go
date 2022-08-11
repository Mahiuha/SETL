// Code generated by jrpc. DO NOT EDIT.

package dummydef

import (
	context "context"
	"testing"

	taxi "github.com/jakewright/home-automation/libraries/go/taxi"
)

// DummyService is the public interface of this service
type DummyService interface {
	Log(ctx context.Context, body *LogRequest) *LogFuture
	Panic(ctx context.Context, body *PanicRequest) *PanicFuture
}

// LogFuture represents an in-flight Log request
type LogFuture struct {
	done <-chan struct{}
	rsp  *LogResponse
	err  error
}

// Wait blocks until the response is ready
func (f *LogFuture) Wait() (*LogResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// PanicFuture represents an in-flight Panic request
type PanicFuture struct {
	done <-chan struct{}
	rsp  *PanicResponse
	err  error
}

// Wait blocks until the response is ready
func (f *PanicFuture) Wait() (*PanicResponse, error) {
	<-f.done
	return f.rsp, f.err
}

// Client makes requests to this service
type Client struct {
	dispatcher taxi.Dispatcher
}

// Compile-time assertion that the client implements the interface
var _ DummyService = (*Client)(nil)

// NewClient returns a new client
func NewClient(dispatcher taxi.Dispatcher) *Client {
	return &Client{
		dispatcher: dispatcher,
	}
}

// Log dispatches an RPC to the service
func (c *Client) Log(ctx context.Context, body *LogRequest) *LogFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "http://dummy/log",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &LogFuture{
		done: done,
		rsp:  &LogResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(ftr.rsp)
	}()

	return ftr
}

// Panic dispatches an RPC to the service
func (c *Client) Panic(ctx context.Context, body *PanicRequest) *PanicFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "http://dummy/panic",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &PanicFuture{
		done: done,
		rsp:  &PanicResponse{},
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
var _ DummyService = (*MockClient)(nil)

// NewMockClient returns a new mock client
func NewMockClient(ctx context.Context, t *testing.T) *MockClient {
	f := taxi.NewTestFixture(t)

	return &MockClient{
		dispatcher: &taxi.MockClient{Handler: f},
	}
}

// Log dispatches an RPC to the mock client
func (c *MockClient) Log(ctx context.Context, body *LogRequest) *LogFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "http://dummy/log",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &LogFuture{
		done: done,
		rsp:  &LogResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}

// Panic dispatches an RPC to the mock client
func (c *MockClient) Panic(ctx context.Context, body *PanicRequest) *PanicFuture {
	taxiFtr := c.dispatcher.Dispatch(ctx, &taxi.RPC{
		Method: "POST",
		URL:    "http://dummy/panic",
		Body:   body,
	})

	done := make(chan struct{})
	ftr := &PanicFuture{
		done: done,
		rsp:  &PanicResponse{},
	}

	go func() {
		defer close(done)
		ftr.err = taxiFtr.DecodeResponse(&ftr.rsp)
	}()

	return ftr
}