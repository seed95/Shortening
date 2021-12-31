package gin

import (
	"espad_task/pkg/translate"
	"espad_task/server"
	"github.com/gin-gonic/gin"
)

func getLang(c *gin.Context) []translate.Language {
	return server.GetLanguage(c.GetHeader("Accept-Language"))
}
