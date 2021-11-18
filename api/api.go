package api

import (
	"context"

	"github.com/NpoolPlatform/review-service/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedReviewServiceServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterReviewServiceServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterReviewServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
