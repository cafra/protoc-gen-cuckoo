package cuckoo

import (
	"errors"

	"github.com/carolove/cuckoo/model"
	"github.com/cuigh/auxo/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	SERVICE = "MessageService"
)

var clients map[string]*client

func init() {
	clients = make(map[string]*client)
	addAllClient()
}

func addAllClient() {
	clientsConfig := config.GetAllClients()
	for _, config := range clientsConfig {
		NewRPCClient(config.Addr + config.Port)
	}
}

type client struct {
	conn         *grpc.ClientConn
	wapperClient model.MessageServiceClient
}

func NewRPCClient(addr string) error {
	c := &client{}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.wapperClient = model.NewMessageServiceClient(conn)
	clients[addr] = c
	return nil
}

func GetRPCClient(addr string) (*client, error) {
	if c, isExist := clients[addr]; isExist {
		return c, nil
	}
	return nil, errors.New("系统不存在该client")
}

func RemoveRPCClient(addr string) error {
	if client, isExist := clients[addr]; isExist {
		delete(clients, addr)
		client.close()
	}
	return errors.New("未找到对应的client")
}

func (c *client) close() error {
	return c.conn.Close()
}

func (c *client) AppendEntries(in *model.AppendEntriesRequest) (out *model.AppendEntriesAck) {
	out, err := c.wapperClient.AppendEntries(context.Background(), in)
	if err != nil {
		return nil
	}
	return out
}

func (c *client) RequestVote(in *model.RequestVoteRequest) (out *model.RequestVoteAck) {
	out, err := c.wapperClient.RequestVote(context.Background(), in)
	if err != nil {
		return nil
	}
	return out
}
