// Copyright 2020 UTN - FRRe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package foo implements bar functionality
package foo

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	bar "github.com/FRRe-DACS/service-mesh/pkg/dacs/bar"
	"github.com/FRRe-DACS/service-mesh/pkg/dacs/utils"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server represents the gRPC server for Foo Services
type Server struct {
	BarClient bar.BarServiceClient
}

var logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "bar")))

// CreateFoo Creates a new Foo in the system
func (s *Server) CreateFoo(ctx context.Context, in *CreateFooRequest) (*CreateFooResponse, error) {
	logger.Info(fmt.Sprintf("Creating Foo: %d", in.Foo.Id))

	response, err := s.BarClient.GetBar(ctx, &bar.GetBarRequest{
		Id: in.Foo.Bar.Id,
	})

	if err != nil {
		return nil, fmt.Errorf("could create Foo: %s", err)
	}

	// TODO Add logic here
	return &CreateFooResponse{
		Foo: &Foo{
			Id:  1,
			Bar: response.Bar,
		},
	}, nil
}

// StartFooRESTServer Starts s REST Reverse proxy service for Foo
func StartFooRESTServer(address, grpcAddress string) error {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping
	err = RegisterFooServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Foo  Service: %s", err)
	}

	logger.Info(fmt.Sprintf("Starting HTTP/1.1 REST server on %s", address))
	http.ListenAndServe(address, mux)

	return nil
}

// StartFooGRPCServer Starts a gRPC Server for Foo
func StartFooGRPCServer(address string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	var conn *grpc.ClientConn
	var barHostname = utils.GetEnv("BAR_HOSTNAME", "localhost")
	var barPort = utils.GetEnv("BAR_GRPC_PORT", "9090")
	var barURL = fmt.Sprintf("%s:%s", barHostname, barPort)

	logger.Info(fmt.Sprintf("Connecting to Bar Service at %s", barURL))

	conn, err = grpc.Dial(barURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := bar.NewBarServiceClient(conn)

	// create a server instance
	s := Server{
		BarClient: client,
	}

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	RegisterFooServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	// start the server
	logger.Info(fmt.Sprintf("Starting HTTP/2 gRPC server on %s", address))

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}
