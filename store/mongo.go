package store

import (
	"context"
	"fmt"
	"log"

	config "gitlab.com/logitab/back-end-team/location-tracker-go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStore stores connection instance per database
type MongoStore struct {
	Conn    *mongo.Database
	lhStore *locationHistoryStore
}

var (
	// DB stores database connection
	DB *MongoStore
)

// Init connects to mongodb
func init() {
	db, err := Connect()
	if err != nil || db == nil {
		panic("Cannot init db")
	}
	DB = db
	InitGlobalLocationHistoryStore()
}

// Connect connects to a store.
func Connect() (*MongoStore, error) {
	cfg := config.Cfg.Mongo

	// Set client options
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("%s://%s:%s", cfg.Protocol, cfg.IP, cfg.Port))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get Db by isntance
	db := client.Database(cfg.DBName)
	s := &MongoStore{
		Conn:    db,
		lhStore: &locationHistoryStore{db},
	}

	return s, nil
}
