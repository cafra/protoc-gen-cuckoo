// Code generated by protoc-gen-cuckoo 0.1, DO NOT EDIT.
// source: kv.proto

package model

import (
	"fmt"
	"log"

	etcdv3 "github.com/carolove/cuckoo/net/grpc/lb/local"
	"github.com/cuigh/auxo/config"
	"github.com/cuigh/auxo/util/lazy"
)

var (
	kv = &kvRPCClient{lazyValue: lazy.Value{New: createKV}}
)

func GetKV() KVClient {
	v, err := kv.lazyValue.Get()
	if err != nil {
		fmt.Println("KV service Get failed!")
		log.Fatalln("KVClient | Get | err=", err)
	}

	return v.(KVClient)
}

func createKV() (interface{}, error) {
	key := "rpc.client.kv"
	if !config.Exist(key) {
		log.Fatalln("createKV | Exist | key=", key)
	}

	conn, err := etcdv3.NewConn("protoc-gen-cuckoo/kv", config.GetString(key))
	if err != nil {
		log.Fatalln("createKV | NewConn | err=", err)
	}

	client := NewKVClient(conn)
	return interface{}(client), nil
}

type kvRPCClient struct {
	lazyValue lazy.Value
}
