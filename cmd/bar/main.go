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
package main

import (
	"fmt"
	"os"

	"github.com/FRRe-DACS/service-mesh/pkg/dacs/bar"
	"go.uber.org/zap"
)

const (
	// Default port for REST Web API
	defaultRestPort = "8080"

	// Default port for gRPC API
	defaultGpcpPort = "9090"
)

var logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))

func main() {
	grpcPort := getEnv("GRPC_PORT", defaultGpcpPort)
	restPort := getEnv("REST_PORT", defaultRestPort)

	grpcAddress := fmt.Sprintf("%s:%s", "0.0.0.0", grpcPort)
	restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", restPort)

	// fire the gRPC server in a goroutine
	go func() {
		err := bar.StartBarGRPCServer(grpcAddress)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := bar.StartBarRESTServer(restAddress, grpcAddress)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
		}
	}()

	// infinite loop
	logger.Info("Entering infinite loop")

	select {}
}

// getEnv Get environment variable value or fallback value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
