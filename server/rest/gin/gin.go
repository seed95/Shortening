package gin

import (
	"espad_task/config"
	"espad_task/pkg/log"
	"espad_task/pkg/translate"
	"espad_task/server/rest"
	"espad_task/service"
	"fmt"
	goGin "github.com/gin-gonic/gin"
)

type (
	handler struct {
		cfg        *config.Rest
		shortening service.Shortening
		logger     log.Logger
		translator translate.Translator
	}

	Option struct {
		Cfg        *config.Rest
		Shortening service.Shortening
		Logger     log.Logger
		Translator translate.Translator
	}
)

var r = goGin.Default()

func New(opt *Option) rest.Server {
	return &handler{
		cfg:        opt.Cfg,
		shortening: opt.Shortening,
		logger:     opt.Logger,
		translator: opt.Translator,
	}
}

func (s *handler) Start() error {
	s.setRoutes()
	return r.Run(fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
}

func (s *handler) setRoutes() {
	r.GET("/:key", s.Redirect)
	r.POST("/generate", s.GenerateShort)
	r.GET("/get/:key", s.GetOriginal)
}
