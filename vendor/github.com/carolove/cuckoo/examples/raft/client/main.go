package main

import (
	"fmt"
	"os"

	"github.com/carolove/cuckoo"
	"github.com/carolove/cuckoo/model"
)

const (
	ADDR = "127.0.0.1:9090"
)

func main() {
	cli, err := cuckoo.GetRPCClient(ADDR)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	in := &model.RequestVoteRequest{}
	in.Vote = "hello world!@Client"
	out := cli.RequestVote(in)
	fmt.Println(out.Success, out.Msg)
}
