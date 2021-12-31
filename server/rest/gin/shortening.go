package gin

import (
	"espad_task/build/messages"
	"espad_task/pkg/derrors"
	"espad_task/pkg/log"
	"espad_task/server/params"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func (s *handler) Redirect(c *gin.Context) {

	lang := getLang(c)

	key := c.Param("key")
	originalLink, err := s.shortening.GetOriginalLink(key)
	if err != nil {
		msg, code := derrors.HttpError(err)
		c.JSON(code, s.translator.Translate(msg, lang...))
		return
	}

	redirectLink := url.URL{Path: originalLink}
	c.Redirect(http.StatusFound, redirectLink.RequestURI())

}

func (s *handler) GenerateShort(c *gin.Context) {

	lang := getLang(c)

	reqParam := new(params.GenerateShortRequest)

	if err := c.Bind(reqParam); err != nil {
		s.logger.Error(&log.Field{
			Section:  "handler.user",
			Function: "createUser",
			Message:  s.translator.Translate(err.Error()),
		})

		c.JSON(http.StatusBadRequest, s.translator.Translate(messages.ParseQueryError, lang...))
		return
	}

	shortLink, err := s.shortening.GenerateShort(reqParam.OriginalLink, reqParam.Alias, reqParam.Expiration)
	if err != nil {
		msg, code := derrors.HttpError(err)

		c.JSON(code, s.translator.Translate(msg, lang...))
		return
	}

	res := &params.GenerateShortResponse{
		OriginalLink: reqParam.OriginalLink,
		ShortLink:    shortLink,
		Expiration:   reqParam.Expiration,
	}

	c.JSON(http.StatusOK, res)
}

func (s *handler) GetOriginal(c *gin.Context) {
	lang := getLang(c)

	key := c.Param("key")

	originalLink, err := s.shortening.GetOriginalLink(key)
	if err != nil {
		msg, code := derrors.HttpError(err)
		c.JSON(code, s.translator.Translate(msg, lang...))
		return
	}

	res := &params.GetOriginalResponse{
		OriginalLink: originalLink,
		Key:          key,
	}

	c.JSON(http.StatusOK, res)

}
