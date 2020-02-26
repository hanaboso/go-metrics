# Hanaboso GO Metrics

**Download**
```
go mod download github.com/hanaboso/go-metrics
```

**Usage MongoDB**
```
import "github.com/hanaboso/go-metrics"

metrics := metrics.Connect("mongodb://mongodb/database?connectTimeoutMS=2500&heartbeatFrequencyMS=2500")
metrics.Send("metrics", map[string]interface{}{"tag": "Tag"}, map[string]interface{}{"field": "Field"})
metrics.Disconnect()
```

**Usage InfluxDB**
```
import "github.com/hanaboso/go-metrics"

metrics := metrics.Connect("influxdb://influxdb:8089")
metrics.Send("metrics", map[string]interface{}{"tag": "Tag"}, map[string]interface{}{"field": "Field"})
metrics.Disconnect()
```
