package url

import "time"

type (
	Url struct {
		OriginalLink string
		ShortLink    string
		Expiration   time.Duration
	}
)
