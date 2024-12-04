module s

replace f => ./f

go 1.21

toolchain go1.23.4

require (
	f v0.0.0-00010101000000-000000000000
	knative.dev/func-go v0.20.0
)

require (
	github.com/fillmore-labs/name-service v0.0.0-20230823210441-7d06e47aa1d5 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.30.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.57.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
