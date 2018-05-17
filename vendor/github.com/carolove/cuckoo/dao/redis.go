package dao

import (
	"os"

	"github.com/cafra/utils/db"
	"github.com/cuigh/auxo/config"
	"github.com/cuigh/auxo/log"
	"github.com/cuigh/auxo/util/lazy"
)

var (
	redisValue0 = lazy.Value{New: createRedis0}
	redisValue1 = lazy.Value{New: createRedis1}
	redisValue2 = lazy.Value{New: createRedis2}
)

func GetLoanRedis(node int) *db.RedisDao {
	var v interface{}
	var err error
	switch node {
	case 1:
		v, err = redisValue1.Get()
	case 2:
		v, err = redisValue2.Get()
	default:
		v, err = redisValue0.Get()
	}

	if err != nil {
		log.Get("default").Errorf("GetLoanRedis | node=%d | redisValue.Get() | open loan redis failed, err=%v", node, err)
		os.Exit(0)
	}
	return v.(*db.RedisDao)
}

func createRedis0() (d interface{}, err error) {
	redisDao, err := db.NewRedisDao(config.GetString("db.redis.cm_yn_loan0.address"), false)
	if err != nil {
		log.Get("default").Errorf("initMysql |loan mysql err=%v", err)
		os.Exit(0)
	}
	d = interface{}(redisDao)
	return d, err
}

func createRedis1() (d interface{}, err error) {
	redisDao, err := db.NewRedisDao(config.GetString("db.redis.cm_yn_loan1.address"), false)
	if err != nil {
		log.Get("default").Errorf("initMysql |loan mysql err=%v", err)
		os.Exit(0)
	}
	d = interface{}(redisDao)
	return d, err
}

func createRedis2() (d interface{}, err error) {
	redisDao, err := db.NewRedisDao(config.GetString("db.redis.cm_yn_loan2.address"), false)
	if err != nil {
		log.Get("default").Errorf("initMysql |loan mysql err=%v", err)
		os.Exit(0)
	}
	d = interface{}(redisDao)
	return d, err
}
