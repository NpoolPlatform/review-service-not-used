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

	"github.com/NpoolPlatform/review-service/message/npool"
)

func assertReview(t *testing.T, actual, expected *npool.Review) {
	assert.Equal(t, actual.EntityType, expected.EntityType)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.Message, expected.Message)
	assert.Equal(t, actual.ObjectID, expected.ObjectID)
	assert.Equal(t, actual.Domain, expected.Domain)
}

func TestCreateReview(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	review := npool.Review{
		EntityType: "good",
		State:      "wait",
		ObjectID:   uuid.New().String(),
		Domain:     fmt.Sprintf("cloud-hashing-goods-npool-top-%v", uuid.New().String()),
	}
	firstCreateInfo := npool.CreateReviewResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateReviewRequest{
			Info: &review,
		}).
		Post("http://localhost:35759/v1/create/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.New())
			assertReview(t, firstCreateInfo.Info, &review)
		}
	}

	review.ID = firstCreateInfo.Info.ID
	review.State = "approved"
	review.ReviewerID = uuid.New().String()
	review.Message = "Good good good"

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateReviewRequest{
			Info: &review,
		}).
		Post("http://localhost:35759/v1/update/review")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateReviewResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertReview(t, info.Info, &review)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetReviewsByDomainRequest{
			Domain: review.Domain,
		}).
		Post("http://localhost:35759/v1/get/reviews/by/domain")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetReviewsByDomainResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, len(info.Infos), 1)
		}
	}
}
