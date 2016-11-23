//
// +build medium

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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMongoDBPlugin(t *testing.T) {
	config := plugin.Config{
		"uri":      "localhost:27017",
		"username": "",
		"password": "",
	}
	Convey("Create MognoDB Collector", t, func() {
		mongodbCol := MongoDBCollector{}
		Convey("So MognoDB should not be nil", func() {
			So(mongodbCol, ShouldNotBeNil)
		})
		Convey("So MognoDB should be of MongoDB type", func() {
			So(mongodbCol, ShouldHaveSameTypeAs, MongoDBCollector{})
		})
		Convey("mongodbCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := mongodbCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})
	Convey("Get Metric MognoDB Types", t, func() {
		mongodbCol := MongoDBCollector{}
		var cfg = plugin.Config{}
		metrics, err := mongodbCol.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 84)
	})

	Convey("Collect Memory Metrics", t, func() {
		mongodbCol := MongoDBCollector{}

		mts := []plugin.Metric{}
		for _, v := range nsTypes.memory {

			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElements("memory", v), Config: config})
		}
		metrics, err := mongodbCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 6)
		So(metrics[0].Data, ShouldNotBeNil)

	})

	Convey("Collect Ops Metrics", t, func() {
		mongodbCol := MongoDBCollector{}

		mts := []plugin.Metric{}
		for _, v := range nsTypes.opscounters {

			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElements("opscounters", v), Config: config})
		}
		metrics, err := mongodbCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 6)
		So(metrics[0].Data, ShouldNotBeNil)

	})

	Convey("Collect Network Metrics", t, func() {
		mongodbCol := MongoDBCollector{}

		mts := []plugin.Metric{}
		for _, v := range nsTypes.network {

			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElements("network", v), Config: config})
		}
		metrics, err := mongodbCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 3)
		So(metrics[0].Data, ShouldNotBeNil)

	})
	Convey("Collect Tmalloc Metrics", t, func() {
		mongodbCol := MongoDBCollector{}

		mts := []plugin.Metric{}
		for _, v := range nsTypes.tmalloc {
			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElements("tmalloc", v), Config: config})
		}
		metrics, err := mongodbCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 11)
		So(metrics[0].Data, ShouldNotBeNil)

	})
	Convey("Collect Metric Metrics", t, func() {
		mongodbCol := MongoDBCollector{}

		mts := []plugin.Metric{}
		for _, v := range nsTypes.metrics {
			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElements("metric", v), Config: config})
		}
		metrics, err := mongodbCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 47)
		So(metrics[0].Data, ShouldNotBeNil)

	})
	Convey("Collect Extrainfo Metrics", t, func() {
		mongodbCol := MongoDBCollector{}

		mts := []plugin.Metric{}
		for _, v := range nsTypes.extrainfo {
			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin).AddStaticElements("extrainfo", v), Config: config})
		}
		metrics, err := mongodbCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 2)
		So(metrics[0].Data, ShouldNotBeNil)

	})
}
