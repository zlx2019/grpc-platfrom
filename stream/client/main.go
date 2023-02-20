package main

import (
	produce "com.zero/stream/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("gRPC Client Conn Server err: %s \n", err.Error())
		return
	}
	defer conn.Close()
	// 建立 grpcClient
	grpcClient := produce.NewProduceServiceClient(conn)
	streamClient, err := grpcClient.GetStream(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		reply, err := streamClient.Recv()
		if err != nil && err == io.EOF {
			// 服务的数据发送完毕,已关闭连接,直接退出即可
			break
		}
		fmt.Printf("接收到Server的数据, 内容: %s \n", reply.GetBody())
	}
}

func testGetStream() {
	conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("gRPC Client Conn Server err: %s \n", err.Error())
		return
	}
	defer conn.Close()
	// 建立 grpcClient
	grpcClient := produce.NewProduceServiceClient(conn)
	streamClient, err := grpcClient.GetStream(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		reply, err := streamClient.Recv()
		if err != nil && err == io.EOF {
			// 服务的数据发送完毕,已关闭连接,直接退出即可
			break
		}
		fmt.Printf("接收到Server的数据, 内容: %s \n", reply.GetBody())
	}
}

func testSendStream() {
	conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("gRPC Client Conn Server err: %s \n", err.Error())
		return
	}
	defer conn.Close()
	// 建立 grpcClient
	grpcClient := produce.NewProduceServiceClient(conn)
	// 发起请求,获取数据流客户端
	streamClient, err := grpcClient.SendStream(context.Background())
	// 流式发送十条消息到服务端
	i := 0
	for {
		// 1秒发送1条
		time.Sleep(time.Second)
		// 发送数据
		err = streamClient.Send(&produce.ProduceRequest{ProduceId: int32(i)})
		if err != nil {
			fmt.Println("发送数据失败...")
			continue
		}
		i++

		fmt.Printf("已发送完第 %d 条数据,发送时间: %v \n", int32(i), time.Now().Format("15:04:05"))
		if i >= 10 {
			// 当发送完数据后,响应给服务端 数据已经发送完毕,然后等服务端的结束响应
			fmt.Println("数据发送完毕,等待客户端响应...")
			reply, _ := streamClient.CloseAndRecv()
			fmt.Printf("客户端响应: 【ok: %v,msg: %s】 连接结束!", reply.GetCompleted(), reply.GetMsg())
			break
		}
	}
}
