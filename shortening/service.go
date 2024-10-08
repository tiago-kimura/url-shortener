package shortening

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/tiago-kimura/url-shortener/config"
	"github.com/tiago-kimura/url-shortener/internal/hashEncode"
)

const (
	https = "https://"
)

type Service interface {
	ShortenUrl(url string) (string, error)
	GetUrlOriginal(urlId string) (UrlShortener, error)
	DeleteUrlShortener(urlId string) error
}

type ShorteningService struct {
	repository Repository
	cache      RedisCache
	config     config.Config
	rules      *CompositeRule
}

func NewShorteningService(repository Repository, cache RedisCache, config config.Config, rules *CompositeRule) ShorteningService {
	return ShorteningService{
		repository: repository,
		cache:      cache,
		config:     config,
		rules:      rules,
	}
}

func (s ShorteningService) ShortenUrl(url string) (string, error) {
	urlId := ""
	var err error
	urlShortener := UrlShortener{
		UrlId:       urlId,
		UrlOriginal: url,
	}
	err = s.rules.ProcessRules(urlShortener)
	if err != nil {
		return urlId, err
	}
	urlSlice := strings.Split(url, https)
	urlId = hashEncode.GenerateHashMD5(urlSlice[1], s.config.MinLenthToShorten)
	err = s.isUrlIdExist(urlId)
	if err != nil {
		return urlId, err
	}
	urlShortener.UrlId = urlId
	err = s.repository.PersistUrlShort(urlShortener)

	return https + urlShortener.UrlId, err
}

func (s ShorteningService) GetUrlOriginal(urlId string) (UrlShortener, error) {
	urlShortener, err := s.repository.GetByUrlId(urlId)
	if err != nil {
		return UrlShortener{}, err
	}
	return urlShortener, nil
}

func (s ShorteningService) DeleteUrlShortener(urlId string) error {
	err := s.repository.DeleteByUrlId(urlId)
	if err != nil {
		return err
	}
	return nil
}

func (s ShorteningService) isUrlIdExist(urlId string) error {
	existUrl, err := s.repository.GetByUrlId(urlId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	if existUrl.UrlOriginal != "" {
		return errors.New("URL id already exists")
	}
	return nil
}
