package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/hanaboso/go-metrics/pkg/udp"
)

const schemeInfluxDB = "influxdb://"

var i int
var i8 int8
var i6 int16
var i2 int32
var i4 int64
var f4 float64
var f2 float32
var b bool
var s string

type influxDbMetrics struct {
	connection *udp.Connection
}

// Send metrics to InfluxDB
func (metrics influxDbMetrics) Send(name string, tags map[string]interface{}, fields map[string]interface{}) error {
	message, err := metrics.createMessage(name, tags, fields)

	if err != nil {
		return err
	}

	if _, err = metrics.connection.UDP.Write([]byte(message)); err != nil {
		return err
	}

	return nil
}

// Disconnect from InfluxDB
func (metrics influxDbMetrics) Disconnect() {
	metrics.connection.Disconnect()
}

// IsConnected checks connection status
func (metrics influxDbMetrics) IsConnected() bool {
	return metrics.connection.IsConnected()
}

func (metrics influxDbMetrics) createMessage(name string, tags map[string]interface{}, fields map[string]interface{}) (string, error) {
	if len(fields) == 0 {
		return "", errors.New("fields must not be empty")
	}

	innerTags := metrics.join(metrics.prepareItems(tags, false))
	innerFlags := metrics.join(metrics.prepareItems(fields, true))

	return fmt.Sprintf("%s,%s %s %d", name, innerTags, innerFlags, time.Now().UnixNano()), nil
}

func (metrics influxDbMetrics) join(items map[string]interface{}) string {
	res := ""

	if len(items) == 0 {
		return res
	}

	for k, item := range items {
		res += fmt.Sprintf("%s=%s,", k, item)
	}

	return strings.TrimSuffix(res, ",")
}

func (metrics influxDbMetrics) prepareItems(items map[string]interface{}, escape bool) map[string]interface{} {
	for k, item := range items {
		t := reflect.TypeOf(item)

		if item == "" {
			items[k] = "\"\""
		} else if t == reflect.TypeOf(i) || t == reflect.TypeOf(i8) || t == reflect.TypeOf(i6) || t == reflect.TypeOf(i2) || t == reflect.TypeOf(i4) {
			items[k] = fmt.Sprintf("%d", item)
		} else if t == reflect.TypeOf(f2) || t == reflect.TypeOf(f4) {
			items[k] = fmt.Sprintf("%f", item)
		} else if t == reflect.TypeOf(b) {
			switch item {
			case true:
				items[k] = "true"
				break
			case false:
				items[k] = "false"
				break
			}
		} else if item == nil {
			items[k] = "null"
		} else if t == reflect.TypeOf(s) {
			if escape == true {
				items[k] = metrics.escapeString(fmt.Sprintf("%s", item))
			} else {
				items[k] = fmt.Sprintf("%s", item)
			}
		} else {
			delete(items, k)
		}
	}

	return items
}

func (metrics influxDbMetrics) escapeString(s string) string {
	return fmt.Sprintf("\"%s\"", strings.Replace(s, "\"", "\\\"", -1))
}

func createInfluxDbMetrics(dsn string) Interface {
	connection := &udp.Connection{}
	connection.Connect(strings.Replace(dsn, schemeInfluxDB, "", -1))

	return influxDbMetrics{connection}
}
