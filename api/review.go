// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/review-service/message/npool"

	"github.com/NpoolPlatform/review-service/pkg/crud/review"
	mw "github.com/NpoolPlatform/review-service/pkg/middleware/review"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	myconst "github.com/NpoolPlatform/review-service/pkg/message/const"
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

func (s *Server) GetReviewsByAppDomain(ctx context.Context, in *npool.GetReviewsByAppDomainRequest) (*npool.GetReviewsByAppDomainResponse, error) {
	resp, err := review.GetByAppDomain(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get reviews by app domain error: %v", err)
		return &npool.GetReviewsByAppDomainResponse{}, status.Error(codes.Internal, "internal server error")
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

func (s *Server) SubmitReviewResult(ctx context.Context, in *npool.SubmitReviewResultRequest) (*npool.SubmitReviewResultResponse, error) {
	resp, err := mw.SubmitReviewResult(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("submit review result: %v", err)
		return &npool.SubmitReviewResultResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetReviewsByObjectIDs(ctx context.Context, in *npool.GetReviewsByObjectIDsRequest) (*npool.GetReviewsByObjectIDsResponse, error) {
	if in.GetObjectIDs() == nil {
		logger.Sugar().Error("GetReviewsByObjectIDs error: object ids can not be empty")
		return nil, status.Error(codes.InvalidArgument, "object ids can not be empty")
	}

	objectIDs := []uuid.UUID{}
	for _, objectID := range in.GetObjectIDs() {
		id, err := uuid.Parse(objectID)
		if err != nil {
			logger.Sugar().Errorf("GetReviewsByObjectIDs error: invalid objectID <%v>, %v", objectID, err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid objectID <%v>", objectID)
		}
		objectIDs = append(objectIDs, id)
	}

	ctx, cancel := context.WithTimeout(ctx, myconst.GrpcTimeout)
	defer cancel()

	resp, err := review.GetByObjectIDs(ctx, objectIDs)
	if err != nil {
		logger.Sugar().Errorf("GetReviewsByObjectIDs error: internal server error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &npool.GetReviewsByObjectIDsResponse{
		Infos: resp,
	}, nil
}
