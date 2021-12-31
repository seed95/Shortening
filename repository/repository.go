package repository

import "espad_task/domain/url"

type (
	Repository interface {
		Url
	}

	Url interface {
		AddUrl(url *url.Url) error
		ExistShortUrl(shortUrl string) (bool, error)
		GetUrl(shortUrl string) (string, error)
		DeleteUrl(shortUrl string) error
	}
)
