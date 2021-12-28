package review

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/review-service/message/npool"
	constant "github.com/NpoolPlatform/review-service/pkg/const" //nolint
	crud "github.com/NpoolPlatform/review-service/pkg/crud/review"

	goodsconst "github.com/NpoolPlatform/cloud-hashing-goods/pkg/message/const"
	inspireconst "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/message/const"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"google.golang.org/grpc"

	"golang.org/x/xerrors"
)

// For new review object, add get object method here
func getObject(conn *grpc.ClientConn, info *npool.Review) (interface{}, error) { //nolint
	switch info.Domain {
	case goodsconst.ServiceName:
		return nil, nil
	case inspireconst.ServiceName:
		return nil, nil
	}
	return nil, xerrors.Errorf("unknown review object")
}

// For new review object, add automatic object rule check here
func checkAutomaticRule(conn *grpc.ClientConn, info *npool.Review, obj interface{}) (autoReviewed bool, result string, err error) { //nolint
	switch info.Domain {
	case goodsconst.ServiceName:
		return false, constant.StateWait, nil
	case inspireconst.ServiceName:
		return false, constant.StateWait, nil
	}
	return false, constant.StateRejected, xerrors.Errorf("unknown review object")
}

func Submit(ctx context.Context, in *npool.SubmitReviewRequest) (*npool.SubmitReviewResponse, error) {
	info := in.GetInfo()

	conn, err := grpc2.GetGRPCConn(info.Domain, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get connection: %v", err)
	}
	defer conn.Close()

	obj, err := getObject(conn, info)
	if err != nil {
		return nil, xerrors.Errorf("fail get object: %v", err)
	}

	resp, err := crud.Create(ctx, &npool.CreateReviewRequest{
		Info: info,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail create review: %v", err)
	}

	auto, result, err := checkAutomaticRule(conn, info, obj)
	if err != nil {
		return nil, xerrors.Errorf("fail check automatic rule: %v", err)
	}
	if auto {
		resp.Info.State = result
		resp.Info.Message = fmt.Sprintf("Automatically review %v", result)

		resp1, err := crud.Update(ctx, &npool.UpdateReviewRequest{
			Info: resp.Info,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail update review state: %v", err)
		}
		// Do not need submit result here, invoker will get result by return
		info = resp1.Info
	}

	return &npool.SubmitReviewResponse{
		Info: info,
	}, nil
}

func submitGoodReviewResult(conn *grpc.ClientConn, info *npool.Review, obj interface{}) error {
	// TODO: submit good review result
	return nil
}

func submitInspireReviewResult(conn *grpc.ClientConn, info *npool.Review, obj interface{}) error {
	// TODO: submit inspire review result
	return nil
}

// For new review object, add submit result to object service here
func submitReviewResult(conn *grpc.ClientConn, info *npool.Review, obj interface{}) error {
	switch info.Domain {
	case goodsconst.ServiceName:
		return submitGoodReviewResult(conn, info, obj)
	case inspireconst.ServiceName:
		return submitInspireReviewResult(conn, info, obj)
	}
	return xerrors.Errorf("unknown review object")
}

func SubmitReviewResult(ctx context.Context, in *npool.SubmitReviewResultRequest) (*npool.SubmitReviewResultResponse, error) {
	info := in.GetInfo()

	conn, err := grpc2.GetGRPCConn(info.Domain, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get connection: %v", err)
	}
	defer conn.Close()

	obj, err := getObject(conn, info)
	if err != nil {
		return nil, xerrors.Errorf("fail get object: %v", err)
	}

	resp, err := crud.Update(ctx, &npool.UpdateReviewRequest{
		Info: info,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail update review state: %v", err)
	}

	if err := submitReviewResult(conn, info, obj); err != nil {
		return nil, xerrors.Errorf("fail submit review result: %v", err)
	}

	return &npool.SubmitReviewResultResponse{
		Info: resp.Info,
	}, nil
}
