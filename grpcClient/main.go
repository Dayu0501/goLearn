package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcClient/services"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prodRes.ProdStock)

	prodResName, err := prodClient.GetProdName(context.Background(), &services.ProdRequestName{Name: "hello"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(prodResName.Name)
}
