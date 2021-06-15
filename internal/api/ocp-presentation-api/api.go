// Package api implements a simple API for Ozon Code Platform Presentation API.
package api

import (
	"context"

	"github.com/ozoncp/ocp-presentation-api/internal/model"
	repo "github.com/ozoncp/ocp-presentation-api/internal/repo/presentation"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedPresentationAPIServer
	repo repo.Repo
}

func NewPresentationAPI(repo repo.Repo) desc.PresentationAPIServer {
	return &api{
		repo: repo,
	}
}

func (a *api) CreatePresentationV1(
	ctx context.Context,
	req *desc.CreatePresentationV1Request,
) (
	*desc.CreatePresentationV1Response,
	error,
) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "CreatePresentationV1").
		Uint64("LessonID", req.LessonId).
		Uint64("UserID", req.UserId).
		Str("Name", req.Name).
		Str("Description", req.Description).
		Msg("Input data")

	presentation := model.Presentation{
		LessonID:    req.LessonId,
		UserID:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := a.repo.CreatePresentation(ctx, presentation)
	if err != nil {
		log.Error().Msgf("Failed to insert the data: %v", err)
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.CreatePresentationV1Response{
		PresentationId: id,
	}

	return response, nil
}

func (a *api) DescribePresentationV1(
	ctx context.Context,
	req *desc.DescribePresentationV1Request,
) (
	*desc.DescribePresentationV1Response,
	error,
) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "DescribePresentationV1").
		Uint64("PresentationID", req.PresentationId).
		Msg("Input data")

	presentation, err := a.repo.DescribePresentation(ctx, req.PresentationId)
	if err != nil {
		if err == repo.ErrPresentationNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Err(err)
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.DescribePresentationV1Response{
		Presentation: &desc.Presentation{
			Id:          presentation.ID,
			LessonId:    presentation.LessonID,
			UserId:      presentation.UserID,
			Name:        presentation.Name,
			Description: presentation.Description,
		},
	}

	return response, nil
}

func (a *api) ListPresentationsV1(
	ctx context.Context,
	req *desc.ListPresentationsV1Request,
) (
	*desc.ListPresentationsV1Response,
	error,
) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "ListPresentationsV1").
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("Input data")

	presentations, err := a.repo.ListPresentations(ctx, req.Limit, req.Offset)
	if err != nil {
		if err == repo.ErrPresentationNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Msgf("Failed to fill the list: %v", err)
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	list := make([]*desc.Presentation, 0, len(presentations))
	for _, presentation := range presentations {
		result := &desc.Presentation{
			Id:          presentation.ID,
			LessonId:    presentation.LessonID,
			UserId:      presentation.UserID,
			Name:        presentation.Name,
			Description: presentation.Description,
		}

		list = append(list, result)
	}

	response := &desc.ListPresentationsV1Response{
		Presentations: list,
	}

	return response, nil
}

func (a *api) RemovePresentationV1(
	ctx context.Context,
	req *desc.RemovePresentationV1Request,
) (
	*desc.RemovePresentationV1Response,
	error,
) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "RemovePresentationV1").
		Uint64("PresentationID", req.PresentationId).
		Msg("Input data")

	removed, err := a.repo.RemovePresentation(ctx, req.PresentationId)
	if err != nil {
		log.Error().Msgf("Failed to remove the data: %v", err)
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.RemovePresentationV1Response{
		Found: removed,
	}

	return response, nil
}
