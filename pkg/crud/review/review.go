package review

import (
	"context"

	"github.com/NpoolPlatform/review-service/message/npool"

	"github.com/NpoolPlatform/review-service/pkg/db"
	"github.com/NpoolPlatform/review-service/pkg/db/ent"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateReview(info *npool.Review) error {
	if _, err := uuid.Parse(info.GetObjectID()); err != nil {
		return xerrors.Errorf("invalid object id: %v", err)
	}
	return nil
}

func dbRowToReview(row *ent.Review) *npool.Review {
	return &npool.Review{
		ID:         row.ID.String(),
		EntityType: row.EntityType,
		ReviewerID: row.ReviewerID.String(),
		State:      string(row.State),
		Message:    row.Message,
		ObjectID:   row.ObjectID.String(),
		Domain:     row.Domain,
	}
}

func Create(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	if err := validateReview(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid patameter: %v", err)
	}

	info, err := db.Client().
		Review.
		Create().
		SetEntityType(in.GetInfo().GetEntityType()).
		SetState("wait").
		SetMessage("").
		SetReviewerID(uuid.UUID{}).
		SetObjectID(uuid.MustParse(in.GetInfo().GetObjectID())).
		SetDomain(in.GetInfo().GetDomain()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create review: %v", err)
	}

	return &npool.CreateReviewResponse{
		Info: dbRowToReview(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateReviewRequest) (*npool.UpdateReviewResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	reviewerID, err := uuid.Parse(in.GetInfo().GetReviewerID())
	if err != nil {
		return nil, xerrors.Errorf("invalid reviewer id: %v", err)
	}

	info, err := db.Client().
		Review.
		UpdateOneID(id).
		SetState(review.State(in.GetInfo().GetState())).
		SetMessage(in.GetInfo().GetMessage()).
		SetReviewerID(reviewerID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update review: %v", err)
	}

	return &npool.UpdateReviewResponse{
		Info: dbRowToReview(info),
	}, nil
}

func GetReviewsByDomain(ctx context.Context, in *npool.GetReviewsByDomainRequest) (*npool.GetReviewsByDomainResponse, error) {
	return nil, nil
}
