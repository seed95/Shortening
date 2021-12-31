package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/server"
)

func getLang(c *gin.Context) []translate.Language {
	return server.GetLanguage(c.GetHeader("Accept-Language"))
}
