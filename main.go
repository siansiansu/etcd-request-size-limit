package main

import (
	"context"
	"fmt"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	ctx := context.Background()

	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:          []string{"127.0.0.1:2379"},
		MaxCallSendMsgSize: 2 * 1024 * 1024,
		MaxCallRecvMsgSize: 3 * 1024 * 1024,
	})

	defer cli.Close()

	fmt.Println("====================================TEST 1====================================")
	_, err := cli.Put(ctx, "foo", strings.Repeat("a", 1.5*1024*1024+100))
	if err != nil {
		fmt.Println("client writes exceeding --max-request-bytes will be rejected from etcd server, because 1.5*1024*1024+100 > 1.5*1024*1024")
	}

	fmt.Println("====================================TEST 2====================================")
	_, err = cli.Put(ctx, "foo", strings.Repeat("a", 5*1024*1024))
	if err != nil {
		fmt.Println("client writes exceeding MaxCallSendMsgSize will be rejected from client-side, because 5*1024*1024 > 2*1024*1024")
	}

	fmt.Println("====================================TEST 3====================================")

	for i := range []int{0, 1, 2, 3, 4} {
		_, err = cli.Put(ctx, fmt.Sprintf("foo%d", i), strings.Repeat("a", 1.5*1024*1024-100))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("key foo1,foo2,foo3,foo4 is put successfully.")

	fmt.Println("====================================TEST 4====================================")
	_, err = cli.Get(ctx, "foo", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("client reads exceeding MaxCallRecvMsgSize will be rejected from client-side")
	}
}
