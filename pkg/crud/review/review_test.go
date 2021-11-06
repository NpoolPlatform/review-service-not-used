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

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	review := npool.Review{
		EntityType: "good",
		ReviewerID: uuid.New().String(),
		State:      "wait",
		Message:    "Invalid good",
		ObjectID:   uuid.New().String(),
		Domain:     "cloud-hashing-goods-npool-top",
	}

	_, err := Create(context.Background(), &npool.CreateReviewRequest{
		Info: &review,
	})
	assert.Nil(t, err)
}
