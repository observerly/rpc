/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/rpc
//	@license	Copyright © 2021-2023 observerly

/*****************************************************************************************************************/

package main

/*****************************************************************************************************************/

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/observerly/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*****************************************************************************************************************/

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

/*****************************************************************************************************************/

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewPingClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	// Send the connected status to the server:
	r, err := c.IsConnected(ctx, &pb.PingRequest{})

	if err != nil {
		log.Fatalf("could not get connected status: %v", err)
	}

	log.Printf("Connected: %v", r.GetConnected())
}

/*****************************************************************************************************************/
