package shortening

import (
	"time"
)

type UrlShortener struct {
	UrlId       string
	UrlOriginal string
	CreatedAt   time.Time
}
