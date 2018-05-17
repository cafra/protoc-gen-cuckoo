package local

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

const schema = "wonamingv3"

type localResolver struct {
	defaultSrvAddr []string
	cc             resolver.ClientConn
}

// NewResolver initialize an etcd client
func NewResolver(defaultSrvAddr ...string) resolver.Builder {
	return &localResolver{defaultSrvAddr: defaultSrvAddr}
}

func (r *localResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {

	r.cc = cc
	var addrList []resolver.Address
	for _, addr := range r.defaultSrvAddr {
		addrList = append(addrList, resolver.Address{Addr: addr})
	}

	if len(addrList) == 0 {
		log.Fatalln("localResolver | Build | addrList is empty")
	}

	r.cc.NewAddress(addrList)

	return r, nil
}

func (r localResolver) Scheme() string {
	return schema
}

func (r localResolver) ResolveNow(rn resolver.ResolveNowOption) {
	log.Println("ResolveNow") // TODO check
}

// Close closes the resolver.
func (r localResolver) Close() {
	log.Println("Close")
}

// NewResolver initialize
// defaultSrvAddr
func NewConn(name string, defaultSrvAddr ...string) (conn *grpc.ClientConn, err error) {

	r := NewResolver(defaultSrvAddr...)
	resolver.Register(r)

	return grpc.Dial(r.Scheme()+"://author/"+name, grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
}
