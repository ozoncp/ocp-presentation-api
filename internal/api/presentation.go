// Package api implements a simple API for Ozon Code Platform Presentation API.
package api

import (
	"context"

	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-presentation-api"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errPresentationNotFound = "presentation not found"
)

type api struct {
	desc.UnimplementedPresentationAPIServer
}

func NewPresentationAPI() desc.PresentationAPIServer {
	return &api{}
}

func (a *api) CreatePresentationV1(
	ctx context.Context,
	req *desc.CreatePresentationV1Request,
) (
	*desc.CreatePresentationV1Response,
	error,
) {
	log.Print(errPresentationNotFound)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errPresentationNotFound)

	return nil, err
}

func (a *api) DescribePresentationV1(
	ctx context.Context,
	req *desc.DescribePresentationV1Request,
) (
	*desc.DescribePresentationV1Response,
	error,
) {
	log.Print(errPresentationNotFound)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errPresentationNotFound)
	return nil, err
}

func (a *api) ListPresentationV1(
	ctx context.Context,
	req *desc.ListPresentationsV1Request,
) (
	*desc.ListPresentationsV1Response,
	error,
) {
	log.Print(errPresentationNotFound)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errPresentationNotFound)
	return nil, err
}

func (a *api) RemovePresentationV1(
	ctx context.Context,
	req *desc.RemovePresentationV1Request,
) (
	*desc.RemovePresentationV1Response,
	error,
) {
	log.Print(errPresentationNotFound)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errPresentationNotFound)
	return nil, err
}
