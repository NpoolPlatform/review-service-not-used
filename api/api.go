package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/review-service"
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
