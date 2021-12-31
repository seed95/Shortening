package redis

import (
	"context"
	"fmt"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/log/logrus"
	"github.com/seed95/shortening/pkg/translate/i18n"
	"testing"
)

var (
	repoTest *repository
)

func setupTest(t *testing.T) {
	cfg := &config.Redis{
		Host: "localhost",
		Port: 6379,
	}

	opt := &logrus.Option{
		Path:         "../../../logs/test",
		Pattern:      "%Y-%m-%dT%H:%M",
		RotationSize: "20MB",
		RotationTime: "24h",
		MaxAge:       "720h",
	}
	logger, err := logrus.New(opt)
	if err != nil {
		t.Fatal(err)
	}

	translator, err := i18n.New("../../../build/i18n")
	if err != nil {
		t.Fatal(err)
	}

	rdb := goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",
		DB:       1, //Test data base
	})

	repoTest = &repository{
		cfg:        cfg,
		ctx:        context.Background(),
		rdb:        rdb,
		logger:     logger,
		translator: translator,
	}

	repoTest.rdb.FlushDB(repoTest.ctx)
}

func teardownTest() {
	repoTest = nil
}
