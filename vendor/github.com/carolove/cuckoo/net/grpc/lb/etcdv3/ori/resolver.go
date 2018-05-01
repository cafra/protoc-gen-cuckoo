package grpclb

import (
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

const schema = "wonamingv3"

var cli *clientv3.Client

type etcdResolver struct {
	rawAddr string
	cc      resolver.ClientConn
}

// NewResolver initialize an etcd client
func NewConn(etcdAddr, name string) (conn *grpc.ClientConn, err error) {
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			return nil, err
		}
	}

	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)

	return grpc.Dial("/"+schema+"/"+name, grpc.WithBalancer(b), grpc.WithInsecure())
}
