package shortening

import (
	"github.com/seed95/shortening/build/messages"
	"github.com/seed95/shortening/domain/url"
	"github.com/seed95/shortening/pkg/derrors"
	"github.com/seed95/shortening/pkg/log"
	"github.com/xhit/go-str2duration/v2"
	"time"
)

func (h *handler) GenerateShort(originalLink, alias, expirationStr string) (string, error) {

	expiration := h.defaultExpire

	if len(expirationStr) != 0 {
		var err error
		expiration, err = str2duration.ParseDuration(expirationStr)
		if err != nil {
			h.logger.Error(&log.Field{
				Section:  "service.shortening",
				Function: "GenerateShort",
				Params:   map[string]interface{}{"expiration": expirationStr},
				Message:  "",
			})

			return "", derrors.New(derrors.Invalid, messages.InvalidExpiration)
		}
	}

	if len(alias) != 0 {
		return h.addShort(alias, originalLink, expiration)
	}

	return h.generateShort(originalLink, expiration)
}

func (h *handler) generateShort(originalLink string, expiration time.Duration) (string, error) {

	hashLink := hash(originalLink)
	key := encode(hashLink)
	urlObject := &url.Url{
		OriginalLink: originalLink,
		ShortLink:    key,
		Expiration:   expiration,
	}
	if err := h.urlRepo.AddUrl(urlObject); err != nil {
		return "", err
	}

	return h.baseUrl + key, nil
}

func (h *handler) addShort(alias, originalLink string, expiration time.Duration) (string, error) {

	if err := h.validation.Alias(alias); err != nil {
		return "", err
	}

	exist, err := h.urlRepo.ExistShortUrl(alias)
	if err != nil {
		return "", err
	}

	if exist {
		h.logger.Error(&log.Field{
			Section:  "service.shortening",
			Function: "GenerateShorterLink",
			Params:   map[string]interface{}{"alias": alias},
			Message:  h.translator.Translate(messages.ExistAlias),
		})

		return "", derrors.New(derrors.Invalid, messages.ExistAlias)

	}

	urlObject := &url.Url{
		OriginalLink: originalLink,
		ShortLink:    alias,
		Expiration:   expiration,
	}

	if err := h.urlRepo.AddUrl(urlObject); err != nil {
		return "", err
	}
	return h.baseUrl + alias, nil
}

func (h *handler) GetOriginalLink(key string) (string, error) {

	originalUrl, err := h.urlRepo.GetUrl(key)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (h *handler) GetShortLink(key string) string {
	return h.baseUrl + key
}
