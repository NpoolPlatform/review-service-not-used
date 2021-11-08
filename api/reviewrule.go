package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/review-service/message/npool"

	"github.com/NpoolPlatform/review-service/pkg/crud/review-rule" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateReviewRule(ctx context.Context, in *npool.CreateReviewRuleRequest) (*npool.CreateReviewRuleResponse, error) {
	resp, err := reviewrule.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("create review rule error: %w", err)
		return &npool.CreateReviewRuleResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateReviewRule(ctx context.Context, in *npool.UpdateReviewRuleRequest) (*npool.UpdateReviewRuleResponse, error) {
	resp, err := reviewrule.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update review rule error: %w", err)
		return &npool.UpdateReviewRuleResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetReviewRule(ctx context.Context, in *npool.GetReviewRuleRequest) (*npool.GetReviewRuleResponse, error) {
	resp, err := reviewrule.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("update review rule error: %w", err)
		return &npool.GetReviewRuleResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetReviewRulesByDomain(ctx context.Context, in *npool.GetReviewRulesByDomainRequest) (*npool.GetReviewRulesByDomainResponse, error) {
	resp, err := reviewrule.GetByDomain(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get review rules by domain error: %w", err)
		return &npool.GetReviewRulesByDomainResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
