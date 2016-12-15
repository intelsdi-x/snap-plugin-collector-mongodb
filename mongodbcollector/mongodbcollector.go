/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mongodbcollector

import (
	"reflect"

	mwrapper "github.com/intelsdi-x/snap-plugin-collector-mongodb/mongo"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	// Name of plugin
	Name = "mongodb"
	// Vendor  prefix
	Vendor = "intel"
	// Plugin plugin name
	Plugin = "mongodb"
	// Version of plugin
	Version = 2

	nsMetricPostion = 2
	nsSubMetric     = 3
)

// MongoDBCollector type
type MongoDBCollector struct {
}

// CollectMetrics returns collected metrics
func (MongoDBCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	meta := make(map[string]string)

	_, memoryMetrics := filterNamespace("memory", mts)
	_, connectionsMetrics := filterNamespace("connections", mts)
	_, opscountersMetrics := filterNamespace("opscounters", mts)
	_, extrainfoMetrics := filterNamespace("extrainfo", mts)
	_, tmallocMetrics := filterNamespace("tmalloc", mts)
	_, networkMetrics := filterNamespace("network", mts)
	_, metricMetrics := filterNamespace("metric", mts)

	dbConfig, err := getDBConfig(mts[0])
	if err != nil {
		return nil, err
	}
	conn, err := mwrapper.GetConnection(dbConfig["uri"], dbConfig["passsword"], dbConfig["username"])

	if err != nil {
		return nil, err
	}
	serverStatus, err := mwrapper.GetServerStatus(conn)
	if err != nil {
		return nil, err
	}

	metrics = append(metrics, getMemoryMetrics(serverStatus, memoryMetrics, meta)...)
	metrics = append(metrics, getConnectionMetrics(serverStatus, connectionsMetrics, meta)...)
	metrics = append(metrics, getOpsCounterMetrics(serverStatus, opscountersMetrics, meta)...)
	metrics = append(metrics, getExtraInfoMetrics(serverStatus, extrainfoMetrics, meta)...)
	metrics = append(metrics, getTmallocMetrics(serverStatus, tmallocMetrics, meta)...)
	metrics = append(metrics, getNetworkMetrics(serverStatus, networkMetrics, meta)...)
	metrics = append(metrics, getMetricMetrics(serverStatus, metricMetrics, meta)...)

	return metrics, nil

}

// GetConfigPolicy returns a config policy
func (MongoDBCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	configKey := []string{"intel", "mongodb"}
	policy.AddNewStringRule(configKey,
		"uri",
		false)
	policy.AddNewStringRule(configKey,
		"username",
		false)
	policy.AddNewStringRule(configKey,
		"password",
		false)
	return *policy, nil
}

// GetMetricTypes returns metric types that can be collected
func (MongoDBCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {

	var metrics []plugin.Metric
	ns := plugin.NewNamespace(Vendor, Plugin)
	metricTypes := reflect.ValueOf(&nsTypes).Elem()
	typeOfmetrics := metricTypes.Type()
	for i := 0; i < metricTypes.NumField(); i++ {
		for j := 0; j < metricTypes.Field(i).Len(); j++ {
			metrics = append(metrics, createMetric(ns.AddStaticElements(
				typeOfmetrics.Field(i).Name, metricTypes.Field(i).Index(j).String())))
		}
	}

	return metrics, nil
}
