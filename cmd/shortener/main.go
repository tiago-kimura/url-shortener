package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiago-kimura/url-shortener/config"
	"github.com/tiago-kimura/url-shortener/shortening"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	cfg := config.Load()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabase))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	repository := shortening.NewRepository(db)
	cache := shortening.NewRedisCache(rdb)

	minLengthRule := &shortening.MinLengthRule{MinLength: cfg.MinLenthToShorten}
	//hashExistsRule := &shortening.HashExistsRule{Repository: repository}
	validUrl := &shortening.ValidUrl{}
	rules := shortening.NewCompositeRule(minLengthRule, validUrl)
	urlService := shortening.NewShorteningService(repository, cache, *cfg, rules)

	urlHandler := NewHandler(&urlService)

	router := mux.NewRouter()
	router.HandleFunc("/shorten", urlHandler.ShortenURL).Methods("POST")
	router.HandleFunc("/getOriginalUrl/{urlId}", urlHandler.ResolveURL).Methods("GET")
	router.HandleFunc("/delete/{urlId}", urlHandler.DeleteURL).Methods("DELETE")

	log.Println("Starting server on :" + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
