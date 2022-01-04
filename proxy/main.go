package main

import (
	"flag"
	"fmt"

	grpc_proxy "github.com/bradleyjkemp/grpc-tools/grpc-proxy"
	"google.golang.org/grpc"

	// "reflect"
	"time"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
	grpc_proxy.RegisterDefaultFlags()
	flag.Parse()
	proxy, _ := grpc_proxy.New(
		grpc_proxy.WithInterceptor(intercept),
		grpc_proxy.DefaultFlags(),
		grpc_proxy.Port(30000),
	)
	proxy.Start()
}

type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface {}) error {
	if err := w.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	fmt.Printf("Receive a message %s at %v\n", m.(*[]uint8), time.Now().Format(time.RFC3339))

	return nil
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message %s at %v\n", m.([]uint8), time.Now().Format(time.RFC3339))
	if err := w.ServerStream.SendMsg(m); err != nil {
		return err
	}

	return nil
}

func intercept(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	newStream := newWrappedStream(ss)
	return handler(srv, newStream)
}
