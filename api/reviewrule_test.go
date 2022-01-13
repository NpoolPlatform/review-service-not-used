package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/review-service"
)

func assertReviewRule(t *testing.T, actual, expected *npool.ReviewRule) {
	assert.Equal(t, actual.ObjectType, expected.ObjectType)
	assert.Equal(t, actual.Domain, expected.Domain)
	assert.Equal(t, actual.Rules, expected.Rules)
}

func TestCreateReviewRule(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	reviewrule := npool.ReviewRule{
		ObjectType: "good",
		Domain:     fmt.Sprintf("cloud-hashing-goods-npool-top-%v", uuid.New().String()),
	}
	firstCreateInfo := npool.CreateReviewRuleResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateReviewRuleRequest{
			Info: &reviewrule,
		}).
		Post("http://localhost:50050/v1/create/review/rule")

	reviewrule.Rules = "{}"
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assertReviewRule(t, firstCreateInfo.Info, &reviewrule)
		}
	}

	reviewrule.ID = firstCreateInfo.Info.ID
	reviewrule.Rules = "{\"aaa\": \"bbbbb\"}"

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateReviewRuleRequest{
			Info: &reviewrule,
		}).
		Post("http://localhost:50050/v1/update/review/rule")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateReviewRuleResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertReviewRule(t, info.Info, &reviewrule)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetReviewRulesByDomainRequest{
			Domain: reviewrule.Domain,
		}).
		Post("http://localhost:50050/v1/get/review/rules/by/domain")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetReviewRulesByDomainResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, len(info.Infos), 1)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetReviewRuleRequest{
			ID: reviewrule.ID,
		}).
		Post("http://localhost:50050/v1/get/review/rule")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetReviewRuleResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertReviewRule(t, info.Info, &reviewrule)
		}
	}
}
