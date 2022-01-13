package reviewrule

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/review-service"

	"github.com/NpoolPlatform/review-service/pkg/db"
	"github.com/NpoolPlatform/review-service/pkg/db/ent"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/reviewrule"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func dbRowToReviewRule(row *ent.ReviewRule) *npool.ReviewRule {
	return &npool.ReviewRule{
		ID:         row.ID.String(),
		ObjectType: row.ObjectType,
		Domain:     row.Domain,
		Rules:      row.Rules,
	}
}

func Create(ctx context.Context, in *npool.CreateReviewRuleRequest) (*npool.CreateReviewRuleResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		ReviewRule.
		Create().
		SetObjectType(in.GetInfo().GetObjectType()).
		SetDomain(in.GetInfo().GetDomain()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create review: %v", err)
	}

	return &npool.CreateReviewRuleResponse{
		Info: dbRowToReviewRule(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateReviewRuleRequest) (*npool.UpdateReviewRuleResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		ReviewRule.
		UpdateOneID(id).
		SetRules(in.GetInfo().GetRules()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update review: %v", err)
	}

	return &npool.UpdateReviewRuleResponse{
		Info: dbRowToReviewRule(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetReviewRuleRequest) (*npool.GetReviewRuleResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		ReviewRule.
		Query().
		Where(
			reviewrule.And(
				reviewrule.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review rule: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty review rules")
	}

	return &npool.GetReviewRuleResponse{
		Info: dbRowToReviewRule(infos[0]),
	}, nil
}

func GetByDomain(ctx context.Context, in *npool.GetReviewRulesByDomainRequest) (*npool.GetReviewRulesByDomainResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		ReviewRule.
		Query().
		Where(
			reviewrule.And(
				reviewrule.Domain(in.GetDomain()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query review: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty review")
	}

	rules := []*npool.ReviewRule{}
	for _, info := range infos {
		rules = append(rules, dbRowToReviewRule(info))
	}

	return &npool.GetReviewRulesByDomainResponse{
		Infos: rules,
	}, nil
}
