package api

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-user-api/internal/models"
	"github.com/ozoncp/ocp-user-api/internal/producer"
	"github.com/ozoncp/ocp-user-api/internal/repo"
	desc "github.com/ozoncp/ocp-user-api/pkg/ocp-user-api"
)

type api struct {
	desc.UnimplementedOcpUserApiServer
	userRepo      repo.Repo
	eventProducer producer.Producer
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

	users := make([]*desc.User, 0, len(searchResult.Items))
	for _, user := range searchResult.Items {
		users = append(users, repoUserToProtoUser(&user))
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
		Name:       req.Profile.GetName(),
		Surname:    req.Profile.GetSurname(),
		Patronymic: req.Profile.GetPatronymic(),
		Email:      req.Profile.GetEmail(),
	}
	userId, err := a.userRepo.CreateUser(ctx, user)

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = a.eventProducer.SendEvent(producer.Event{
		Type: producer.Created,
		Payload: map[string]interface{}{
			"Id": userId,
		},
	})

	if err != nil {
		log.Error().Err(err).Msg("publish telemetry event")
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

	isDeleted, err := a.userRepo.RemoveUser(ctx, req.UserId)

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if isDeleted {
		log.Info().Msgf("user %d was deleted", req.UserId)

		err := a.eventProducer.SendEvent(producer.Event{
			Type: producer.Removed,
			Payload: map[string]interface{}{
				"Id": req.UserId,
			},
		})

		if err != nil {
			log.Error().Err(err).Msg("publish telemetry event")
		}
	}

	return &desc.RemoveUserV1Response{
		Deleted: isDeleted,
	}, nil
}

func (a *api) UpdateUserV1(
	ctx context.Context,
	req *desc.UpdateUserV1Request,
) (*desc.UpdateUserV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("update user %d", req.UserId)

	updated, err := a.userRepo.UpdateUser(ctx, &models.User{
		CalendarId: req.UserParams.CalendarId,
		ResumeId:   req.UserParams.ResumeId,
		Name:       req.UserParams.Profile.GetName(),
		Surname:    req.UserParams.Profile.GetSurname(),
		Patronymic: req.UserParams.Profile.GetPatronymic(),
		Email:      req.UserParams.Profile.GetEmail(),
	})

	if err != nil {
		log.Error().Err(err).Msg("internal error")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if updated {
		log.Info().Msgf("user %d was updated", req.UserId)

		err := a.eventProducer.SendEvent(producer.Event{
			Type: producer.Updated,
			Payload: map[string]interface{}{
				"Id": req.UserId,
			},
		})

		if err != nil {
			log.Error().Err(err).Msg("publish telemetry event")
		}
	}

	return &desc.UpdateUserV1Response{
		Updated: updated,
	}, nil
}

func (a *api) MultiCreateUserV1(
	ctx context.Context,
	req *desc.MultiCreateUserV1Request,
) (*desc.MultiCreateUserV1Response, error) {
	traceId := uuid.New().String()
	span, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("TraceId[%s]", traceId))
	defer span.Finish()

	logger := log.With().
		Str("TraceId", traceId).
		Logger()
	logger.Info().Msgf("create new %d users", len(req.Users))

	if err := req.Validate(); err != nil {
		logger.Error().Err(err).Msg("invalid argument")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	users := make([]models.User, 0, len(req.Users))
	for _, user := range req.Users {
		users = append(users, models.User{
			CalendarId: user.CalendarId,
			ResumeId:   user.ResumeId,
			Name:       user.Profile.GetName(),
			Surname:    user.Profile.GetSurname(),
			Patronymic: user.Profile.GetPatronymic(),
			Email:      user.Profile.GetEmail(),
		})
	}
	ids, err := a.userRepo.CreateUsers(ctx, users)

	if err != nil {
		logger.Error().Err(err).Msg("internal error")

		if len(ids) == 0 {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	for _, id := range ids {
		err := a.eventProducer.SendEvent(producer.Event{
			Type: producer.Created,
			Payload: map[string]interface{}{
				"Id": id,
			},
		})

		if err != nil {
			logger.Error().Err(err).Msg("publish telemetry event")
		}
	}

	return &desc.MultiCreateUserV1Response{
		Count: int64(len(users)),
	}, nil
}

func NewOcpUserApi(userRepo repo.Repo, eventProducer producer.Producer) desc.OcpUserApiServer {
	return &api{
		userRepo:      userRepo,
		eventProducer: eventProducer,
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
