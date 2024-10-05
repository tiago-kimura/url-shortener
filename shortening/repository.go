package shortening

import "database/sql"

type Repository interface {
	PersistUrlShort(url UrlShortener) error
	GetUrlOrigin(urlShort string) (UrlShortener, error)
	DeleteUrlShort(urlShort string) error
}

type RepositoryImpl struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{Db: db}
}

func (r *RepositoryImpl) PersistUrlShort(url UrlShortener) error {
	stmt, err := r.Db.Prepare("INSERT INTO url_shortener (url_id, url_origin, create_at, update_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url.UrlId, url.UrlOrigin, url.CreateAt, url.UpdateAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) GetUrlOrigin(urlId string) (UrlShortener, error) {
	var urlShort UrlShortener
	err := r.Db.QueryRow("Select * from url_shortener where url_id = " + urlId).Scan(&urlShort)
	if err != nil {
		return urlShort, err
	}
	return urlShort, nil
}

func (r *RepositoryImpl) DeleteUrlShort(urlId string) error {
	stmt, err := r.Db.Prepare("DELETE url_shortener where url_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(urlId)
	if err != nil {
		return err
	}
	return nil
}
