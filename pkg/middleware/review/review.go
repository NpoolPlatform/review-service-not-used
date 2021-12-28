package review

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/review-service/message/npool"
	crud "github.com/NpoolPlatform/review-service/pkg/crud/review"

	"golang.org/x/xerrors"
)

func Submit(ctx context.Context, in *npool.SubmitReviewRequest) (*npool.SubmitReviewResponse, error) {
	info := in.GetInfo()

	_, err := crud.Create(ctx, &npool.CreateReviewRequest{
		Info: info,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create review: %v", err)
	}

	// TODO: get object from service
	logger.Sugar().Infof("info ==> %v", info)

	// TODO: check if match automatically review

	// TODO: if automatically, notify service

	return &npool.SubmitReviewResponse{
		Info: info,
	}, nil
}
