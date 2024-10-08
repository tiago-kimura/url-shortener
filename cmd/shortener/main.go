package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiago-kimura/url-shortener/config"
	"github.com/tiago-kimura/url-shortener/shortening"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
)

func main() {

	cfg := config.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	repository := shortening.NewRepository(db)
	cache := shortening.NewRedisCache(rdb)

	maxLengthRule := &shortening.MaxLengthRule{MaxLength: cfg.MaxLenthToShorten}
	hashExistsRule := &shortening.HashExistsRule{Repository: repository}
	rules := shortening.NewCompositeRule(maxLengthRule, hashExistsRule)
	urlService := shortening.NewShorteningService(repository, *cache, *cfg, rules)

	urlHandler := NewHandler(&urlService)

	router := mux.NewRouter() // TODO: separe route
	router.HandleFunc("/shorten", urlHandler.ShortenURL).Methods("POST")
	router.HandleFunc("/resolve/{urlId}", urlHandler.ResolveURL).Methods("GET")
	router.HandleFunc("/delete/{urlId}", urlHandler.DeleteURL).Methods("DELETE")

	log.Println("Starting server on :" + cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
