# snap collector plugin - mongodb
<!-- last updated for mongoDB 3.5.7 -->
## Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Type | Description
----------|------|----------------
/intel/mongodb/memory/bits | int | basic size of allocation (bit)
/intel/mongodb/memory/resident | int | amount of memory used by the application (MB)
/intel/mongodb/memory/virtual | int | virtual memory usage (MB)
/intel/mongodb/memory/supported | int | flag if memory info is supported
/intel/mongodb/memory/mapped | int | amount of virtual memory used to map the database into memory (MB)
/intel/mongodb/memory/mappedwithjournal | int | amount of virtual memory for database including mapped files (MB)
/intel/mongodb/connections/current | int | number of sessions currently open in the transport layer
/intel/mongodb/connections/available | int | number of available sessions that could still be opened
/intel/mongodb/connections/totalcreated | int | total number of sessions that have ever been created by the transport layer
/intel/mongodb/extrainfo/heap_usage_bytes | int | (experimental)
/intel/mongodb/extrainfo/page_faults | int | number of page faults experienced by application
/intel/mongodb/opscounters/insert | int | counts performed insert operations
/intel/mongodb/opscounters/query | int | counts performed query operations
/intel/mongodb/opscounters/update | int | counts performed update operations
/intel/mongodb/opscounters/delete | int | counts performed delete operations
/intel/mongodb/opscounters/getmore | int | counts performed getmore operations
/intel/mongodb/opscounters/command | int | counts performed command operations
/intel/mongodb/tmalloc/current_allocated_bytes | int | number of bytes used by the application
/intel/mongodb/tmalloc/heap_size | int | bytes of system memory reserved for allocation
/intel/mongodb/tmalloc/pageheap_free_bytes | int | number of bytes in free, mapped pages in page heap
/intel/mongodb/tmalloc/pageheap_unmapped_bytes | int | number of bytes in free, unmapped pages in page heap that have been released to OS
/intel/mongodb/tmalloc/max_total_thread_cache_bytes | int | upper limit on total number of bytes stored across all per-thread caches
/intel/mongodb/tmalloc/current_total_thread_cache_bytes | int | number of bytes used across all thread caches
/intel/mongodb/tmalloc/total_free_bytes | int | total number of free bytes across all caches
/intel/mongodb/tmalloc/central_cache_free_bytes | int | number of free bytes in the central cache
/intel/mongodb/tmalloc/transfer_cache_free_bytes | int | number of free bytes that are waiting to be transfered between the central cache and a thread cache
/intel/mongodb/tmalloc/thread_cache_free_bytes | int | number of free bytes in thread caches
/intel/mongodb/tmalloc/aggressive_memory_decommit | int | status of aggressive memory decommit mode
/intel/mongodb/network/bytesin | int | number of bytes in received requests
/intel/mongodb/network/bytesout | int | number of bytes in emitted responses
/intel/mongodb/network/numrequests | int | number of received requests
/intel/mongodb/metrics/document_deleted | int | total number of documents deleted
/intel/mongodb/metrics/document_inserted | int | total number of documents inserted
/intel/mongodb/metrics/document_returned | int | total number of documents returned
/intel/mongodb/metrics/document_updated | int | total number of documents updated
/intel/mongodb/metrics/operation_fastmod | int | number of update operations that neither cause documents to grow nor require updates to the index
/intel/mongodb/metrics/operation_idhack | int | number of queries that contain the _id field
/intel/mongodb/metrics/operation_scanandorder | int | total number of queries that return sorted numbers that cannot perform the sort operation using an index
/intel/mongodb/metrics/operation_write_conflicts | int | total number of queries that encounted write conflicts
/intel/mongodb/metrics/queryexecutor_scanned | int | total number of index items scanned during queries and query-plan evaluation
/intel/mongodb/metrics/queryexecutor_scannedobjects | int | total number of documents scanned during queries and query-plan evaluation
/intel/mongodb/metrics/record_moves | int | reports the total number of times documents move within the on-disk representation of the MongoDB data set
/intel/mongodb/metrics/ttl_deleteddocuments | int | total number of documents deleted from collections with a ttl index
/intel/mongodb/metrics/ttl_passes | int | number of times the background process removes documents from collections with a ttl index
/intel/mongodb/metrics/storage_freelist_search_bucketexhausted | int | number of times that mongod has checked the free list without finding a suitably large record allocation
/intel/mongodb/metrics/storage_freelist_search_requests | int | number of times mongod has searched for available record allocations
/intel/mongodb/metrics/storage_freelist_search_scanned | int | number of available record allocations mongod has searched
/intel/mongodb/metrics/repl_executor_counters_eventcreated | int | number of created replication events
/intel/mongodb/metrics/repl_executor_counters_eventwait | int | number of created replication events
/intel/mongodb/metrics/repl_executor_counters_cancels | int | number of canceled replication events
/intel/mongodb/metrics/repl_executor_counters_waits | int | number of replication waits
/intel/mongodb/metrics/repl_executor_counters_schedulednetcmd | int | number of scheduled remote operations
/intel/mongodb/metrics/repl_executor_counters_scheduleddbwork | int | number of scheduled db-work operations
/intel/mongodb/metrics/repl_executor_counters_scheduledxclwork | int | number of operations scheduled under exclusive lock
/intel/mongodb/metrics/repl_executor_counters_scheduledworkat | int | number of delayed scheduled operations
/intel/mongodb/metrics/repl_executor_counters_scheduledwork | int | number of additional scheduled operations
/intel/mongodb/metrics/repl_executor_counters_schedulingfailures | int | number of failed scheduling attempts
/intel/mongodb/metrics/repl_executor_queues_networkinprogress | int | number of remote scheduled commands blocked on the network
/intel/mongodb/metrics/repl_executor_queues_dbworkinprogress | int | number of scheduled db-work operations pending execution
/intel/mongodb/metrics/repl_executor_queues_exclusiveinprogress | int | number of scheduled pending operations requiring exclusive lock
/intel/mongodb/metrics/repl_executor_queues_sleepers | int | counts operations waiting for a timer
/intel/mongodb/metrics/repl_executor_queues_ready | int | counts queued operations ready for execution
/intel/mongodb/metrics/repl_executor_queues_free | int | (experminental)
/intel/mongodb/metrics/repl_executor_unsignaledevents | int | counts events that have yet to be signaled
/intel/mongodb/metrics/repl_executor_eventwaiters | int | counts threads awaiting for events to be signaled
/intel/mongodb/metrics/repl_apply_batchesnum | int | counts executed batches
/intel/mongodb/metrics/repl_apply_batches_totalmillis | int | counts batches total execution time
/intel/mongodb/metrics/repl_apply_ops | int | number of oplog entries applied
/intel/mongodb/metrics/repl_buffer_count | int | count of items in the buffer
/intel/mongodb/metrics/repl_buffer_maxsizebytes | int | max size of the buffer (bytes)
/intel/mongodb/metrics/repl_buffer_sizebytes | int | size of items in the buffer (bytes)
/intel/mongodb/metrics/repl_network_bytes | int | bytes read via the oplog reader
/intel/mongodb/metrics/repl_network_getmores_num | int | number of times reading batches off the network
/intel/mongodb/metrics/repl_network_getmores_totalmillis | int | time spent reading batches off the network
/intel/mongodb/metrics/repl_network_ops | int | oplog entries read via the oplog reader
/intel/mongodb/metrics/repl_network_readerscreated | int | number of readers created
/intel/mongodb/metrics/repl_preload_docs_num | int | count (of batches) spent fetching pages before application
/intel/mongodb/metrics/repl_preload_docs_totalmillis | int | time spent fetching pages before application
/intel/mongodb/metrics/repl_preload_indexesnum | int | count (of batches) spent fetching index before application
/intel/mongodb/metrics/repl_preload_indexes_totalmillis | int | time spent fetching index before application
/intel/mongodb/metrics/repl_pcursor_timedout | int | number of timed out cursors
/intel/mongodb/metrics/repl_pcursor_open_notimeout | int | number of cursors not timed out
/intel/mongodb/metrics/repl_pcursor_open_pinned | int | number of open pinned cursors
/intel/mongodb/metrics/repl_pcursor_open_total | int | total number of open cursors
