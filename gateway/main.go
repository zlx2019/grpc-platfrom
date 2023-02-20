package main

import (
	"context"
	"fmt"
	"gateway/pb/student"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"net/http"
	"strconv"
)

type StudentService struct {
	student.UnimplementedStudentServiceServer
}

// Students 获取学生列表
func (*StudentService) Students(ctx context.Context, e *emptypb.Empty) (*student.StudentListReply, error) {
	i := 1
	listReply := &student.StudentListReply{}
	for {
		istr := strconv.Itoa(i)
		u := "username" + istr
		p := "password" + istr
		e := "email" + istr
		student := &student.StudentReply{Id: uint64(i), Username: u, Password: p, Email: e}
		listReply.Students = append(listReply.Students, student)
		i++
		if i > 10 {
			break
		}
	}
	//return nil, status.Errorf(codes.Unimplemented, "method Students not implemented")
	return listReply, nil
}

// 根据id获取学生信息
func (*StudentService) GetStudent(ctx context.Context, req *student.QueryRequest) (*student.StudentReply, error) {
	fmt.Printf("根据ID获取用户信息: %d \n", req.Id)
	return &student.StudentReply{Id: 1001, Username: "awdwadaw", Password: "wadawda", Email: "awdwada"}, nil
}

// 保存用户信息
func (*StudentService) SaveStudent(ctx context.Context, req *student.SaveRequest) (*emptypb.Empty, error) {
	fmt.Printf("保存用户信息 : %v", req)
	return &emptypb.Empty{}, nil
}

func main() {
	go StartGRPC()
	go StartGateway()
	select {}

}

// StartGateway 启动Gateway
func StartGateway() {
	// 创建grpc的client,指定服务端地址
	clientConn, err := grpc.DialContext(context.Background(),
		"127.0.0.1:8888", grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// 创建一个对外开放的HTTP服务
	mux := runtime.NewServeMux()

	// 创建gateway 服务
	gatewayServer := &http.Server{Addr: ":9999", Handler: mux}

	// 注册服务映射处理
	_ = student.RegisterStudentServiceHandler(context.Background(), mux, clientConn)
	// 启动服务
	err = gatewayServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func StartGRPC() {
	// 创建tcp服务
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 注册服务
	student.RegisterStudentServiceServer(grpcServer, &StudentService{})
	// 启动grpc服务
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
