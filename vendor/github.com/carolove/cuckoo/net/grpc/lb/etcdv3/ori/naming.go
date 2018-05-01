package grpclb

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc/naming"
)

// Register register service with name as prefix to etcd, multi etcd addr should use ; to split
func Register(etcdAddr, name string, addr string, ttl int64) error {
	var err error

	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			return err
		}
	}

	ticker := time.NewTicker(time.Second * time.Duration(ttl))

	r := &etcdnaming.GRPCResolver{Client: cli}
	go func() {
		for {
			r.Update(context.TODO(), "/"+schema+"/"+name, naming.Update{Op: naming.Add, Addr: addr})
			if err != nil {
				log.Println(err)
			}

			<-ticker.C
		}
	}()

	return nil
}

// UnRegister remove service from etcd
func UnRegister(name string, addr string) {

	if cli != nil {
		r := &etcdnaming.GRPCResolver{Client: cli}
		r.Update(context.TODO(), "/"+schema+"/"+name, naming.Update{Op: naming.Delete, Addr: addr})
	}
}
