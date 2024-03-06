/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/rpc/m
//	@license	Copyright © 2021-2023 observerly

/*****************************************************************************************************************/

package main

/*****************************************************************************************************************/

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/observerly/rpc/proto"
	"google.golang.org/grpc"
)

/*****************************************************************************************************************/

var (
	port = flag.Int("port", 50051, "The server port")
)

/*****************************************************************************************************************/

// server is used to implement ping.PingServer.
type server struct {
	pb.UnimplementedPingServer
}

/*****************************************************************************************************************/

// Some random boolean generator to similate a connected status:
func IsConnected() bool {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

/*****************************************************************************************************************/

// IsConnected implements ping.PingServer
func (s *server) IsConnected(ctx context.Context, in *pb.PingRequest) (*pb.PongReply, error) {
	connected := IsConnected()

	log.Printf("Received: %v", connected)

	return &pb.PongReply{Connected: connected}, nil
}

/*****************************************************************************************************************/

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server:
	s := grpc.NewServer()

	// Register the ping server:
	pb.RegisterPingServer(s, &server{})

	// Serve the server:
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

/*****************************************************************************************************************/
