package main

import (
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/ozoncp/ocp-user-api/internal/api"
	"github.com/ozoncp/ocp-user-api/internal/producer"
	"github.com/ozoncp/ocp-user-api/internal/repo"
	desc "github.com/ozoncp/ocp-user-api/pkg/ocp-user-api"
)

const (
	grpcPort    = "7002"
	serviceName = "OcpUserApi"
)

func runGrpc() {
	listen, err := net.Listen("tcp", ":"+grpcPort)

	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
	}

	db, err := sqlx.Connect("postgres", "postgres://ocpuser:ocpuser@localhost:5432/ocp?sslmode=disable")
	if err != nil {
		log.Error().Err(err).Msg("error open db")
		return
	}

	userRepo := repo.NewRepo(db)
	eventProducer := producer.NewProducer("ocp", "user")
	grpcServer := grpc.NewServer()
	desc.RegisterOcpUserApiServer(grpcServer, api.NewOcpUserApi(userRepo, eventProducer))

	log.Info().Str("address", "localhost"+grpcPort).Msg("grpc server started")

	if err := grpcServer.Serve(listen); err != nil {
		log.Error().Err(err).Msg("failed to serve")
	}
}

func initTracing() {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("jaeger tracer creation failed")
	}

	opentracing.SetGlobalTracer(tracer)
}

func main() {
	initTracing()
	runGrpc()
}
