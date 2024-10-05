package shortening

import (
	"time"
)

type UrlShortener struct {
	UrlId     string
	UrlOrigin string
	CreateAt  time.Time
	UpdateAt  time.Time
}
