// Package api implements a simple API for Ozon Code Platform Slide API.
package api

import (
	"context"
	"errors"
	"fmt"
	"unsafe"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/model"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/repo"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-slide-api"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedSlideAPIServer
	repo      repo.Repo
	chunkSize uint
}

func NewSlideAPI(repo repo.Repo, chunkSize uint) desc.SlideAPIServer {
	if chunkSize == 0 {
		panic("invalid argument")
	}

	return &api{
		repo:      repo,
		chunkSize: chunkSize,
	}
}

func (a *api) CreateSlideV1(
	ctx context.Context,
	request *desc.CreateSlideV1Request,
) (
	*desc.CreateSlideV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreateSlideV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "CreateSlideV1").
		Uint64("PresentationID", request.Slide.PresentationId).
		Uint64("Number", request.Slide.Number).
		Str("Type", model.ContentType(request.Slide.Type).String()).
		Msg("Input data")

	slide := model.Slide{
		PresentationID: request.Slide.PresentationId,
		Number:         request.Slide.Number,
		Type:           model.ContentType(request.Slide.Type),
	}

	id, err := a.repo.CreateSlide(ctx, slide)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.CreateSlideV1Response{
		SlideId: id,
	}

	return response, nil
}

func (a *api) MultiCreateSlidesV1(
	ctx context.Context,
	request *desc.MultiCreateSlidesV1Request,
) (
	*desc.MultiCreateSlidesV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateSlidesV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "MultiCreateSlidesV1").
		Int("length", len(request.Slides)).
		Msg("Input data")

	slides := make([]model.Slide, a.chunkSize)
	var numberOfCreatedSlides int64

	for i, n := uint(0), uint(len(request.Slides)); i < n; i += a.chunkSize {
		end := i + a.chunkSize
		if end > n {
			end = n
		}

		number, err := func() (int64, error) {
			innerSpan := tracer.StartSpan(
				fmt.Sprintf("MultiCreateSlidesV1: %d Bytes (B)", a.chunkSize*uint(unsafe.Sizeof(model.Slide{}))),
				opentracing.ChildOf(span.Context()),
			)
			defer innerSpan.Finish()

			var j int
			for ; i < end; i, j = i+1, j+1 {
				slides[j].PresentationID = request.Slides[i].PresentationId
				slides[j].Number = request.Slides[i].Number
				slides[j].Type = model.ContentType(request.Slides[i].Type)
			}

			return a.repo.MultiCreateSlides(ctx, slides[:j])
		}()

		if err != nil {
			log.Error().Err(err).Msg("Failed to insert the data")
			return nil, status.Error(codes.ResourceExhausted, err.Error())
		}

		numberOfCreatedSlides += number
	}

	response := &desc.MultiCreateSlidesV1Response{
		NumberOfCreatedSlides: numberOfCreatedSlides,
	}

	return response, nil
}

func (a *api) UpdateSlideV1(
	ctx context.Context,
	request *desc.UpdateSlideV1Request,
) (
	*desc.UpdateSlideV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdateSlideV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "UpdateSlideV1").
		Uint64("PresentationID", request.Slide.PresentationId).
		Uint64("Number", request.Slide.Number).
		Str("Type", model.ContentType(request.Slide.Type).String()).
		Msg("Input data")

	presentation := model.Slide{
		ID:             request.Slide.Id,
		PresentationID: request.Slide.PresentationId,
		Number:         request.Slide.Number,
		Type:           model.ContentType(request.Slide.Type),
	}

	found, err := a.repo.UpdateSlide(ctx, presentation)

	if err != nil {
		if err == repo.ErrSlideNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to update the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.UpdateSlideV1Response{
		Found: found,
	}

	return response, nil
}

func (a *api) DescribeSlideV1(
	ctx context.Context,
	request *desc.DescribeSlideV1Request,
) (
	*desc.DescribeSlideV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("DescribeSlideV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "DescribeSlideV1").
		Uint64("SlideID", request.SlideId).
		Msg("Input data")

	slide, err := a.repo.DescribeSlide(ctx, request.SlideId)
	if err != nil {
		if err == repo.ErrSlideNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to describe the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	response := &desc.DescribeSlideV1Response{
		Slide: &desc.Slide{
			Id:             slide.ID,
			PresentationId: slide.PresentationID,
			Number:         slide.Number,
			Type:           desc.ContentType(slide.Type),
		},
	}

	return response, nil
}

func (a *api) ListSlidesV1(
	ctx context.Context,
	request *desc.ListSlidesV1Request,
) (
	*desc.ListSlidesV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("ListSlidesV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "ListSlidesV1").
		Uint64("Limit", request.Limit).
		Uint64("Offset", request.Offset).
		Msg("Input data")

	slides, err := a.repo.ListSlides(ctx, request.Limit, request.Offset)
	if err != nil {
		if errors.Is(err, repo.ErrSlideNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to fill the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	list := make([]*desc.Slide, 0, len(slides))
	for _, slide := range slides {
		result := &desc.Slide{
			Id:             slide.ID,
			PresentationId: slide.PresentationID,
			Number:         slide.Number,
			Type:           desc.ContentType(slide.Type),
		}

		list = append(list, result)
	}

	response := &desc.ListSlidesV1Response{
		Slides: list,
	}

	return response, nil
}

func (a *api) RemoveSlideV1(
	ctx context.Context,
	request *desc.RemoveSlideV1Request,
) (
	*desc.RemoveSlideV1Response,
	error,
) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("RemoveSlideV1")
	defer span.Finish()

	if err := request.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Str("API", "RemoveSlideV1").
		Uint64("SlideID", request.SlideId).
		Msg("Input data")

	removed, err := a.repo.RemoveSlide(ctx, request.SlideId)
	response := &desc.RemoveSlideV1Response{
		Found: removed,
	}

	if err != nil {
		if errors.Is(err, repo.ErrSlideNotFound) {
			return response, status.Error(codes.NotFound, err.Error())
		}

		log.Error().Err(err).Msg("Failed to remove the data")
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	return response, nil
}
