module github.com/ozoncp/ocp-user-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jmoiron/sqlx v1.3.4 // indirect
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.10.1
	github.com/ozoncp/ocp-user-api/pkg/ocp-user-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0 // indirect
	google.golang.org/genproto v0.0.0-20210608205507-b6d2f5bf0d7d // indirect
	google.golang.org/grpc v1.38.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/ozoncp/ocp-user-api/pkg/ocp-user-api => ./pkg/ocp-user-api
