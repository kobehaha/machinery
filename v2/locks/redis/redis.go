package redis

import (
	"errors"
	"strings"
	"time"
	"fmt"
	"math/rand"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/go-redis/redis/v8"
	"github.com/RichardKnop/machinery/v2/log"
)

var (
	ErrRedisLockFailed = errors.New("redis lock: failed to acquire lock")
)

type Lock struct {
	rclient  redis.UniversalClient
	retries  int
	interval time.Duration
}

func New(cnf *config.Config, addrs []string, db, retries int) Lock {
	if retries <= 0 {
		return Lock{}
	}
	lock := Lock{retries: retries}

	var password string

	parts := strings.Split(addrs[0], "@")
	if len(parts) == 2 {
		password = parts[0]
		addrs[0] = parts[1]
	}

	ropt := &redis.UniversalOptions{
		Addrs:    addrs,
		DB:       db,
		Password: password,
	}
	if cnf.Redis != nil {
		ropt.MasterName = cnf.Redis.MasterName
	}

	lock.rclient = redis.NewUniversalClient(ropt)

	return lock
}

func (r Lock) LockWithRetries(key string, unixTsToExpireNs int64) error {
	for i := 0; i <= r.retries; i++ {
		err := r.Lock(key, unixTsToExpireNs)
		if err == nil {
			//成功拿到锁，返回
			return nil
		} else {
			log.DEBUG.Println(fmt.Sprintf("retry %d to get lock failed, basic info is  %s", i, err.Error()))
		}
		r := rand.Intn(5)
		time.Sleep(time.Duration(r) * time.Second)
	}
	return ErrRedisLockFailed
}

func (r Lock) Lock(key string, unixTsToExpireNs int64) error {
	now := time.Now().UnixNano()
	expiration := time.Duration(unixTsToExpireNs + 1 - now)
	ctx := r.rclient.Context()

	success, err := r.rclient.SetNX(ctx, key, unixTsToExpireNs, expiration).Result()
	if err != nil {
		return err
	}
	if !success {
		ttl, err := r.rclient.TTL(ctx, key).Result()
		if err != nil {
			return err
		}
		log.DEBUG.Printf("set lock failed, current exist lock ttl is %d", ttl.Seconds())
		return ErrRedisLockFailed
	}
	return nil
}
