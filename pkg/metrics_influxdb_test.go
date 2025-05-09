package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfluxDbMetrics(t *testing.T) {
	metrics = createInfluxDbMetrics(getInfluxDbDsn())
	assert.True(t, metrics.IsConnected())

	metrics.Disconnect()
	assert.False(t, metrics.IsConnected())
}

func TestInfluxDbSend(t *testing.T) {
	metrics = createInfluxDbMetrics(getInfluxDbDsn())

	assert.Nil(t, metrics.Send("metrics", data, map[string]interface{}{
		"nil":     nil,
		"true":    true,
		"false":   false,
		"float":   1.0,
		"empty":   "",
		"array":   []string{"string"},
		"string":  "string",
		"integer": 1,
	}))
	assert.Nil(t, metrics.Send("metrics", map[string]interface{}{}, data))
	assert.NotNil(t, metrics.Send("metrics", data, map[string]interface{}{}))
}
