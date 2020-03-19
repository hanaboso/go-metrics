package pkg

import (
	"os"
	"testing"

	"github.com/hanaboso/go-mongodb"
	"github.com/stretchr/testify/assert"
)

var metrics Interface
var connection = mongodb.Connection{}
var data = map[string]interface{}{"key": "value"}

func TestConnect(t *testing.T) {
	assert.Implements(t, (*Interface)(nil), Connect(getMongoDbDsn()))
	assert.Implements(t, (*Interface)(nil), Connect(getInfluDbDsn()))

	assert.Panics(t, func() {
		Connect("Unknown")
	})
}

func getMongoDbDsn() string {
	if dsn := os.Getenv("METRICS_MONGO_DSN"); dsn != "" {
		return dsn
	}

	return "mongodb://127.0.0.26/database?connectTimeoutMS=2500&serverSelectionTimeoutMS=2500&socketTimeoutMS=2500&heartbeatFrequencyMS=2500"
}

func getInfluDbDsn() string {
	if dsn := os.Getenv("METRICS_INFLUX_DSN"); dsn != "" {
		return dsn
	}

	return "influxdb://127.0.0.26:8089"
}
