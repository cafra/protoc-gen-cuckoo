package etcd

import (
	"github.com/cuigh/auxo/config"
	"github.com/cuigh/auxo/errors"
)

type Options struct {
	Type    string
	Address []string
}

func LoadOptions() (*Options, error) {
	key := "global.etcd"
	if !config.Exist(key) {
		return nil, errors.Format("can't find etcd config")
	}

	opts := &Options{}
	err := config.UnmarshalOption(key, opts)
	if err != nil {
		return nil, err
	}
	return opts, nil
}
