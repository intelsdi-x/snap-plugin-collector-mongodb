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
	mwrapper "github.com/intelsdi-x/snap-plugin-collector-mongodb/mongo"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

var nsTypes = struct {
	memory      []string
	connections []string
	extrainfo   []string
	opscounters []string
	tmalloc     []string
	network     []string
	metrics     []string
}{
	memory:      []string{"bits", "resident", "virtual", "supported", "mapped", "mappedwithjournal"},
	connections: []string{"current", "available", "totalcreated"},
	extrainfo:   []string{"heap_usage_bytes", "page_faults"},
	opscounters: []string{"insert", "query", "update", "delete", "getmore", "command"},
	tmalloc: []string{"current_allocated_bytes", "heap_size", "pageheap_free_bytes", "pageheap_unmapped_bytes", "max_total_thread_cache_bytes", "current_total_thread_cache_bytes",
		"total_free_bytes", "central_cache_free_bytes", "transfer_cache_free_bytes", "thread_cache_free_bytes", "aggressive_memory_decommit"},
	network: []string{"bytesin", "bytesout", "numrequests"},
	metrics: []string{"document_deleted", "document_inserted", "document_returned", "document_updated",
		"operation_fastmod", "operation_idhack", "operation_scanandorder", "operation_write_conflicts",
		"queryexecutor_scanned", "queryexecutor_scannedobjects",
		"record_moves",
		"ttl_deleteddocuments", "ttl_passes",
		"storage_freelist_search_bucketexhausted", "storage_freelist_search_requests", "storage_freelist_search_scanned",
		"repl_executor_counters_eventcreated", "repl_executor_counters_eventwait", "repl_executor_counters_cancels", "repl_executor_counters_waits", "repl_executor_counters_schedulednetcmd", "repl_executor_counters_scheduleddbwork", "repl_executor_counters_scheduledxclwork", "repl_executor_counters_scheduledworkat", "repl_executor_counters_scheduledwork", "repl_executor_counters_schedulingfailures",
		"repl_executor_queues_networkinprogress", "repl_executor_queues_dbworkinprogress", "repl_executor_queues_exclusiveinprogress", "repl_executor_queues_sleepers", "repl_executor_queues_ready", "repl_executor_queues_free",
		"repl_executor_unsignaledevents", "repl_executor_eventwaiters",
		"repl_apply_batchesnum", "repl_apply_batches_totalmillis", "repl_apply_ops",
		"repl_buffer_count", "repl_buffer_maxsizebytes", "repl_buffer_sizebytes",
		"repl_network_bytes", "repl_network_getmores_num", "repl_network_getmores_totalmillis", "repl_network_ops", "repl_network_readerscreated",
		"repl_preload_docs_num", "repl_preload_docs_totalmillis", "repl_preload_indexesnum", "repl_preload_indexes_totalmillis",
		"repl_pcursor_timedout", "repl_pcursor_open_notimeout", "repl_pcursor_open_pinned", "repl_pcursor_open_total",
	},
}

func getMemoryMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {

		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "bits":
			metrics = append(metrics, createMeasurement(mt, status.Mem.Bits, ns, meta))
		case "resident":
			metrics = append(metrics, createMeasurement(mt, status.Mem.Resident, ns, meta))
		case "virtual":
			metrics = append(metrics, createMeasurement(mt, status.Mem.Virtual, ns, meta))
		case "supported":
			metrics = append(metrics, createMeasurement(mt, status.Mem.Supported, ns, meta))
		case "mapped":
			metrics = append(metrics, createMeasurement(mt, status.Mem.Mapped, ns, meta))
		case "mappedwithjournal":
			metrics = append(metrics, createMeasurement(mt, status.Mem.MappedWithJournal, ns, meta))
		}
	}
	return metrics
}

func getConnectionMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "current":
			metrics = append(metrics, createMeasurement(mt, status.Connections.Current, ns, meta))
		case "available":
			metrics = append(metrics, createMeasurement(mt, status.Connections.Available, ns, meta))
		case "totalcreated":
			metrics = append(metrics, createMeasurement(mt, status.Connections.TotalCreated, ns, meta))

		}
	}
	return metrics
}

func getOpsCounterMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "insert":
			metrics = append(metrics, createMeasurement(mt, status.Opcounters.Insert, ns, meta))
		case "query":
			metrics = append(metrics, createMeasurement(mt, status.Opcounters.Query, ns, meta))
		case "update":
			metrics = append(metrics, createMeasurement(mt, status.Opcounters.Update, ns, meta))
		case "delete":
			metrics = append(metrics, createMeasurement(mt, status.Opcounters.Delete, ns, meta))
		case "getmore":
			metrics = append(metrics, createMeasurement(mt, status.Opcounters.Getmore, ns, meta))
		case "command":
			metrics = append(metrics, createMeasurement(mt, status.Opcounters.Command, ns, meta))
		}
	}
	return metrics
}

func getExtraInfoMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "heap_usage_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Extra_Info.HeapUsageBytes, ns, meta))
		case "page_faults":
			metrics = append(metrics, createMeasurement(mt, status.Extra_Info.PageFaults, ns, meta))
		}
	}
	return metrics
}

func getTmallocMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "current_allocated_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Generic.CurrentAllocatedBytes, ns, meta))
		case "heap_size":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Generic.HeapSize, ns, meta))
		case "pageheap_free_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.PageheapFreeBytes, ns, meta))
		case "pageheap_unmapped_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.PageheapUnmappedBytes, ns, meta))
		case "max_total_thread_cache_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.MaxTotalThreadCacheBytes, ns, meta))
		case "current_total_thread_cache_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.CurrentTotalThreadCacheBytes, ns, meta))
		case "total_free_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.TotalFreeBytes, ns, meta))
		case "central_cache_free_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.CentralCacheFreeBytes, ns, meta))
		case "transfer_cache_free_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.TransferCacheFreeBytes, ns, meta))
		case "thread_cache_free_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.ThreadCacheFreeBytes, ns, meta))
		case "aggressive_memory_decommit":
			metrics = append(metrics, createMeasurement(mt, status.Tcmalloc.Tcmalloc.AggressiveMemoryDecommit, ns, meta))

		}
	}
	return metrics
}

func getNetworkMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "bytesin":
			metrics = append(metrics, createMeasurement(mt, status.Network.BytesIn, ns, meta))
		case "bytesout":
			metrics = append(metrics, createMeasurement(mt, status.Network.BytesOut, ns, meta))
		case "numrequests":
			metrics = append(metrics, createMeasurement(mt, status.Network.NumRequests, ns, meta))
		}
	}
	return metrics
}

func getMetricMetrics(status mwrapper.ServerStatus, mts []plugin.Metric, meta map[string]string) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := make([]plugin.NamespaceElement, len(mt.Namespace))
		copy(ns, mt.Namespace)

		switch ns[nsSubMetric].Value {
		case "document_deleted":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Document.Deleted, ns, meta))
		case "document_inserted":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Document.Inserted, ns, meta))
		case "document_returned":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Document.Returned, ns, meta))
		case "document_updated":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Document.Updated, ns, meta))
		case "operation_fastmod":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Operation.Fastmod, ns, meta))
		case "operation_idhack":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Operation.Idhack, ns, meta))
		case "operation_scanandorder":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Operation.ScanAndOrder, ns, meta))
		case "operation_write_conflicts":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Operation.WriteConflicts, ns, meta))
		case "record_moves":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Record.Moves, ns, meta))
		case "ttl_deleteddocuments":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.TTL.DeletedDocuments, ns, meta))
		case "ttl_passes":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.TTL.Passes, ns, meta))
		case "storage_freelist_search_bucketexhausted":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Storage.Freelist.Search.BucketExhausted, ns, meta))
		case "storage_freelist_search_requests":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Storage.Freelist.Search.Requests, ns, meta))
		case "storage_freelist_search_scanned":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Storage.Freelist.Search.Scanned, ns, meta))
		case "repl_executor_counters_eventcreated":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.EventCreated, ns, meta))
		case "repl_executor_counters_eventwait":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.EventWait, ns, meta))
		case "repl_executor_counters_cancels":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.Cancels, ns, meta))
		case "repl_executor_counters_waits":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.Waits, ns, meta))
		case "repl_executor_counters_schedulednetcmd":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.ScheduledNetCmd, ns, meta))
		case "repl_executor_counters_scheduleddbwork":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.ScheduledDBWork, ns, meta))
		case "repl_executor_counters_scheduledxclwork":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.ScheduledXclWork, ns, meta))
		case "repl_executor_counters_scheduledworkat":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.ScheduledWorkAt, ns, meta))
		case "repl_executor_counters_scheduledwork":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.ScheduledWork, ns, meta))
		case "repl_executor_counters_schedulingfailures":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Counters.SchedulingFailures, ns, meta))
		case "repl_executor_queues_networkinprogress":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Queues.NetworkInProgress, ns, meta))
		case "repl_executor_queues_dbworkinprogress":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Queues.DbWorkInProgress, ns, meta))
		case "repl_executor_queues_exclusiveinprogress":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Queues.ExclusiveInProgress, ns, meta))
		case "repl_executor_queues_sleepers":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Queues.Sleepers, ns, meta))
		case "repl_executor_queues_ready":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Queues.Ready, ns, meta))
		case "repl_executor_queues_free":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.Queues.Free, ns, meta))
		case "repl_executor_unsignaledevents":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.UnsignaledEvents, ns, meta))
		case "repl_executor_eventwaiters":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Executor.EventWaiters, ns, meta))
		case "repl_apply_batchesnum":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Apply.Batches.Num, ns, meta))
		case "repl_apply_batches_totalmillis":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Apply.Batches.TotalMillis, ns, meta))
		case "repl_apply_ops":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Apply.Ops, ns, meta))
		case "repl_buffer_count":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Buffer.Count, ns, meta))
		case "repl_buffer_maxsizebytes":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Buffer.MaxSizeBytes, ns, meta))
		case "repl_buffer_sizebytes":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Buffer.SizeBytes, ns, meta))
		case "repl_network_bytes":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Network.Bytes, ns, meta))
		case "repl_network_getmores_num":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Network.Getmores.Num, ns, meta))
		case "repl_network_getmores_totalmillis":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Network.Getmores.TotalMillis, ns, meta))
		case "repl_network_ops":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Network.Ops, ns, meta))
		case "repl_network_readerscreated":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Network.ReadersCreated, ns, meta))
		case "repl_preload_docs_num":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Preload.Docs.Num, ns, meta))
		case "repl_preload_docs_totalmillis":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Preload.Docs.TotalMillis, ns, meta))
		case "repl_preload_indexesnum":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Preload.Indexes.Num, ns, meta))
		case "repl_preload_indexes_totalmillis":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Repl.Preload.Indexes.TotalMillis, ns, meta))
		case "repl_cursor_timedout":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Cursor.TimedOut, ns, meta))
		case "repl_cursor_open_notimeout":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Cursor.Open.NoTimeout, ns, meta))
		case "repl_cursor_open_pinned":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Cursor.Open.Pinned, ns, meta))
		case "repl_cursor_open_total":
			metrics = append(metrics, createMeasurement(mt, status.Metrics.Cursor.Open.Total, ns, meta))

		}
	}
	return metrics
}
