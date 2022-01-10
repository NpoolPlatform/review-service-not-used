module github.com/NpoolPlatform/review-service

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/cloud-hashing-goods v0.0.0-20211224023221-a715aef93510
	github.com/NpoolPlatform/cloud-hashing-inspire v0.0.0-20211224023242-bd6da3111653
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211222114515-4928e6cf3f1f
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0
