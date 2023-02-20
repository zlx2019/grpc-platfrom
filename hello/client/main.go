package main

import (
	"com.zero/grpc_hello/pb/member"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//1. 创建一个GRPC客户端,连接到服务端,关闭安全认证,否则无法连接
	conn, err := grpc.Dial("localhost:6565", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("gRPC Client Conn Server err: %s \n", err.Error())
		return
	}
	defer conn.Close()
	// 2. 创建HelloService的客户端
	//client := pb.NewHelloServiceClient(conn)
	memberServiceClient := member.NewMemberServiceClient(conn)
	memberRequest := &member.MemberRequest{Id: 1001}
	memberReply, err := memberServiceClient.GetMember(context.Background(), memberRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println(memberReply.UserName)
	// 3. 准备需要的参数
	//request := &pb.Request{
	//	//ImgType: &pb.Request_Jpg{"JPG格式"},
	//	ImgType: &pb.Request_Png{"Png格式"},
	//}
	// 3. 调用接口
	//response, err := client.SendMessage(context.Background(), request)
	//if err != nil {
	//	fmt.Printf("gRPC 接口调用失败 err: %s \n", err.Error())
	//	return
	//}
	//fmt.Printf("gRPC 接口调用成功 resp: %s", response.RespMessage)
}
