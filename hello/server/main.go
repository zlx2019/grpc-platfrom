package main

import (
	"com.zero/grpc_hello/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// HelloServer GO服务端需要定义一个结构体与 Proto中的Service一一对应
// 然后实现proto编译后生成的一个结构体,然后重写它的方法
type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

// SendMessage 重写proto中定义的rpc SendMessage方法
func (server *HelloServer) SendMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// 获取到请求参数中的Message
	imgType := ""
	jpg := req.GetJpg()
	png := req.GetPng()
	if jpg != "" {
		imgType = jpg
	} else {
		imgType = png
	}
	fmt.Printf("收到gRPC Client 请求, message: %s  \n", imgType)
	// 然后构建一个响应参数
	respMsg := &pb.Response{RespMessage: "gRpc Server的返回参数"}
	// 然后将参数返回,并且没有err,所以直接返回nil
	return respMsg, nil
}
func main() {
	// 1.创建一个tcp服务
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2.创建一个grpc服务
	grpcServer := grpc.NewServer()
	// 3.创建实现完成的接口服务
	helloService := &HelloServer{}
	// 4.将接口服务注册到grpc服务中
	pb.RegisterHelloServiceServer(grpcServer, helloService)
	fmt.Printf("gRPC Server Start listen: %s \n", listen.Addr())
	// 5.借助tpc服务 开启grpc服务器
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("GRPC服务启动失败 原因: %s \n", err.Error())
	}
}
