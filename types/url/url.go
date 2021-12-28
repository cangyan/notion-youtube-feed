package url

type Url struct {
	Url string `json:"url,omitempty"`
}

func ValueOf(s string) Url {
	return Url{Url: s}
}
