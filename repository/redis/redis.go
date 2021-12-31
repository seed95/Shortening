package redis

import (
	"context"
	"fmt"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/log"
	"github.com/seed95/shortening/pkg/translate"
	repo "github.com/seed95/shortening/repository"
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
