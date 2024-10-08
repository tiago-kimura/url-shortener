package shortening

import (
	"github.com/tiago-kimura/url-shortener/config"
	"github.com/tiago-kimura/url-shortener/internal/hashEncode"
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
	// TODO: remove http/https
	if len(url) > s.config.MaxLenthToShorten {
		urlId = hashEncode.GenerateHashSHA256(url, s.config.MaxLenthToShorten)
		urlShortener := UrlShortener{
			UrlId:       urlId,
			UrlOriginal: url,
		}
		err = s.rules.ProcessRules(urlShortener)
		if err != nil {
			return urlId, err
		}
		err = s.repository.PersistUrlShort(urlShortener)
	}

	return urlId, err
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
