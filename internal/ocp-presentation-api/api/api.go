// Package api implements a simple API for Ozon Code Platform Presentation API.
package api

import (
	"context"
	"errors"
	"fmt"
	"unsafe"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-presentation-api/repo"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInvalidArgument = errors.New("invalid argument")

type api struct {
	desc.UnimplementedPresentationAPIServer
	repo      repo.Repo
	chunkSize uint
}

func NewPresentationAPI(repo repo.Repo, chunkSize uint) desc.PresentationAPIServer {
	if chunkSize == 0 {
		panic(ErrInvalidArgument)
	}

	return &api{
		repo:      repo,
		chunkSize: chunkSize,
	}
}

func (a *api) CreatePresentationV1(
	ctx context.Context,
	request *desc.CreatePresentationV1Request,
) (
	*desc.CreatePresentationV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreatePresentationV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "CreatePresentationV1").
		Uint64("LessonID", request.Presentation.LessonId).
		Uint64("UserID", request.Presentation.UserId).
		Str("Name", request.Presentation.Name).
		Str("Description", request.Presentation.Description).
		Msg("Input data")

	presentation := model.Presentation{
		LessonID:    request.Presentation.LessonId,
		UserID:      request.Presentation.UserId,
		Name:        request.Presentation.Name,
		Description: request.Presentation.Description,
	}

	id, err := a.repo.CreatePresentation(ctx, presentation)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.CreatePresentationV1Response{
		PresentationId: id,
	}

	return response, nil
}

func (a *api) MultiCreatePresentationsV1(
	ctx context.Context,
	request *desc.MultiCreatePresentationsV1Request,
) (
	*desc.MultiCreatePresentationsV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreatePresentationsV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "MultiCreatePresentationsV1").
		Int("length", len(request.Presentations)).
		Msg("Input data")

	presentations := make([]model.Presentation, a.chunkSize)
	var numberOfCreatedPresentations int64

	for i, n := uint(0), uint(len(request.Presentations)); i < n; i += a.chunkSize {
		end := i + a.chunkSize
		if end > n {
			end = n
		}

		number, err := func() (int64, error) {
			innerSpan := tracer.StartSpan(
				fmt.Sprintf("MultiCreatePresentationsV1: %d Bytes (B)", a.chunkSize*uint(unsafe.Sizeof(model.Presentation{}))),
				opentracing.ChildOf(span.Context()),
			)
			defer innerSpan.Finish()

			var j int
			for ; i < end; i, j = i+1, j+1 {
				presentations[j].LessonID = request.Presentations[i].LessonId
				presentations[j].UserID = request.Presentations[i].UserId
				presentations[j].Name = request.Presentations[i].Name
				presentations[j].Description = request.Presentations[i].Description
			}

			return a.repo.MultiCreatePresentations(ctx, presentations[:j])
		}()

		if err != nil {
			log.Error().Err(err).Msg("Failed to insert the data")
			return nil, status.Error(codes.ResourceExhausted, err.Error())
		}

		numberOfCreatedPresentations += number
	}

	response := &desc.MultiCreatePresentationsV1Response{
		NumberOfCreatedPresentations: numberOfCreatedPresentations,
	}

	return response, nil
}

func (a *api) UpdatePresentationV1(
	ctx context.Context,
	request *desc.UpdatePresentationV1Request,
) (
	*desc.UpdatePresentationV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdatePresentationV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "UpdatePresentationV1").
		Uint64("LessonID", request.Presentation.LessonId).
		Uint64("UserID", request.Presentation.UserId).
		Str("Name", request.Presentation.Name).
		Str("Description", request.Presentation.Description).
		Msg("Input data")

	presentation := model.Presentation{
		ID:          request.Presentation.Id,
		LessonID:    request.Presentation.LessonId,
		UserID:      request.Presentation.UserId,
		Name:        request.Presentation.Name,
		Description: request.Presentation.Description,
	}

	found, err := a.repo.UpdatePresentation(ctx, presentation)

	if err != nil {
		if err == repo.ErrPresentationNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to update the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.UpdatePresentationV1Response{
		Found: found,
	}

	return response, nil
}

func (a *api) DescribePresentationV1(
	ctx context.Context,
	request *desc.DescribePresentationV1Request,
) (
	*desc.DescribePresentationV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("DescribePresentationV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "DescribePresentationV1").
		Uint64("PresentationID", request.PresentationId).
		Msg("Input data")

	presentation, err := a.repo.DescribePresentation(ctx, request.PresentationId)
	if err != nil {
		if err == repo.ErrPresentationNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to describe the data")
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
	request *desc.ListPresentationsV1Request,
) (
	*desc.ListPresentationsV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ListPresentationsV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "ListPresentationsV1").
		Uint64("Limit", request.Limit).
		Uint64("Offset", request.Offset).
		Msg("Input data")

	presentations, err := a.repo.ListPresentations(ctx, request.Limit, request.Offset)
	if err != nil {
		if errors.Is(err, repo.ErrPresentationNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to fill the data")
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
	request *desc.RemovePresentationV1Request,
) (
	*desc.RemovePresentationV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("RemovePresentationV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "RemovePresentationV1").
		Uint64("PresentationID", request.PresentationId).
		Msg("Input data")

	removed, err := a.repo.RemovePresentation(ctx, request.PresentationId)
	response := &desc.RemovePresentationV1Response{
		Found: removed,
	}

	if err != nil {
		if errors.Is(err, repo.ErrPresentationNotFound) {
			return response, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to remove the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	return response, nil
}
