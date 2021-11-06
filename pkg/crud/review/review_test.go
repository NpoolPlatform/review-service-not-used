package review

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/review-service/message/npool"
	"github.com/NpoolPlatform/review-service/pkg/test-init" //nolint

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

func assertReview(t *testing.T, actual, expected *npool.Review) {
	assert.Equal(t, actual.EntityType, expected.EntityType)
	assert.Equal(t, actual.State, expected.State)
	assert.Equal(t, actual.Message, expected.Message)
	assert.Equal(t, actual.ObjectID, expected.ObjectID)
	assert.Equal(t, actual.Domain, expected.Domain)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	review := npool.Review{
		EntityType: "good",
		State:      "wait",
		ObjectID:   uuid.New().String(),
		Domain:     "cloud-hashing-goods-npool-top",
	}

	resp, err := Create(context.Background(), &npool.CreateReviewRequest{
		Info: &review,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assert.Equal(t, resp.Info.ReviewerID, uuid.UUID{}.String())
		assertReview(t, resp.Info, &review)
	}

	review.State = "approved"
	review.ReviewerID = uuid.New().String()
	review.Message = "Good good good"
	review.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateReviewRequest{
		Info: &review,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.Equal(t, resp1.Info.ReviewerID, review.ReviewerID)
		assertReview(t, resp1.Info, &review)
	}
}
