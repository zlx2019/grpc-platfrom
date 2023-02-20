package main

import (
	produce "com.zero/stream/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net"
	"time"
)

type ProduceServer struct {
	produce.UnimplementedProduceServiceServer
}

// SendStream 重写SendStream方法,服务提供者
func (*ProduceServer) SendStream(stream produce.ProduceService_SendStreamServer) error {
	// 循环读取客户端的数据流
	for {
		// 阻塞式等待客户端数据流
		request, err := stream.Recv()
		if err != nil && err == io.EOF {
			// 接收到 err = EOF 表示数据发送完成。响应客户端,跳出循环直接结束.并且关闭连接
			stream.SendAndClose(&produce.ProduceReply{Completed: true, Msg: "消费完毕"})
			fmt.Println("消费完成,关闭连接!")
			break
		}
		// 读取到数据流
		fmt.Printf("Server接收到数据,消费第 %d 条数据, 消费时间: %v\n", request.ProduceId, time.Now().Format("15:04:05"))
	}
	return nil
}

// GetStream 重写GetStream方法,提供服务
func (*ProduceServer) GetStream(e *emptypb.Empty, stream produce.ProduceService_GetStreamServer) error {
	i := 0
	for {
		i++
		body := fmt.Sprintf("这是来自于服务端的第 %d 条消息 产生时间: %s", i, time.Now().Format("15:04:05"))
		stream.Send(&produce.BodyReply{Body: body})
		time.Sleep(time.Second)
		if i >= 10 {
			// 发送数据完毕,注意：这里不会需要通知,直接返回nil即可
			break
		}
	}
	return nil
}

func main() {
	// 1.创建一个tcp服务
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer := grpc.NewServer()
	produceServer := &ProduceServer{}
	produce.RegisterProduceServiceServer(grpcServer, produceServer)
	fmt.Printf("gRPC Server Start listen: %s \n", listen.Addr())
	// 5.借助tpc服务 开启grpc服务器
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("GRPC服务启动失败 原因: %s \n", err.Error())
	}
}
