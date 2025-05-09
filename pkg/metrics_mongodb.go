package pkg

import (
	"context"
	"github.com/hanaboso/go-mongodb"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/connstring"
	"time"
)

const schemeMongoDB = connstring.SchemeMongoDB + "://"
const schemeMongoDBSRV = connstring.SchemeMongoDBSRV + "://"

type mongoDbMetrics struct {
	connection *mongodb.Connection
}

// Send metrics to MongoDB
func (metrics mongoDbMetrics) Send(name string, tags map[string]interface{}, fields map[string]interface{}) error {
	innerContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := metrics.connection.Database.Collection(name).InsertOne(innerContext, map[string]interface{}{
		"tags":   tags,
		"fields": fields,
	})

	return err
}

// Disconnect from MongoDB
func (metrics mongoDbMetrics) Disconnect() {
	metrics.connection.Disconnect()
}

// IsConnected checks connection status
func (metrics mongoDbMetrics) IsConnected() bool {
	return metrics.connection.IsConnected()
}

func createMongoDbMetrics(dsn string) Interface {
	connection := &mongodb.Connection{}
	connection.Connect(dsn)

	return mongoDbMetrics{connection}
}
