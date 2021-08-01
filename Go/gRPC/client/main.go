package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := product.NewProductInfoClient(conn)
	ctx := context.Background() //

	id := AddProduct(ctx, client)
	GetProduct(ctx, client, id)
}

// 添加一个测试的商品
func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	aMac := &product.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	
}
