package udp

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var connection = Connection{}

func TestUdp(t *testing.T) {
	connection.Connect(strings.Replace(os.Getenv("METRICS_INFLUX_DSN"), "influxdb://", "", -1))
	assert.True(t, connection.IsConnected())

	connection.Disconnect()
	assert.False(t, connection.IsConnected())
}
