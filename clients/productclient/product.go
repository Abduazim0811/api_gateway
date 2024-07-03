package productclient

import (
	"Api/protos/productproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcProduct() productproto.ProductsClient {
	conn, err := grpc.NewClient("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client:", err)
	}

	return productproto.NewProductsClient(conn)
}
