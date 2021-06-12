package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-user-api/internal/models"
	"github.com/ozoncp/ocp-user-api/internal/repo"
	desc "github.com/ozoncp/ocp-user-api/pkg/ocp-user-api"
)

type api struct {
	desc.UnimplementedOcpUserApiServer
	userRepo repo.Repo
}

func (a *api) ListUsersV1(
	ctx context.Context,
	req *desc.ListUsersV1Request,
) (*desc.ListUsersV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("search users: Limit: %d, Offset: %d", req.Limit, req.Offset)

	searchParams := models.UserSearchParams{
		Count:  req.Limit,
		Offset: req.Offset,
	}
	searchResult, err := a.userRepo.SearchUsers(ctx, searchParams)

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	users := make([]*desc.User, len(searchResult.Items))
	for i, user := range searchResult.Items {
		users[i] = repoUserToProtoUser(&user)
	}

	log.Info().Msgf("found %d users, NextOffset: %d", len(users), searchResult.NextOffset)

	return &desc.ListUsersV1Response{
		Users:      users,
		NextOffset: searchResult.NextOffset,
	}, nil
}

func (a *api) DescribeUserV1(
	ctx context.Context,
	req *desc.DescribeUserV1Request,
) (*desc.DescribeUserV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("get user %d", req.UserId)

	user, err := a.userRepo.GetUser(ctx, req.UserId)

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		log.Info().Msgf("user %d not found", req.UserId)
		return nil, status.Error(codes.NotFound, "user was not found")
	}

	log.Debug().Msgf(" found user% %v", user)

	return &desc.DescribeUserV1Response{
		User: repoUserToProtoUser(user),
	}, nil
}

func (a *api) CreateUserV1(
	ctx context.Context,
	req *desc.CreateUserV1Request,
) (*desc.CreateUserV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user := &models.User{
		CalendarId: req.CalendarId,
		ResumeId:   req.ResumeId,
		Name:       req.Profile.Name,
		Surname:    req.Profile.Surname,
		Patronymic: req.Profile.Patronymic,
		Email:      req.Profile.Email,
	}
	userId, err := a.userRepo.CreateUser(ctx, user)

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info().Msgf("create new user %d", userId)

	return &desc.CreateUserV1Response{
		UserId: userId,
	}, nil
}

func (a *api) RemoveUserV1(
	ctx context.Context,
	req *desc.RemoveUserV1Request,
) (*desc.RemoveUserV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("remove user %d", req.UserId)

	err := a.userRepo.RemoveUser(ctx, req.UserId)

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info().Msgf("user %d was deleted", req.UserId)

	return &desc.RemoveUserV1Response{
		Deleted: true,
	}, nil
}

func NewOcpUserApi(userRepo repo.Repo) desc.OcpUserApiServer {
	return &api{
		userRepo: userRepo,
	}
}

func repoUserToProtoUser(user *models.User) *desc.User {
	return &desc.User{
		Id:         user.Id,
		CalendarId: user.CalendarId,
		ResumeId:   user.ResumeId,
		Profile: &desc.UserProfile{
			Name:       user.Name,
			Surname:    user.Surname,
			Patronymic: user.Patronymic,
			Email:      user.Email,
		},
	}
}
