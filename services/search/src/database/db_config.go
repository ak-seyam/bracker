package database

import (
	"os"

	"github.com/A-Siam/bracker/search/src/common/loggers"

	db "github.com/elastic/go-elasticsearch/v7"
)

var config = db.Config{
	Addresses: []string{os.Getenv("DB_BASE_URL")},
	Username:  os.Getenv("DB_USERNAME"),
	Password:  os.Getenv("DB_PASSWORD"),
}
var client, err = db.NewClient(config)

func GetDB() *db.Client {
	loggers.InfoLogger.Println("Initializing ElasticSearch Client")
	if err != nil {
		panic(err)
	}
	return client
}
