package validation

import (
	"espad_task/config"
	"espad_task/pkg/log"
	"espad_task/pkg/translate"
	"espad_task/service"
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
