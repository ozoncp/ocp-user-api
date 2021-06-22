module github.com/ozoncp/ocp-user-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.0
	github.com/golang/mock v1.5.0
	github.com/google/uuid v1.2.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.10.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-user-api/pkg/ocp-user-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.8.0 // indirect
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-user-api/pkg/ocp-user-api => ./pkg/ocp-user-api
