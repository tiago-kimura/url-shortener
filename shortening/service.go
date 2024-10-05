package shortening

//import "github.com/tiago-kimura/url-shortener/internal/hashEncoder"

type Service interface {
	ShortenUrl(url string) (string, error)
	GetUrlOrigin(urlId string) (UrlShortener, error)
	DeleteUrlShortener(urlId string) error
}

type ShorteningService struct {
	repository Repository
}

func NewShorteningService(repository Repository) ShorteningService {
	return ShorteningService{
		repository: repository,
	}
}

func (s ShorteningService) ShortenUrl(url string) (string, error) {
	//	urlId := hashEncoder.GenerateHashSHA(url)

	return "urlId", nil
}

func (s ShorteningService) GetUrlOrigin(urlId string) (UrlShortener, error) {
	return UrlShortener{}, nil
}

func (s ShorteningService) DeleteUrlShortener(urlId string) error {
	return nil
}
