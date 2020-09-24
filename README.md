# Hanaboso GO Metrics

**Usage MongoDB**
```
import metrics "github.com/hanaboso/go-metrics/pkg"

metrics := metrics.Connect("mongodb://mongodb/database?connectTimeoutMS=2500&heartbeatFrequencyMS=2500")
metrics.Send("metrics", map[string]interface{}{"tag": "Tag"}, map[string]interface{}{"field": "Field"})
metrics.Disconnect()
```

**Usage InfluxDB**
```
import metrics "github.com/hanaboso/go-metrics/pkg"

metrics := metrics.Connect("influxdb://influxdb:8089")
metrics.Send("metrics", map[string]interface{}{"tag": "Tag"}, map[string]interface{}{"field": "Field"})
metrics.Disconnect()
```
