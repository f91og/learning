package main

import (
	"context"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*product.Product
}

// 添加商品
func (s *server) AddProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}

	s.productMap[req.Id] = req
	resp.Value = req.Id
	return // AddProduct函数的返回值列表里指定了参数名字，所以这里return的时候就可以省略参数名直接return
}

// 获取商品
func (s *server) GetProduct(ctx context.Context, req *product.ProductId) (resp *product.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*product.Product)
	}
	resp = s.productMap[req.Value] //
	return
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("net listen err ", err)
		return
	}
	s := grpc.NewServer()                           //
	product.RegisterProductInfoServer(s, &server{}) //
	log.Println("start gRPC listen on port 50051")

	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}

}
