package main

import (
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-user-api/internal/api"
	"github.com/ozoncp/ocp-user-api/internal/repo"
	desc "github.com/ozoncp/ocp-user-api/pkg/ocp-user-api"
)

const (
	grpcPort    = "7002"
)

func runGrpc() {
	listen, err := net.Listen("tcp", ":" + grpcPort)

	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
	}

	db, err := sqlx.Connect("postgres", "postgres://ocpuser:ocpuser@localhost:5432/ocp?sslmode=disable")
	if err != nil {
		log.Error().Err(err).Msg("error open db")
		return
	}

	userRepo := repo.NewRepo(db)
	grpcServer := grpc.NewServer()
	desc.RegisterOcpUserApiServer(grpcServer, api.NewOcpUserApi(userRepo))

	log.Info().Str("address", "localhost"+grpcPort).Msg("grpc server started")

	if err := grpcServer.Serve(listen); err != nil {
		log.Error().Err(err).Msg("failed to serve")
	}
}

func main() {
	runGrpc()
}
