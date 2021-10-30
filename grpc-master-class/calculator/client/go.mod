module calculator/client

go 1.16

replace calculator/protocol => ../protocol

require (
	google.golang.org/grpc v1.41.0
	calculator/protocol v0.0.0-00010101000000-000000000000
)
