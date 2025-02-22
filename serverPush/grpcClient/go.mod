module github.com/mehulsuthar-000/grpcClient

go 1.23.4

require (
	github.com/mehulsuthar-000/serverPush/protofiles v0.0.0
	google.golang.org/grpc v1.70.0
)

require (
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/mehulsuthar-000/serverPush/protofiles => ../protofiles
