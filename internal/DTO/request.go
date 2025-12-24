package dto

type Request struct {
	Url         string `json:"url"`
	CustomAlias string `json:"custom_alias`
	ExpiresIn   int    `json:"expires_in"`
}
