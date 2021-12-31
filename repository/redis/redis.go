package redis

import (
	"context"
	"espad_task/config"
	"espad_task/pkg/log"
	"espad_task/pkg/translate"
	repo "espad_task/repository"
	"fmt"
	goRedis "github.com/go-redis/redis/v8"
)

type (
	repository struct {
		cfg        *config.Redis
		ctx        context.Context
		rdb        *goRedis.Client
		logger     log.Logger
		translator translate.Translator
	}

	Option struct {
		Cfg        *config.Redis
		Logger     log.Logger
		Translator translate.Translator
	}
)

func New(opt *Option) repo.Repository {

	rdb := goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%s:%d", opt.Cfg.Host, opt.Cfg.Port),
		Password: opt.Cfg.Password,
		DB:       opt.Cfg.DbIndex,
	})

	return &repository{
		cfg:        opt.Cfg,
		ctx:        context.Background(),
		rdb:        rdb,
		logger:     opt.Logger,
		translator: opt.Translator,
	}
}
