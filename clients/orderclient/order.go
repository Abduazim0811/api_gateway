package orderclient

import (
	"Api/protos/orderproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcOrder() order.OrderServiceClient {
	conn, err := grpc.NewClient("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client:", err)
	}

	return order.NewOrderServiceClient(conn)
}
