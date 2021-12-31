package validation

import (
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/log"
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/service"
)

type (
	handler struct {
		cfg        *config.Application
		logger     log.Logger
		translator translate.Translator
	}

	Option struct {
		Cfg        *config.Application
		Logger     log.Logger
		Translator translate.Translator
	}
)

func New(opt *Option) service.Validation {

	return &handler{
		cfg:        opt.Cfg,
		logger:     opt.Logger,
		translator: opt.Translator,
	}
}
