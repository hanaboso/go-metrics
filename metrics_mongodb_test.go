package metrics

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoDbMetrics(t *testing.T) {
	metrics = createMongoDbMetrics(getMongoDbDsn())
	assert.True(t, metrics.IsConnected())

	metrics.Disconnect()
	assert.False(t, metrics.IsConnected())
}

func TestMongoDbSend(t *testing.T) {
	connection.Connect(getMongoDbDsn())
	_ = connection.Database.Drop(context.Background())

	innerContext, cancel := connection.Context()
	defer cancel()

	assert.Nil(t, createMongoDbMetrics(getMongoDbDsn()).Send("metrics", data, map[string]interface{}{
		"nil":     nil,
		"true":    true,
		"false":   false,
		"float":   1.0,
		"empty":   "",
		"array":   []string{"string"},
		"string":  "string",
		"integer": 1,
	}))

	count, err := connection.Database.Collection("metrics").CountDocuments(innerContext, primitive.D{{}})
	assert.Nil(t, err)
	assert.Equal(t, 1, int(count))
}
