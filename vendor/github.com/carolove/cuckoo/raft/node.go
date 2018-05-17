package cuckoo

import (
	"sync"

	"github.com/carolove/cuckoo/usi"
	"github.com/cuigh/auxo/config"
)

type node struct {
	srv        *server
	msgCh      chan *usi.Msg
	masterAddr string
	lock       sync.Mutex // lock only use for set master addr
}

func NewNode() (usi.USinterface, error) {
	n := &node{}
	config := config.GetAllServer()
	srv, err := newRPCServer(config.Port)
	if err != nil {
		return nil, err
	}

	n.srv = srv
	// init msgCh as nonblocking
	n.msgCh = make(chan *usi.Msg, 10)
	n.lock = sync.Mutex{}
	return n, nil
}

// server未start并且master未选择返回错误
func (nd *node) GetMasterAddr() string {
	return nd.masterAddr
}

// rugosa未初始化send返回错误
func (nd *node) SendMsg(msg *usi.Msg) error {
	if msg != nil {
		nd.msgCh <- msg
	}
	return nil
}

// 一旦创建server便可以获取channel接口
func (nd *node) GetMsgCh() chan *usi.Msg {
	return nd.msgCh
}

func (nd *node) Start() error {
	return nd.srv.start()
}

func (nd *node) Stop() error {
	return nd.srv.stop()
}
