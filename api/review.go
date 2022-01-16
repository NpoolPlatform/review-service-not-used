// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/review-service"

	"github.com/NpoolPlatform/review-service/pkg/crud/review"
	mw "github.com/NpoolPlatform/review-service/pkg/middleware/review"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReview(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	resp, err := review.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create review error: %v", err)
		return &npool.CreateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateReview(ctx context.Context, in *npool.UpdateReviewRequest) (*npool.UpdateReviewResponse, error) {
	resp, err := review.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update review error: %v", err)
		return &npool.UpdateReviewResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetReviewsByDomain(ctx context.Context, in *npool.GetReviewsByDomainRequest) (*npool.GetReviewsByDomainResponse, error) {
	resp, err := review.GetByDomain(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get reviews by domain error: %v", err)
		return &npool.GetReviewsByDomainResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetReviewsByAppDomain(ctx context.Context, in *npool.GetReviewsByAppDomainRequest) (*npool.GetReviewsByAppDomainResponse, error) {
	resp, err := review.GetByAppDomain(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get reviews by app domain error: %v", err)
		return &npool.GetReviewsByAppDomainResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetReviewsByAppDomainObjectTypeID(ctx context.Context, in *npool.GetReviewsByAppDomainObjectTypeIDRequest) (*npool.GetReviewsByAppDomainObjectTypeIDResponse, error) {
	resp, err := review.GetByAppDomainObjectTypeID(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get reviews by app domain object type id error: %v", err)
		return &npool.GetReviewsByAppDomainObjectTypeIDResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) SubmitReview(ctx context.Context, in *npool.SubmitReviewRequest) (*npool.SubmitReviewResponse, error) {
	resp, err := mw.Submit(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("submit review: %v", err)
		return &npool.SubmitReviewResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) SubmitReviewResult(ctx context.Context, in *npool.SubmitReviewResultRequest) (*npool.SubmitReviewResultResponse, error) {
	resp, err := mw.SubmitReviewResult(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("submit review result: %v", err)
		return &npool.SubmitReviewResultResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
