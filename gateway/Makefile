# 生成go、go-grpc、grpc-gateway文件
build:
	protoc -I=./proto --go_out=./  \
		  	--go-grpc_out=./  \
		  	--grpc-gateway_out ./  \
		  	./proto/student.proto