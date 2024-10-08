package shortening

import (
	"time"
)

type UrlShortener struct {
	UrlId       string    `gorm:"primaryKey;index;size:10"`
	UrlOriginal string    `gorm:"not null"`
	CreateAt    time.Time `gorm:"autoCreateTime"`
}
