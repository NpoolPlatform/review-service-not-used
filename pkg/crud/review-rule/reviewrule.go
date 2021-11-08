package reviewrule

import (
	"context"

	"github.com/NpoolPlatform/review-service/message/npool"

	"github.com/NpoolPlatform/review-service/pkg/db"
	"github.com/NpoolPlatform/review-service/pkg/db/ent"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/reviewrule"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func dbRowToReviewRule(row *ent.ReviewRule) *npool.ReviewRule {
	return &npool.ReviewRule{
		ID:         row.ID.String(),
		EntityType: row.EntityType,
		Domain:     row.Domain,
		Rules:      row.Rules,
	}
}

func Create(ctx context.Context, in *npool.CreateReviewRuleRequest) (*npool.CreateReviewRuleResponse, error) {
	info, err := db.Client().
		ReviewRule.
		Create().
		SetEntityType(in.GetInfo().GetEntityType()).
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

	info, err := db.Client().
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

	infos, err := db.Client().
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
	infos, err := db.Client().
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
