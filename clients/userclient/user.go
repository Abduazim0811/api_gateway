package userclient

import (
	"Api/protos/userproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcUser() userproto.UserServiceClient {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client:", err)
	}

	return userproto.NewUserServiceClient(conn)
}
