package domain

type Url struct {
	id        int
	long_url  string
	short_url string
}

func CreateNewUrl() *Url {
	return &Url{id: 1, long_url: "https://google.com", short_url: "googl"}
}
