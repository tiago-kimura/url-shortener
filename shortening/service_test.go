package shortening

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tiago-kimura/url-shortener/config"
	"github.com/tiago-kimura/url-shortener/internal/hashEncode"
)

type MockRepository struct {
	mock.Mock
}

type MockRedisCache struct {
	mock.Mock
}

func (m *MockRepository) PersistUrlShort(urlShortener UrlShortener) error {
	args := m.Called(urlShortener)
	return args.Error(0)
}

func (m *MockRepository) GetByUrlId(urlId string) (UrlShortener, error) {
	args := m.Called(urlId)
	return args.Get(0).(UrlShortener), args.Error(1)
}

func (m *MockRepository) DeleteByUrlId(urlId string) error {
	args := m.Called(urlId)
	return args.Error(0)
}

func (m *MockRedisCache) Get(urlId string) (string, error) {
	args := m.Called(urlId)
	return args.String(0), args.Error(1)
}

func (m *MockRedisCache) Set(urlId, url string, ttl time.Duration) error {
	args := m.Called(urlId, url, ttl)
	return args.Error(0)
}

func (m *MockRedisCache) Delete(urlId string) error {
	args := m.Called(urlId)
	return args.Error(0)
}

func TestShortenUrl_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockCache := new(MockRedisCache)
	cfg := config.Config{MinLenthToShorten: 6}
	rules := new(CompositeRule)

	service := NewShorteningService(mockRepo, mockCache, cfg, rules)

	url := "https://mercadolivre.com/promocoes"
	urlId := hashEncode.GenerateHashMD5("mercadolivre.com/promocoes", cfg.MinLenthToShorten)
	urlShortener := UrlShortener{
		UrlId:       urlId,
		UrlOriginal: url,
	}

	mockRepo.On("GetByUrlId", urlId).Return(UrlShortener{}, sql.ErrNoRows)
	mockRepo.On("PersistUrlShort", urlShortener).Return(nil)

	shortenedUrl, err := service.ShortenUrl(url)

	assert.NoError(t, err)
	assert.Equal(t, https+urlId, shortenedUrl)
	mockRepo.AssertExpectations(t)
}

func TestShortenUrl_UrlIdExists(t *testing.T) {
	mockRepo := new(MockRepository)
	mockCache := new(MockRedisCache)
	cfg := config.Config{MinLenthToShorten: 6}
	rules := new(CompositeRule)

	service := NewShorteningService(mockRepo, mockCache, cfg, rules)

	url := "https://mercadolivre.com/promocoes"
	urlId := hashEncode.GenerateHashMD5("mercadolivre.com/promocoes", cfg.MinLenthToShorten)
	urlShortener := UrlShortener{
		UrlId:       urlId,
		UrlOriginal: url,
	}

	mockRepo.On("GetByUrlId", urlId).Return(urlShortener, nil)

	shortenedUrl, err := service.ShortenUrl(url)

	assert.Error(t, err)
	assert.Equal(t, urlId, shortenedUrl)
	mockRepo.AssertExpectations(t)
}

func TestGetUrlOriginal_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockCache := new(MockRedisCache)
	cfg := config.Config{}
	rules := new(CompositeRule)

	service := NewShorteningService(mockRepo, mockCache, cfg, rules)

	urlId := "bed8a7f5"
	urlShortener := UrlShortener{
		UrlId:       urlId,
		UrlOriginal: "mercadolivre.com/promocoes",
	}

	mockRepo.On("GetByUrlId", urlId).Return(urlShortener, nil)

	result, err := service.GetUrlOriginal(urlId)

	assert.NoError(t, err)
	assert.Equal(t, urlShortener, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUrlOriginal_NotFound(t *testing.T) {
	mockRepo := new(MockRepository)
	mockCache := new(MockRedisCache)
	cfg := config.Config{}
	rules := new(CompositeRule)

	service := NewShorteningService(mockRepo, mockCache, cfg, rules)

	urlId := "bed8a7f5"

	mockRepo.On("GetByUrlId", urlId).Return(UrlShortener{}, sql.ErrNoRows)

	result, err := service.GetUrlOriginal(urlId)

	assert.Error(t, err)
	assert.Equal(t, UrlShortener{}, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUrlShortener_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	mockCache := new(MockRedisCache)
	cfg := config.Config{}
	rules := new(CompositeRule)

	service := NewShorteningService(mockRepo, mockCache, cfg, rules)

	urlId := "bed8a7f5"

	mockRepo.On("DeleteByUrlId", urlId).Return(nil)

	err := service.DeleteUrlShortener(urlId)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUrlShortener_Error(t *testing.T) {
	mockRepo := new(MockRepository)
	mockCache := new(MockRedisCache)
	cfg := config.Config{}
	rules := new(CompositeRule)

	service := NewShorteningService(mockRepo, mockCache, cfg, rules)

	urlId := "bed8a7f5"

	mockRepo.On("DeleteByUrlId", urlId).Return(errors.New("delete error"))

	err := service.DeleteUrlShortener(urlId)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
