module github.com/NpoolPlatform/review-service

go 1.16

require (
	entgo.io/ent v0.11.2
	github.com/NpoolPlatform/api-manager v0.0.0-20220205130236-69d286e72dba
	github.com/NpoolPlatform/cloud-hashing-goods v0.0.0-20211224023221-a715aef93510
	github.com/NpoolPlatform/cloud-hashing-inspire v0.0.0-20211224023242-bd6da3111653
	github.com/NpoolPlatform/go-service-framework v0.0.0-20220120091626-4e8035637592
	github.com/NpoolPlatform/libent-cruder v0.0.0-20220801075201-cab5db8b6290
	github.com/NpoolPlatform/message v0.0.0-20220501030927-34f296682c0c
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.46.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
