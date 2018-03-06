package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"wshhz.com/grpc/Proto"
)

const (
	// 定义监听端口
	port = "41005"
)

// 数据结构
type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 启动服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}

	// 注册rpc服务
	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	fmt.Println("grpc server in:%s", port)
}

// 定义方法
func (t *Data) GetUser(ctx context.Context, request *inf.UserRq) (response *inf.UserRp, err error) {
	response = &inf.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}

	return
}
