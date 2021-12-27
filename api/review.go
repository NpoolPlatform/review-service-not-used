// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/review-service/message/npool"

	"github.com/NpoolPlatform/review-service/pkg/crud/review"
	mw "github.com/NpoolPlatform/review-service/pkg/middleware/review"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReview(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	resp, err := review.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create review error: %v", err)
		return &npool.CreateReviewResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateReview(ctx context.Context, in *npool.UpdateReviewRequest) (*npool.UpdateReviewResponse, error) {
	resp, err := review.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update review error: %v", err)
		return &npool.UpdateReviewResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetReviewsByDomain(ctx context.Context, in *npool.GetReviewsByDomainRequest) (*npool.GetReviewsByDomainResponse, error) {
	resp, err := review.GetByDomain(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get reviews by domain error: %v", err)
		return &npool.GetReviewsByDomainResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) SubmitReview(ctx context.Context, in *npool.SubmitReviewRequest) (*npool.SubmitReviewResponse, error) {
	resp, err := mw.Submit(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("submit review: %v", err)
		return &npool.SubmitReviewResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
