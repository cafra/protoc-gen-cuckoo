package main

import (
	"testing"

	"github.com/carolove/cuckoo"
	"github.com/carolove/cuckoo/model"
	"github.com/cuigh/auxo/test/assert"
)

func TestGetAllClient(t *testing.T) {
	client, _ := cuckoo.GetRPCClient("127.0.0.1:9090")
	assert.NotEmpty(t, client)
	req := &model.RequestVoteRequest{}
	req.Vote = "hello"
	resp := client.RequestVote(req)
	assert.Equal(t, resp.Msg, "hello world!@server")
	cuckoo.RemoveRPCClient("127.0.0.1:9090")
}
