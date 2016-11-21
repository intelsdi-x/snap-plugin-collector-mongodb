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
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

func createMetric(ns plugin.Namespace) plugin.Metric {
	metricType := plugin.Metric{
		Namespace: ns,
		Version:   Version,
	}
	return metricType
}

func merge(maps ...map[string]string) (output map[string]string) {
	size := len(maps)
	if size == 0 {
		return output
	}
	if size == 1 {
		return maps[0]
	}
	output = make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			output[k] = v
		}
	}
	return output
}

func createMeasurement(mt plugin.Metric, value interface{}, ns plugin.Namespace, meta map[string]string) plugin.Metric {

	return plugin.Metric{
		Timestamp: time.Now(),
		Namespace: ns,
		Data:      value,
		Tags:      merge(mt.Tags, meta),
		Config:    mt.Config,
		Version:   Version,
	}
}

func filterNamespace(metricType string, mts []plugin.Metric) (int, []plugin.Metric) {
	filteredMetrics := []plugin.Metric{}
	for _, m := range mts {
		if m.Namespace.Strings()[nsMetricPostion] == metricType {
			filteredMetrics = append(filteredMetrics, m)
		}
	}
	return len(filteredMetrics), filteredMetrics
}

func getDBConfig(metric plugin.Metric) (map[string]string, error) {
	config := make(map[string]string)
	values := []string{"uri", "username", "password"}
	var err error
	for _, v := range values {
		config[v], err = getStringFromConfig(metric, v)
		if err != nil {
			return config, err
		}
	}
	return config, nil
}

func getStringFromConfig(metric plugin.Metric, value string) (string, error) {
	conf, err := metric.Config.GetString(value)
	if err != nil {
		return "", err
	}
	return conf, nil
}
