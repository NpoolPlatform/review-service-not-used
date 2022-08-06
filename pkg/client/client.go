package review

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/review-service"

	constant "github.com/NpoolPlatform/review-service/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ReviewServiceClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewReviewServiceClient(conn)

	return handler(_ctx, cli)
}

func GetDomainReviews(ctx context.Context, appID, domain, objectType string) ([]*npool.Review, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ReviewServiceClient) (cruder.Any, error) {
		resp, err := cli.GetReviewsByAppDomain(ctx, &npool.GetReviewsByAppDomainRequest{
			AppID:  appID,
			Domain: domain,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.Review), nil
}
