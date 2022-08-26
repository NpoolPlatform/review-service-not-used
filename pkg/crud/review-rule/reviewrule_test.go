package reviewrule

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/review-service"
	testinit "github.com/NpoolPlatform/review-service/pkg/test-init" //nolint

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertReviewRule(t *testing.T, actual, expected *npool.ReviewRule) {
	assert.Equal(t, actual.ObjectType, expected.ObjectType)
	assert.Equal(t, actual.Domain, expected.Domain)
	assert.Equal(t, actual.Rules, expected.Rules)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	reviewrule := npool.ReviewRule{
		ObjectType: "good",
		Domain:     fmt.Sprintf("cloud-hashing-goods-npool-top-%v", uuid.New().String()),
	}

	resp, err := Create(context.Background(), &npool.CreateReviewRuleRequest{
		Info: &reviewrule,
	})
	reviewrule.Rules = "{}"

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertReviewRule(t, resp.Info, &reviewrule)
	}

	reviewrule.Rules = "{\"aaa\": \"bbbbbbb\"}"
	reviewrule.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateReviewRuleRequest{
		Info: &reviewrule,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertReviewRule(t, resp1.Info, &reviewrule)
	}

	resp2, err := GetByDomain(context.Background(), &npool.GetReviewRulesByDomainRequest{
		Domain: reviewrule.Domain,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := Get(context.Background(), &npool.GetReviewRuleRequest{
		ID: reviewrule.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertReviewRule(t, resp3.Info, &reviewrule)
	}
}
