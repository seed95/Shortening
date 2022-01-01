package params

type (
	GenerateShortRequest struct {
		OriginalLink string `json:"original_link"`
		Alias        string `json:"alias"`
		Expiration   string `json:"expiration"`
	}

	GenerateShortResponse struct {
		OriginalLink string `json:"original_link"`
		ShortLink    string `json:"short_link"`
		Expiration   string `json:"expiration"`
	}
)

type (
	GetOriginalResponse struct {
		OriginalLink string `json:"original_link"`
		Key          string `json:"key"`
	}
)
