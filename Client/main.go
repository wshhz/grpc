package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"wshhz.com/grpc/Proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	wg sync.WaitGroup
)

const (
	// 连接服务器地址
	server = "127.0.0.1"

	// 连接服务器端口
	port = "41005"

	// 连接并行度
	parallel = 5

	// 每连接请求次数
	times = 10
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	// 并行请求
	for i := 0; i < int(parallel); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exe()
		}()
	}

	wg.Wait()

	//exe()

	fmt.Printf("time taken: %.2f\r\n", time.Now().Sub(currTime).Seconds())
}

// 执行方法
func exe() {
	// 建立连接
	conn, err := grpc.Dial(server+":"+port, grpc.WithInsecure())
	if conn == nil || err != nil {
		fmt.Printf("connect error : %v\r\n", err)
		return
	}
	defer conn.Close()

	// 新建一个客户端对象
	client := inf.NewDataClient(conn)
	if client == nil {
		fmt.Printf("client = nil\r\n")
		return
	}

	for i := 0; i < int(times); i++ {
		getUser(client)
	}
}

// 请求方法
func getUser(client inf.DataClient) {
	var request inf.UserRq
	r := rand.Intn(parallel)
	request.Id = int32(r)

	ctx := context.Background()

	// 调用远程方法
	response, err := client.GetUser(ctx, &request)
	if err != nil {
		fmt.Printf("dial error :%v\r\n", err)
		return
	}

	// 判断返回结果是否正确
	if id, _ := strconv.Atoi(strings.Split(response.Name, ":")[0]); id != r {
		fmt.Printf("response error %#v", response)
	}

	// 输出远程调用的结果
	fmt.Printf("call success res = %v\r\n", response)
}
