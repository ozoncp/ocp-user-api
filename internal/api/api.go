package api

import (
	"context"

	"github.com/rs/zerolog/log"
	desc "github.com/ozoncp/ocp-user-api/pkg/ocp-user-api"
)

type api struct {
	desc.UnimplementedOcpUserApiServer
}

func (a *api) ListUsersV1(
	ctx context.Context,
	req *desc.ListUsersV1Request,
) (*desc.ListUsersV1Response, error) {
	log.Info().Msg("ListUsersV1")
	return nil, nil
}

func (a *api) DescribeVideoV1(
	ctx context.Context,
	req *desc.DescribeUserV1Request,
) (*desc.DescribeUserV1Response, error) {
	log.Info().Msg("DescribeUserV1")
	return nil, nil
}

func (a *api) CreateUserV1(
	ctx context.Context,
	req *desc.CreateUserV1Request,
) (*desc.CreateUserV1Response, error) {
	log.Info().Msg("CreateUserV1")
	return nil, nil
}

func (a *api) RemoveUserV1(
	ctx context.Context,
	req *desc.RemoveUserV1Request,
) (*desc.RemoveUserV1Response, error) {
	log.Info().Msg("RemoveUserV1")
	return nil, nil
}

func NewOcpUserApi() desc.OcpUserApiServer {
	return &api{}
}
