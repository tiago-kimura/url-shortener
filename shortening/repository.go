package shortening

import (
	"gorm.io/gorm"
)

type Repository interface {
	PersistUrlShort(url UrlShortener) error
	GetByUrlId(urlShort string) (UrlShortener, error)
	DeleteByUrlId(urlShort string) error
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) PersistUrlShort(url UrlShortener) error {
	return r.db.Create(&url).Error
}

func (r *RepositoryImpl) GetByUrlId(urlId string) (UrlShortener, error) {
	var urlShortener UrlShortener
	if err := r.db.First(&urlShortener, "url_id = ?", urlId).Error; err != nil {
		return urlShortener, err
	}
	return urlShortener, nil
}

func (r *RepositoryImpl) DeleteByUrlId(urlId string) error {
	return r.db.Delete(&UrlShortener{}, "url_id = ?", urlId).Error
}
