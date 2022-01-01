package validation

import (
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/log/logrus"
	"github.com/seed95/shortening/pkg/translate/i18n"
	"github.com/seed95/shortening/service"
	"testing"
)

var (
	serviceTest service.Validation
)

func setupService(t *testing.T) {

	cfg := &config.Application{
		AliasMinLength: 6,
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

	serviceTest = New(&Option{
		Cfg:        cfg,
		Logger:     logger,
		Translator: translator,
	})
}
