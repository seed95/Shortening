package service

type (
	Shortening interface {
		GenerateShort(originalLink, alias, expirationStr string) (string, error)
		GetOriginalLink(key string) (string, error)
		GetShortLink(key string) string
	}

	Validation interface {
		Alias(alias string) error
	}
)
