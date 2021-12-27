package review

import (
	"context"

	"github.com/NpoolPlatform/review-service/message/npool"
	crud "github.com/NpoolPlatform/review-service/pkg/crud/review"

	"golang.org/x/xerrors"
)

func Submit(ctx context.Context, in *npool.SubmitReviewRequest) (*npool.SubmitReviewResponse, error) {
	_, err := crud.Create(ctx, &npool.CreateReviewRequest{
		Info: in.GetInfo(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create review: %v", err)
	}

	// TODO: check if match automatically review
	return nil, nil
}
