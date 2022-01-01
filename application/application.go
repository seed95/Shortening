package application

import (
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/log"
	"github.com/seed95/shortening/pkg/log/logrus"
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/pkg/translate/i18n"
	"github.com/seed95/shortening/repository/redis"
	shorteningServer "github.com/seed95/shortening/server/rpc/shortening"
	"github.com/seed95/shortening/service/shortening"
	"github.com/seed95/shortening/service/validation"
)

var cfg = &config.Config{}

type (
	Option struct {
		ConfigFile string
	}
)

func Run(opt *Option) error {

	if err := initConfig(opt.ConfigFile); err != nil {
		return err
	}

	logger, err := initLog()
	if err != nil {
		return err
	}

	translator, err := initTranslator()
	if err != nil {
		logger.Error(&log.Field{
			Section:  "application",
			Function: "Run",
			Message:  err.Error(),
		})

		return err
	}

	repo := redis.New(&redis.Option{
		Cfg:        &cfg.Database.Redis,
		Logger:     logger,
		Translator: translator,
	})

	validationService := validation.New(&validation.Option{
		Cfg:        &cfg.Application,
		Logger:     logger,
		Translator: translator,
	})

	shorteningService, err := shortening.New(&shortening.Option{
		Cfg:        cfg,
		UrlRepo:    repo,
		Validation: validationService,
		Logger:     logger,
		Translator: translator,
	})
	if err != nil {
		return err
	}

	return shorteningServer.Start(&shorteningServer.Option{
		Cfg:               &cfg.Server,
		ShorteningService: shorteningService,
		Logger:            logger,
		Translator:        translator,
	})
}

func initConfig(configFile string) error {
	return config.Parse(configFile, cfg)
}

func initLog() (log.Logger, error) {
	return logrus.New(&logrus.Option{
		Path:         cfg.Logger.Logrus.Path,
		Pattern:      cfg.Logger.Logrus.Pattern,
		RotationSize: cfg.Logger.Logrus.RotationSize,
		RotationTime: cfg.Logger.Logrus.RotationTime,
		MaxAge:       cfg.Logger.Logrus.MaxAge,
	})
}

func initTranslator() (translate.Translator, error) {
	return i18n.New(cfg.Translator.I18n.MessagePath)
}
