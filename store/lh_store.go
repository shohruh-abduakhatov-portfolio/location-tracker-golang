package store

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type locationHistoryStore struct {
	conn *mongo.Database
}

// GlobalLocationHistoryStore keeps DB connection to LocalHistoryStore
var GlobalLocationHistoryStore LocationHistoryStore

// InitGlobalLocationHistoryStore initiates DB connection to LocalHistoryStore
func InitGlobalLocationHistoryStore() {
	GlobalLocationHistoryStore = DB.lhStore
}

func (lhs *locationHistoryStore) Insert(locHist *LocationHistory, driverID string) (interface{}, error) {
	collection := lhs.conn.Collection(driverID)
	insertResult, err := collection.InsertOne(context.TODO(), locHist)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return insertResult.InsertedID, nil
}

func (lh *locationHistoryStore) Update(locHist *LocationHistory, driverID string) error {
	return nil
}
