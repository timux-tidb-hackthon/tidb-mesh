package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "172.24.57.248:50050", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	// conn, err := grpc.DialContext(context.TODO(), *addr, grpc.WithInsecure(), grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
	// 	var d net.Dialer
	// 	return d.DialContext(ctx, "tcp", "127.0.0.1:30000")
	// }) )
	conn, err := grpc.DialContext(context.TODO(), *addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}