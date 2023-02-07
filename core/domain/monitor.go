package domain

type Monitor struct {
	Count           int64 `json:"count"`
	UniqueUserAgent int64 `json:"unique_user_agent"`
}
