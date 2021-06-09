package main

import (
	"net"

	"github.com/ozoncp/ocp-user-api/internal/api"
	desc "github.com/ozoncp/ocp-user-api/pkg/ocp-user-api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

const (
	grpcPort = ":7002"
)

func runGrpc() {
	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
	}

	grpcServer := grpc.NewServer()
	desc.RegisterOcpUserApiServer(grpcServer, api.NewOcpUserApi())

	log.Info().Str("address", "localhost"+grpcPort).Msg("grpc server started")

	if err := grpcServer.Serve(listen); err != nil {
		log.Error().Err(err).Msg("failed to serve")
	}
}

func main() {
	runGrpc()
}
