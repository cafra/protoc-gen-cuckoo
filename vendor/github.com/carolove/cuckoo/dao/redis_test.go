package dao_test

import (
	"fmt"
	"testing"

	"github.com/carolove/cuckoo/dao"
	"github.com/cuigh/auxo/config"
)

func init() {
	config.AddFolder("../config")
}

func Test_Redis(t *testing.T) {
	client := dao.GetLoanRedis()
	client.Set("test:raft:1", "yes")
	fmt.Println(client.Get("test:raft:1"))
}
