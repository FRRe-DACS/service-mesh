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

// Package bar implements bar functionality
package bar

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server represents the gRPC server for Bar Services
type Server struct {
}

var logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "bar")))

// CreateBar Creates a new Bar in the system
func (s *Server) CreateBar(ctx context.Context, in *CreateBarRequest) (*CreateBarResponse, error) {
	logger.Info(fmt.Sprintf("Creating Bar: %d", in.Bar.Id))
	// TODO Add logic here
	return &CreateBarResponse{
		Bar: &Bar{
			Id:          1,
			Name:        "Cosme",
			Description: "Cosme fulanito",
		},
	}, nil
}

// ListBars List bars in the system
func (s *Server) ListBars(ctx context.Context, in *ListBarsRequest) (*ListBarsResponse, error) {
	logger.Info("Listing Bars.")

	var bars = []*Bar{}

	bars = append(bars, &Bar{
		Id:          1,
		Name:        "Cosme",
		Description: "Cosme fulanito",
	})

	return &ListBarsResponse{
		Bars: bars,
	}, nil
}

// GetBar Get an existing bars
func (s *Server) GetBar(ctx context.Context, in *GetBarRequest) (*GetBarResponse, error) {
	logger.Info(fmt.Sprintf("Getting Bar: %d", in.Id))
	var bar = &Bar{
		Id:          1,
		Name:        "Cosme",
		Description: "Cosme fulanito",
	}
	return &GetBarResponse{
		Bar: bar,
	}, nil
}

// UpdateBar Updates an existing bars
func (s *Server) UpdateBar(ctx context.Context, in *UpdateBarRequest) (*UpdateBarResponse, error) {
	logger.Info(fmt.Sprintf("Updating Bar: %s", in.Bar.Id))
	return &UpdateBarResponse{
		Bar: &Bar{
			Id:          1,
			Name:        "Cosme",
			Description: "Cosme fulanito",
		},
	}, nil
}

// DeleteBar Deletes an existing bar
func (s *Server) DeleteBar(ctx context.Context, in *DeleteBarRequest) (*DeleteBarResponse, error) {
	logger.Info(fmt.Sprintf("Deleting Bar: %d", in.Id))
	return &DeleteBarResponse{}, nil
}

// StartBarRESTServer Starts s REST Reverse proxy service for Bar
func StartBarRESTServer(address, grpcAddress string) error {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping
	err = RegisterBarServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Bar  Service: %s", err)
	}

	logger.Info(fmt.Sprintf("Starting HTTP/1.1 REST server on %s", address))
	http.ListenAndServe(address, mux)

	return nil
}

// StartBarGRPCServer Starts a gRPC Server for Bar
func StartBarGRPCServer(address string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	s := Server{}

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	RegisterBarServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	// start the server
	logger.Info(fmt.Sprintf("Starting HTTP/2 gRPC server on %s", address))

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}
