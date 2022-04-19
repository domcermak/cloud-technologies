package main

import (
	"domcermak/ctc/assignments/05/common"
	"fmt"
	"net"

	"domcermak/ctc/assignments/05/etcd"
	"domcermak/ctc/assignments/05/server"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting...")

	flags := server.ParseFlags()
	fmt.Println(flags)

	etcdClient, err := etcd.NewClient(flags.EtcdServerAddr)
	if err != nil {
		panic(err)
	}

	test(etcdClient)

	s := server.NewServer(etcdClient)
	grpcServer := grpc.NewServer()
	server.RegisterEtcdServer(grpcServer, s)
	lis, err := net.Listen("tcp", flags.ServerAddr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("listening on %s...\n", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println("quitting...")
}

func test(etcdClient *etcd.Client) {
	_, _ = common.LogRequest("test", func() (interface{}, error) {
		res, err := etcdClient.Post("foo", "bar")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

		res, err = etcdClient.Get("foo")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

		res, err = etcdClient.Delete("foo")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

		res, err = etcdClient.Get("foo")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

		return nil, nil
	})
}
