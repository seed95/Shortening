package gin

import (
	"fmt"
	goGin "github.com/gin-gonic/gin"
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/log"
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/server/rest"
	"github.com/seed95/shortening/service"
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
