package review

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/review-service"

	"github.com/NpoolPlatform/review-service/pkg/db"
	"github.com/NpoolPlatform/review-service/pkg/db/ent"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateReview(info *npool.Review) error {
	if _, err := uuid.Parse(info.GetObjectID()); err != nil {
		return xerrors.Errorf("invalid object id: %v", err)
	}
	if info.GetDomain() == "" {
		return xerrors.Errorf("invalid domain")
	}
	return nil
}

func dbRowToReview(row *ent.Review) *npool.Review {
	return &npool.Review{
		ID:         row.ID.String(),
		ObjectType: row.ObjectType,
		AppID:      row.AppID.String(),
		ReviewerID: row.ReviewerID.String(),
		State:      string(row.State),
		Message:    row.Message,
		ObjectID:   row.ObjectID.String(),
		Domain:     row.Domain,
		CreateAt:   row.CreateAt,
		Trigger:    row.Trigger,
	}
}

func Create(ctx context.Context, in *npool.CreateReviewRequest) (*npool.CreateReviewResponse, error) {
	if err := validateReview(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid patameter: %v", err)
	}

	appID, err := uuid.Parse(in.GetInfo().GetAppID())
	if err != nil {
		appID = uuid.UUID{}
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	id := uuid.New()
	if in.GetInfo().GetID() != "" {
		id, err = uuid.Parse(in.GetInfo().GetID())
		if err != nil {
			return nil, err
		}
	}
	info, err := cli.
		Review.
		Create().
		SetID(id).
		SetObjectType(in.GetInfo().GetObjectType()).
		SetState("wait").
		SetMessage("").
		SetReviewerID(uuid.UUID{}).
		SetObjectID(uuid.MustParse(in.GetInfo().GetObjectID())).
		SetAppID(appID).
		SetDomain(in.GetInfo().GetDomain()).
		SetTrigger(in.GetInfo().GetTrigger()).
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
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

func GetByAppDomain(ctx context.Context, in *npool.GetReviewsByAppDomainRequest) (*npool.GetReviewsByAppDomainResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Review.
		Query().
		Where(
			review.And(
				review.AppID(appID),
				review.Domain(in.GetDomain()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review: %v", err)
	}

	reviews := []*npool.Review{}
	for _, info := range infos {
		reviews = append(reviews, dbRowToReview(info))
	}

	return &npool.GetReviewsByAppDomainResponse{
		Infos: reviews,
	}, nil
}

func GetByDomain(ctx context.Context, in *npool.GetReviewsByDomainRequest) (*npool.GetReviewsByDomainResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Review.
		Query().
		Where(
			review.And(
				review.Domain(in.GetDomain()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review: %v", err)
	}

	reviews := []*npool.Review{}
	for _, info := range infos {
		reviews = append(reviews, dbRowToReview(info))
	}

	return &npool.GetReviewsByDomainResponse{
		Infos: reviews,
	}, nil
}

func GetByAppDomainObjectTypeID(ctx context.Context, in *npool.GetReviewsByAppDomainObjectTypeIDRequest) (*npool.GetReviewsByAppDomainObjectTypeIDResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	objectID, err := uuid.Parse(in.GetObjectID())
	if err != nil {
		return nil, xerrors.Errorf("invalid object id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Review.
		Query().
		Where(
			review.And(
				review.AppID(appID),
				review.Domain(in.GetDomain()),
				review.ObjectID(objectID),
				review.ObjectType(in.GetObjectType()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review: %v", err)
	}

	reviews := []*npool.Review{}
	for _, info := range infos {
		reviews = append(reviews, dbRowToReview(info))
	}

	return &npool.GetReviewsByAppDomainObjectTypeIDResponse{
		Infos: reviews,
	}, nil
}

func Get(ctx context.Context, in *npool.GetReviewRequest) (*npool.GetReviewResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		Review.
		Query().
		Where(
			review.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review: %v", err)
	}

	var _review *npool.Review
	for _, info := range infos {
		_review = dbRowToReview(info)
		break
	}

	return &npool.GetReviewResponse{
		Info: _review,
	}, nil
}
