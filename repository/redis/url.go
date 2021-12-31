package redis

import (
	"espad_task/build/messages"
	"espad_task/domain/url"
	"espad_task/pkg/derrors"
	"espad_task/pkg/log"
	"github.com/go-redis/redis/v8"
)

func (r *repository) AddUrl(url *url.Url) error {

	if err := r.rdb.Set(r.ctx, url.ShortLink, url.OriginalLink, url.Expiration).Err(); err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.redis",
			Function: "AddUrl",
			Params:   map[string]interface{}{"shortUrl": url},
			Message:  err.Error(),
		})

		return derrors.New(derrors.Unexpected, messages.DBError)
	}

	return nil
}

func (r *repository) ExistShortUrl(shortUrl string) (bool, error) {
	cmd := r.rdb.Exists(r.ctx, shortUrl)
	if cmd.Err() != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.redis",
			Function: "ExistShortUrl",
			Params:   map[string]interface{}{"short_url": shortUrl},
			Message:  cmd.Err().Error(),
		})

		return false, derrors.New(derrors.Unexpected, messages.DBError)
	}

	return cmd.Val() != 0, nil
}

func (r *repository) GetUrl(shortUrl string) (string, error) {

	cmd := r.rdb.Get(r.ctx, shortUrl)

	if cmd.Err() == redis.Nil {
		err := derrors.New(derrors.NotFound, messages.UrlNotFound)
		r.logger.Error(&log.Field{
			Section:  "repository.redis",
			Function: "GetUrl",
			Params:   map[string]interface{}{"short_url": shortUrl},
			Message:  r.translator.Translate(err.Error()),
		})

		return "", err
	}

	if cmd.Err() != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.redis",
			Function: "GetUrl",
			Params:   map[string]interface{}{"short_url": shortUrl},
			Message:  cmd.Err().Error(),
		})

		return "", derrors.New(derrors.Unexpected, messages.DBError)
	}

	return cmd.Val(), nil
}

func (r *repository) DeleteUrl(shortUrl string) error {

	cmd := r.rdb.Del(r.ctx, shortUrl)

	if cmd.Err() != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.redis",
			Function: "DeleteUrl",
			Params:   map[string]interface{}{"short_url": shortUrl},
			Message:  cmd.Err().Error(),
		})

		return derrors.New(derrors.Unexpected, messages.DBError)
	}

	return nil
}
