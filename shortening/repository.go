package shortening

import "database/sql"

type Repository interface {
	PersistUrlShort(url UrlShortener) error
	GetByUrlId(urlShort string) (UrlShortener, error)
	DeleteByUrlId(urlShort string) error
}

type RepositoryImpl struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{Db: db}
}

func (r *RepositoryImpl) PersistUrlShort(url UrlShortener) error {
	stmt, err := r.Db.Prepare("INSERT INTO url_shortener (url_id, url_original) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url.UrlId, url.UrlOriginal)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) GetByUrlId(urlId string) (UrlShortener, error) {
	var urlShortener UrlShortener
	stmt, err := r.Db.Prepare("Select * from url_shortener WHERE url_id =(?)")
	if err != nil {
		return urlShortener, err
	}

	err = stmt.QueryRow(urlId).Scan(&urlShortener.UrlId, &urlShortener.UrlOriginal, &urlShortener.CreatedAt)
	if err != nil {
		return urlShortener, err
	}
	return urlShortener, nil
}

func (r *RepositoryImpl) DeleteByUrlId(urlId string) error {
	stmt, err := r.Db.Prepare("DELETE url_shortener WHERE url_id =(?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(urlId)
	return err
}
