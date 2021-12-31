package shortening

import (
	"espad_task/config"
	"espad_task/pkg/log"
	"espad_task/pkg/translate"
	"espad_task/repository"
	"espad_task/service"
	"fmt"
	"github.com/xhit/go-str2duration/v2"
	"time"
)

type (
	handler struct {
		urlRepo       repository.Url
		validation    service.Validation
		logger        log.Logger
		translator    translate.Translator
		defaultExpire time.Duration
		baseUrl       string
	}

	Option struct {
		Cfg        *config.Config
		UrlRepo    repository.Url
		Validation service.Validation
		Logger     log.Logger
		Translator translate.Translator
	}
)

func New(opt *Option) (service.Shortening, error) {

	expireDuration, err := str2duration.ParseDuration(opt.Cfg.Application.Expire)
	if err != nil {
		opt.Logger.Error(&log.Field{
			Section:  "service.shortening",
			Function: "New",
			Params:   map[string]interface{}{"expire": opt.Cfg.Application.Expire},
			Message:  err.Error(),
		})
		return nil, err
	}

	return &handler{
		urlRepo:       opt.UrlRepo,
		validation:    opt.Validation,
		logger:        opt.Logger,
		translator:    opt.Translator,
		defaultExpire: expireDuration,
		baseUrl:       fmt.Sprintf("%v:%v/", opt.Cfg.Server.Rest.Host, opt.Cfg.Server.Rest.Port),
	}, nil
}
