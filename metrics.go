package metrics

import (
	"fmt"
	"strings"
)

// Interface represents abstract metrics implementation
type Interface interface {
	Send(name string, tags map[string]interface{}, fields map[string]interface{}) error
	Disconnect()
	IsConnected() bool
}

// Connect creates specific metrics implementation
func Connect(dsn string) Interface {
	if strings.HasPrefix(dsn, schemeMongoDB) || strings.HasPrefix(dsn, schemeMongoDBSRV) {
		return createMongoDbMetrics(dsn)
	}

	if strings.HasPrefix(dsn, schemeInfluxDB) {
		return createInfluxDbMetrics(dsn)
	}

	panic(fmt.Sprintf("[Metrics] Supported DSNs are %s, %s and %s!", schemeMongoDB, schemeMongoDBSRV, schemeInfluxDB))
}
