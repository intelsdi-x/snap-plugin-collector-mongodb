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

package mongo

import "time"

//Asserts contains info about Asserts
type Asserts struct {
	Regular   int `json:"regular" bson:"regular"`
	Warning   int `json:"warning" bson:"warning"`
	Msg       int `json:"msg" bson:"msg"`
	User      int `json:"user" bson:"user"`
	Rollovers int `json:"rollovers" bson:"rollovers"`
}

//Connections contains info about Connections
type Connections struct {
	Current      int `json:"current" bson:"current"`
	Available    int `json:"available" bson:"available"`
	TotalCreated int `json:"totalCreated" bson:"totalCreated"`
}

//Extra_Info contains info about ExtraInfo
type Extra_Info struct {
	Note           string `json:"note" bson:"note"`
	HeapUsageBytes int    `json:"heap_usage_bytes" bson:"heap_usage_bytes"`
	PageFaults     int    `json:"page_faults" bson:"page_faults"`
}

//GlobalLock contains info about GlobalLock
type GlobalLock struct {
	TotalTime     int            `json:"totalTime" bson:"totalTime"`
	CurrentQueue  *CurrentQueue  `json:"currentQueue" bson:"currentQueue"`
	ActiveClients *ActiveClients `json:"activeClients" bson:"activeClients"`
}

//ActiveClients contains info about ActiveClients
type ActiveClients struct {
	Total   int `json:"total" bson:"total"`
	Readers int `json:"readers" bson:"readers"`
	Writers int `json:"writers" bson:"writers"`
}

//CurrentQueue contains info about CurrentQueue
type CurrentQueue struct {
	Total   int `json:"total" bson:"total"`
	Readers int `json:"readers" bson:"readers"`
	Writers int `json:"writers" bson:"writers"`
}

//Opcounters contains info about Opcounters
type Opcounters struct {
	Insert  int `json:"insert" bson:"insert"`
	Query   int `json:"query" bson:"query"`
	Update  int `json:"update" bson:"update"`
	Delete  int `json:"delete" bson:"delete"`
	Getmore int `json:"getmore" bson:"getmore"`
	Command int `json:"command" bson:"command"`
}

// AcquireCount contains stats about AcquireCount
type AcquireCount struct {
	R int `json:"r" bson:"r"`
	W int `json:"w" bson:"w"`
}

// Global contains stats about Global
type Global struct {
	AcquireCount *AcquireCount `json:"acquireCount" bson:"acquireCount"`
}

// Database contains stats about Database
type Database struct {
	AcquireCount *AcquireCount `json:"acquireCount" bson:"acquireCount"`
}

// Collection contains stats about Collection
type Collection struct {
	AcquireCount *AcquireCount `json:"acquireCount" bson:"acquireCount"`
}

// Metadata contains stats about Metadata
type Metadata struct {
	AcquireCount *AcquireCount `json:"acquireCount" bson:"acquireCount"`
}

// Locks contains stats about Locks
type Locks struct {
	Global     *Global     `json:"Global" bson:"Global"`
	Database   *Database   `json:"Database" bson:"Database"`
	Collection *Collection `json:"Collection" bson:"Collection"`
	Metadata   *Metadata   `json:"Metadata" bson:"Metadata"`
}

// OpcountersRepl contains stats about OpcountersRepl
type OpcountersRepl struct {
	Insert  int `json:"insert" bson:"insert"`
	Query   int `json:"query" bson:"query"`
	Update  int `json:"update" bson:"update"`
	Delete  int `json:"delete" bson:"delete"`
	Getmore int `json:"getmore" bson:"getmore"`
	Command int `json:"command" bson:"command"`
}

// Network contains stats about Network
type Network struct {
	BytesIn     int `json:"bytesIn" bson:"bytesIn"`
	BytesOut    int `json:"bytesOut" bson:"bytesOut"`
	NumRequests int `json:"numRequests" bson:"numRequests"`
}

// StorageEngine contains stats about StorageEngine
type StorageEngine struct {
	Name                   string `json:"name" bson:"name"`
	SupportsCommittedReads bool   `json:"supportsCommittedReads" bson:"supportsCommittedReads"`
	Persistent             bool   `json:"persistent" bson:"persistent"`
}

// Generic contains stats about Generic
type Generic struct {
	CurrentAllocatedBytes int `json:"current_allocated_bytes" bson:"current_allocated_bytes"`
	HeapSize              int `json:"heap_size" bson:"heap_size"`
}

// TcmallocStat contains stats about TcmallocStat
type TcmallocStat struct {
	PageheapFreeBytes            int    `json:"pageheap_free_bytes" bson:"pageheap_free_bytes"`
	PageheapUnmappedBytes        int    `json:"pageheap_unmapped_bytes" bson:"pageheap_unmapped_bytes"`
	MaxTotalThreadCacheBytes     int    `json:"max_total_thread_cache_bytes" bson:"max_total_thread_cache_bytes"`
	CurrentTotalThreadCacheBytes int    `json:"current_total_thread_cache_bytes" bson:"current_total_thread_cache_bytes"`
	TotalFreeBytes               int    `json:"total_free_bytes" bson:"total_free_bytes"`
	CentralCacheFreeBytes        int    `json:"central_cache_free_bytes" bson:"central_cache_free_bytes"`
	TransferCacheFreeBytes       int    `json:"transfer_cache_free_bytes" bson:"transfer_cache_free_bytes"`
	ThreadCacheFreeBytes         int    `json:"thread_cache_free_bytes" bson:"thread_cache_free_bytes"`
	AggressiveMemoryDecommit     int    `json:"aggressive_memory_decommit" bson:"aggressive_memory_decommit"`
	FormattedString              string `json:"formattedString" bson:"formattedString"`
}

// Tcmalloc contains stats about Tcmalloc
type Tcmalloc struct {
	Generic  *Generic      `json:"generic" bson:"generic"`
	Tcmalloc *TcmallocStat `json:"tcmalloc" bson:"tcmalloc"`
}

// Mem contains stats about Mem
type Mem struct {
	Bits              int  `json:"bits" bson:"bits"`
	Resident          int  `json:"resident" bson:"resident"`
	Virtual           int  `json:"virtual" bson:"virtual"`
	Supported         bool `json:"supported" bson:"supported"`
	Mapped            int  `json:"mapped" bson:"mapped"`
	MappedWithJournal int  `json:"mappedWithJournal" bson:"mappedWithJournal"`
}

// LSM contains stats about LSM
type LSM struct {
	ApplicationWorkUnitsCurrentlyQueued int `json:"application work units currently queued" bson:"application work units currently queued"`
	MergeWorkUnitsCurrentlyQueued       int `json:"merge work units currently queued" bson:"merge work units currently queued"`
	RowsMergedInAnLSMTree               int `json:"rows merged in an LSM tree" bson:"rows merged in an LSM tree"`
	SleepForLSMCheckpointThrottle       int `json:"sleep for LSM checkpoint throttle" bson:"sleep for LSM checkpoint throttle"`
	SleepForLSMMergeThrottle            int `json:"sleep for LSM merge throttle" bson:"sleep for LSM merge throttle"`
	SwitchWorkUnitsCurrentlyQueued      int `json:"switch work units currently queued" bson:"switch work units currently queued"`
	TreeMaintenanceOperationsDiscarded  int `json:"tree maintenance operations discarded" bson:"tree maintenance operations discarded"`
	TreeMaintenanceOperationsExecuted   int `json:"tree maintenance operations executed" bson:"tree maintenance operations executed"`
	TreeMaintenanceOperationsScheduled  int `json:"tree maintenance operations scheduled" bson:"tree maintenance operations scheduled"`
	TreeQueueHitMaximum                 int `json:"tree queue hit maximum" bson:"tree queue hit maximum"`
}

// Async contains stats about Async
type Async struct {
	CurrentWorkQueueLength                    int `json:"current work queue length" bson:"current work queue length"`
	MaximumWorkQueueLength                    int `json:"maximum work queue length" bson:"maximum work queue length"`
	NumberOfAllocationStateRaces              int `json:"number of allocation state races" bson:"number of allocation state races"`
	NumberOfFlushCalls                        int `json:"number of flush calls" bson:"number of flush calls"`
	NumberOfOperationSlotsViewedForAllocation int `json:"number of operation slots viewed for allocation" bson:"number of operation slots viewed for allocation"`
	NumberOfTimesOperationAllocationFailed    int `json:"number of times operation allocation failed" bson:"number of times operation allocation failed"`
	NumberOfTimesWorkerFoundNoWork            int `json:"number of times worker found no work" bson:"number of times worker found no work"`
	TotalAllocations                          int `json:"total allocations" bson:"total allocations"`
	TotalCompactCalls                         int `json:"total compact calls" bson:"total compact calls"`
	TotalInsertCalls                          int `json:"total insert calls" bson:"total insert calls"`
	TotalRemoveCalls                          int `json:"total remove calls" bson:"total remove calls"`
	TotalSearchCalls                          int `json:"total search calls" bson:"total search calls"`
	TotalUpdateCalls                          int `json:"total update calls" bson:"total update calls"`
}

// BlockManager contains stats about BlockManager
type BlockManager struct {
	BlocksPreLoaded           int `json:"blocks pre-loaded" bson:"blocks pre-loaded"`
	BlocksRead                int `json:"blocks read" bson:"blocks read"`
	BlocksWritten             int `json:"blocks written" bson:"blocks written"`
	BytesRead                 int `json:"bytes read" bson:"bytes read"`
	BytesWritten              int `json:"bytes written" bson:"bytes written"`
	BytesWrittenForCheckpoint int `json:"bytes written for checkpoint" bson:"bytes written for checkpoint"`
	MappedBlocksRead          int `json:"mapped blocks read" bson:"mapped blocks read"`
	MappedBytesRead           int `json:"mapped bytes read" bson:"mapped bytes read"`
}

// Cache contains stats about Cache
type Cache struct {
	BytesBelongingToPageImagesInTheCache                       int `json:"bytes belonging to page images in the cache" bson:"bytes belonging to page images in the cache"`
	BytesCurrentlyInTheCache                                   int `json:"bytes currently in the cache" bson:"bytes currently in the cache"`
	BytesNotBelongingToPageImagesInTheCache                    int `json:"bytes not belonging to page images in the cache" bson:"bytes not belonging to page images in the cache"`
	BytesReadIntoCache                                         int `json:"bytes read into cache" bson:"bytes read into cache"`
	BytesWrittenFromCache                                      int `json:"bytes written from cache" bson:"bytes written from cache"`
	CheckpointBlockedPageEviction                              int `json:"checkpoint blocked page eviction" bson:"checkpoint blocked page eviction"`
	EvictionCallsToGetAPage                                    int `json:"eviction calls to get a page" bson:"eviction calls to get a page"`
	EvictionCallsToGetAPageFoundQueueEmpty                     int `json:"eviction calls to get a page found queue empty" bson:"eviction calls to get a page found queue empty"`
	EvictionCallsToGetAPageFoundQueueEmptyAfterLocking         int `json:"eviction calls to get a page found queue empty after locking" bson:"eviction calls to get a page found queue empty after locking"`
	EvictionCurrentlyOperatingInAggressiveMode                 int `json:"eviction currently operating in aggressive mode" bson:"eviction currently operating in aggressive mode"`
	EvictionEmptyScore                                         int `json:"eviction empty score" bson:"eviction empty score"`
	EvictionServerCandidateQueueEmptyWhenToppingUp             int `json:"eviction server candidate queue empty when topping up" bson:"eviction server candidate queue empty when topping up"`
	EvictionServerCandidateQueueNotEmptyWhenToppingUp          int `json:"eviction server candidate queue not empty when topping up" bson:"eviction server candidate queue not empty when topping up"`
	EvictionServerEvictingPages                                int `json:"eviction server evicting pages" bson:"eviction server evicting pages"`
	EvictionServerSleptBecauseWeDidNotMakeProgressWithEviction int `json:"eviction server slept because we did not make progress with eviction" bson:"eviction server slept because we did not make progress with eviction"`
	EvictionServerUnableToReachEvictionGoal                    int `json:"eviction server unable to reach eviction goal" bson:"eviction server unable to reach eviction goal"`
	EvictionState                                              int `json:"eviction state" bson:"eviction state"`
	EvictionWalksAbandoned                                     int `json:"eviction walks abandoned" bson:"eviction walks abandoned"`
	EvictionWorkerThreadEvictingPages                          int `json:"eviction worker thread evicting pages" bson:"eviction worker thread evicting pages"`
	FailedEvictionOfPagesThatExceededTheInMemoryMaximum        int `json:"failed eviction of pages that exceeded the in-memory maximum" bson:"failed eviction of pages that exceeded the in-memory maximum"`
	FilesWithActiveEvictionWalks                               int `json:"files with active eviction walks" bson:"files with active eviction walks"`
	FilesWithNewEvictionWalksStarted                           int `json:"files with new eviction walks started" bson:"files with new eviction walks started"`
	HazardPointerBlockedPageEviction                           int `json:"hazard pointer blocked page eviction" bson:"hazard pointer blocked page eviction"`
	HazardPointerCheckCalls                                    int `json:"hazard pointer check calls" bson:"hazard pointer check calls"`
	HazardPointerCheckEntriesWalked                            int `json:"hazard pointer check entries walked" bson:"hazard pointer check entries walked"`
	HazardPointerMaximumArrayLength                            int `json:"hazard pointer maximum array length" bson:"hazard pointer maximum array length"`
	InMemoryPagePassedCriteriaToBeSplit                        int `json:"in-memory page passed criteria to be split" bson:"in-memory page passed criteria to be split"`
	InMemoryPageSplits                                         int `json:"in-memory page splits" bson:"in-memory page splits"`
	InternalPagesEvicted                                       int `json:"internal pages evicted" bson:"internal pages evicted"`
	InternalPagesSplitDuringEviction                           int `json:"internal pages split during eviction" bson:"internal pages split during eviction"`
	LeafPagesSplitDuringEviction                               int `json:"leaf pages split during eviction" bson:"leaf pages split during eviction"`
	LookasideTableInsertCalls                                  int `json:"lookaside table insert calls" bson:"lookaside table insert calls"`
	LookasideTableRemoveCalls                                  int `json:"lookaside table remove calls" bson:"lookaside table remove calls"`
	MaximumBytesConfigured                                     int `json:"maximum bytes configured" bson:"maximum bytes configured"`
	MaximumPageSizeAtEviction                                  int `json:"maximum page size at eviction" bson:"maximum page size at eviction"`
	ModifiedPagesEvicted                                       int `json:"modified pages evicted" bson:"modified pages evicted"`
	ModifiedPagesEvictedByApplicationThreads                   int `json:"modified pages evicted by application threads" bson:"modified pages evicted by application threads"`
	OverflowPagesReadIntoCache                                 int `json:"overflow pages read into cache" bson:"overflow pages read into cache"`
	OverflowValuesCachedInMemory                               int `json:"overflow values cached in memory" bson:"overflow values cached in memory"`
	PageSplitDuringEvictionDeepenedTheTree                     int `json:"page split during eviction deepened the tree" bson:"page split during eviction deepened the tree"`
	PageWrittenRequiringLookasideRecords                       int `json:"page written requiring lookaside records" bson:"page written requiring lookaside records"`
	PagesCurrentlyHeldInTheCache                               int `json:"pages currently held in the cache" bson:"pages currently held in the cache"`
	PagesEvictedBecauseTheyExceededTheInMemoryMaximum          int `json:"pages evicted because they exceeded the in-memory maximum" bson:"pages evicted because they exceeded the in-memory maximum"`
	PagesEvictedBecauseTheyHadChainsOfDeletedItems             int `json:"pages evicted because they had chains of deleted items" bson:"pages evicted because they had chains of deleted items"`
	PagesEvictedByApplicationThreads                           int `json:"pages evicted by application threads" bson:"pages evicted by application threads"`
	PagesQueuedForEviction                                     int `json:"pages queued for eviction" bson:"pages queued for eviction"`
	PagesQueuedForUrgentEviction                               int `json:"pages queued for urgent eviction" bson:"pages queued for urgent eviction"`
	PagesQueuedForUrgentEvictionDuringWalk                     int `json:"pages queued for urgent eviction during walk" bson:"pages queued for urgent eviction during walk"`
	PagesReadIntoCache                                         int `json:"pages read into cache" bson:"pages read into cache"`
	PagesReadIntoCacheRequiringLookasideEntries                int `json:"pages read into cache requiring lookaside entries" bson:"pages read into cache requiring lookaside entries"`
	PagesRequestedFromTheCache                                 int `json:"pages requested from the cache" bson:"pages requested from the cache"`
	PagesSeenByEvictionWalk                                    int `json:"pages seen by eviction walk" bson:"pages seen by eviction walk"`
	PagesSelectedForEvictionUnableToBeEvicted                  int `json:"pages selected for eviction unable to be evicted" bson:"pages selected for eviction unable to be evicted"`
	PagesWalkedForEviction                                     int `json:"pages walked for eviction" bson:"pages walked for eviction"`
	PagesWrittenFromCache                                      int `json:"pages written from cache" bson:"pages written from cache"`
	PagesWrittenRequiringInMemoryRestoration                   int `json:"pages written requiring in-memory restoration" bson:"pages written requiring in-memory restoration"`
	PercentageOverhead                                         int `json:"percentage overhead" bson:"percentage overhead"`
	TrackedBytesBelongingToInternalPagesInTheCache             int `json:"tracked bytes belonging to internal pages in the cache" bson:"tracked bytes belonging to internal pages in the cache"`
	TrackedBytesBelongingToLeafPagesInTheCache                 int `json:"tracked bytes belonging to leaf pages in the cache" bson:"tracked bytes belonging to leaf pages in the cache"`
	TrackedDirtyBytesInTheCache                                int `json:"tracked dirty bytes in the cache" bson:"tracked dirty bytes in the cache"`
	TrackedDirtyPagesInTheCache                                int `json:"tracked dirty pages in the cache" bson:"tracked dirty pages in the cache"`
	UnmodifiedPagesEvicted                                     int `json:"unmodified pages evicted" bson:"unmodified pages evicted"`
}

// Connection contains stats about Connection
type Connection struct {
	AutoAdjustingConditionResets         int `json:"auto adjusting condition resets" bson:"auto adjusting condition resets"`
	AutoAdjustingConditionWaitCalls      int `json:"auto adjusting condition wait calls" bson:"auto adjusting condition wait calls"`
	FilesCurrentlyOpen                   int `json:"files currently open" bson:"files currently open"`
	MemoryAllocations                    int `json:"memory allocations" bson:"memory allocations"`
	MemoryFrees                          int `json:"memory frees" bson:"memory frees"`
	MemoryReAllocations                  int `json:"memory re-allocations" bson:"memory re-allocations"`
	PthreadMutexConditionWaitCalls       int `json:"pthread mutex condition wait calls" bson:"pthread mutex condition wait calls"`
	PthreadMutexSharedLockReadLockCalls  int `json:"pthread mutex shared lock read-lock calls" bson:"pthread mutex shared lock read-lock calls"`
	PthreadMutexSharedLockWriteLockCalls int `json:"pthread mutex shared lock write-lock calls" bson:"pthread mutex shared lock write-lock calls"`
	TotalFsyncIOs                        int `json:"total fsync I/Os" bson:"total fsync I/Os"`
	TotalReadIOs                         int `json:"total read I/Os" bson:"total read I/Os"`
	TotalWriteIOs                        int `json:"total write I/Os" bson:"total write I/Os"`
}

// WireCursor contains stats about WireCursor
type WireCursor struct {
	CursorCreateCalls       int `json:"cursor create calls" bson:"cursor create calls"`
	CursorInsertCalls       int `json:"cursor insert calls" bson:"cursor insert calls"`
	CursorNextCalls         int `json:"cursor next calls" bson:"cursor next calls"`
	CursorPrevCalls         int `json:"cursor prev calls" bson:"cursor prev calls"`
	CursorRemoveCalls       int `json:"cursor remove calls" bson:"cursor remove calls"`
	CursorResetCalls        int `json:"cursor reset calls" bson:"cursor reset calls"`
	CursorRestartedSearches int `json:"cursor restarted searches" bson:"cursor restarted searches"`
	CursorSearchCalls       int `json:"cursor search calls" bson:"cursor search calls"`
	CursorSearchNearCalls   int `json:"cursor search near calls" bson:"cursor search near calls"`
	CursorUpdateCalls       int `json:"cursor update calls" bson:"cursor update calls"`
	TruncateCalls           int `json:"truncate calls" bson:"truncate calls"`
}

// DataHandle contains stats about DataHandle
type DataHandle struct {
	ConnectionDataHandlesCurrentlyActive       int `json:"connection data handles currently active" bson:"connection data handles currently active"`
	ConnectionSweepCandidateBecameReferenced   int `json:"connection sweep candidate became referenced" bson:"connection sweep candidate became referenced"`
	ConnectionSweepDhandlesClosed              int `json:"connection sweep dhandles closed" bson:"connection sweep dhandles closed"`
	ConnectionSweepDhandlesRemovedFromHashList int `json:"connection sweep dhandles removed from hash list" bson:"connection sweep dhandles removed from hash list"`
	ConnectionSweepTimeOfDeathSets             int `json:"connection sweep time-of-death sets" bson:"connection sweep time-of-death sets"`
	ConnectionSweeps                           int `json:"connection sweeps" bson:"connection sweeps"`
	SessionDhandlesSwept                       int `json:"session dhandles swept" bson:"session dhandles swept"`
	SessionSweepAttempts                       int `json:"session sweep attempts" bson:"session sweep attempts"`
}

// Log contains stats about Log
type Log struct {
	BusyReturnsAttemptingToSwitchSlots    int `json:"busy returns attempting to switch slots" bson:"busy returns attempting to switch slots"`
	ConsolidatedSlotClosures              int `json:"consolidated slot closures" bson:"consolidated slot closures"`
	ConsolidatedSlotJoinRaces             int `json:"consolidated slot join races" bson:"consolidated slot join races"`
	ConsolidatedSlotJoinTransitions       int `json:"consolidated slot join transitions" bson:"consolidated slot join transitions"`
	ConsolidatedSlotJoins                 int `json:"consolidated slot joins" bson:"consolidated slot joins"`
	ConsolidatedSlotUnbufferedWrites      int `json:"consolidated slot unbuffered writes" bson:"consolidated slot unbuffered writes"`
	LogBytesOfPayloadData                 int `json:"log bytes of payload data" bson:"log bytes of payload data"`
	LogBytesWritten                       int `json:"log bytes written" bson:"log bytes written"`
	LogFilesManuallyZeroFilled            int `json:"log files manually zero-filled" bson:"log files manually zero-filled"`
	LogFlushOperations                    int `json:"log flush operations" bson:"log flush operations"`
	LogForceWriteOperations               int `json:"log force write operations" bson:"log force write operations"`
	LogForceWriteOperationsSkipped        int `json:"log force write operations skipped" bson:"log force write operations skipped"`
	LogRecordsCompressed                  int `json:"log records compressed" bson:"log records compressed"`
	LogRecordsNotCompressed               int `json:"log records not compressed" bson:"log records not compressed"`
	LogRecordsTooSmallToCompress          int `json:"log records too small to compress" bson:"log records too small to compress"`
	LogReleaseAdvancesWriteLSN            int `json:"log release advances write LSN" bson:"log release advances write LSN"`
	LogScanOperations                     int `json:"log scan operations" bson:"log scan operations"`
	LogScanRecordsRequiringTwoReads       int `json:"log scan records requiring two reads" bson:"log scan records requiring two reads"`
	LogServerThreadAdvancesWriteLSN       int `json:"log server thread advances write LSN" bson:"log server thread advances write LSN"`
	LogServerThreadWriteLSNWalkSkipped    int `json:"log server thread write LSN walk skipped" bson:"log server thread write LSN walk skipped"`
	LogSyncOperations                     int `json:"log sync operations" bson:"log sync operations"`
	LogSyncTimeDurationUsecs              int `json:"log sync time duration (usecs)" bson:"log sync time duration (usecs)"`
	LogSyncDirOperations                  int `json:"log sync_dir operations" bson:"log sync_dir operations"`
	LogSyncDirTimeDurationUsecs           int `json:"log sync_dir time duration (usecs)" bson:"log sync_dir time duration (usecs)"`
	LogWriteOperations                    int `json:"log write operations" bson:"log write operations"`
	LoggingBytesConsolidated              int `json:"logging bytes consolidated" bson:"logging bytes consolidated"`
	MaximumLogFileSize                    int `json:"maximum log file size" bson:"maximum log file size"`
	NumberOfPreAllocatedLogFilesToCreate  int `json:"number of pre-allocated log files to create" bson:"number of pre-allocated log files to create"`
	PreAllocatedLogFilesNotReadyAndMissed int `json:"pre-allocated log files not ready and missed" bson:"pre-allocated log files not ready and missed"`
	PreAllocatedLogFilesPrepared          int `json:"pre-allocated log files prepared" bson:"pre-allocated log files prepared"`
	PreAllocatedLogFilesUsed              int `json:"pre-allocated log files used" bson:"pre-allocated log files used"`
	RecordsProcessedByLogScan             int `json:"records processed by log scan" bson:"records processed by log scan"`
	TotalInMemorySizeOfCompressedRecords  int `json:"total in-memory size of compressed records" bson:"total in-memory size of compressed records"`
	TotalLogBufferSize                    int `json:"total log buffer size" bson:"total log buffer size"`
	TotalSizeOfCompressedRecords          int `json:"total size of compressed records" bson:"total size of compressed records"`
	WrittenSlotsCoalesced                 int `json:"written slots coalesced" bson:"written slots coalesced"`
	YieldsWaitingForPreviousLogFileClose  int `json:"yields waiting for previous log file close" bson:"yields waiting for previous log file close"`
}

// Reconciliation contains stats about Reconciliation
type Reconciliation struct {
	FastPathPagesDeleted               int `json:"fast-path pages deleted" bson:"fast-path pages deleted"`
	PageReconciliationCalls            int `json:"page reconciliation calls" bson:"page reconciliation calls"`
	PageReconciliationCallsForEviction int `json:"page reconciliation calls for eviction" bson:"page reconciliation calls for eviction"`
	PagesDeleted                       int `json:"pages deleted" bson:"pages deleted"`
	SplitBytesCurrentlyAwaitingFree    int `json:"split bytes currently awaiting free" bson:"split bytes currently awaiting free"`
	SplitObjectsCurrentlyAwaitingFree  int `json:"split objects currently awaiting free" bson:"split objects currently awaiting free"`
}

// Session contains stats about Session
type Session struct {
	OpenCursorCount               int `json:"open cursor count" bson:"open cursor count"`
	OpenSessionCount              int `json:"open session count" bson:"open session count"`
	TableCompactFailedCalls       int `json:"table compact failed calls" bson:"table compact failed calls"`
	TableCompactSuccessfulCalls   int `json:"table compact successful calls" bson:"table compact successful calls"`
	TableCreateFailedCalls        int `json:"table create failed calls" bson:"table create failed calls"`
	TableCreateSuccessfulCalls    int `json:"table create successful calls" bson:"table create successful calls"`
	TableDropFailedCalls          int `json:"table drop failed calls" bson:"table drop failed calls"`
	TableDropSuccessfulCalls      int `json:"table drop successful calls" bson:"table drop successful calls"`
	TableRebalanceFailedCalls     int `json:"table rebalance failed calls" bson:"table rebalance failed calls"`
	TableRebalanceSuccessfulCalls int `json:"table rebalance successful calls" bson:"table rebalance successful calls"`
	TableRenameFailedCalls        int `json:"table rename failed calls" bson:"table rename failed calls"`
	TableRenameSuccessfulCalls    int `json:"table rename successful calls" bson:"table rename successful calls"`
	TableSalvageFailedCalls       int `json:"table salvage failed calls" bson:"table salvage failed calls"`
	TableSalvageSuccessfulCalls   int `json:"table salvage successful calls" bson:"table salvage successful calls"`
	TableTruncateFailedCalls      int `json:"table truncate failed calls" bson:"table truncate failed calls"`
	TableTruncateSuccessfulCalls  int `json:"table truncate successful calls" bson:"table truncate successful calls"`
	TableVerifyFailedCalls        int `json:"table verify failed calls" bson:"table verify failed calls"`
	TableVerifySuccessfulCalls    int `json:"table verify successful calls" bson:"table verify successful calls"`
}

// ThreadState contains stats about ThreadState
type ThreadState struct {
	ActiveFilesystemFsyncCalls int `json:"active filesystem fsync calls" bson:"active filesystem fsync calls"`
	ActiveFilesystemReadCalls  int `json:"active filesystem read calls" bson:"active filesystem read calls"`
	ActiveFilesystemWriteCalls int `json:"active filesystem write calls" bson:"active filesystem write calls"`
}

// ThreadYield contains stats about ThreadYield
type ThreadYield struct {
	PageAcquireBusyBlocked       int `json:"page acquire busy blocked" bson:"page acquire busy blocked"`
	PageAcquireEvictionBlocked   int `json:"page acquire eviction blocked" bson:"page acquire eviction blocked"`
	PageAcquireLockedBlocked     int `json:"page acquire locked blocked" bson:"page acquire locked blocked"`
	PageAcquireReadBlocked       int `json:"page acquire read blocked" bson:"page acquire read blocked"`
	PageAcquireTimeSleepingUsecs int `json:"page acquire time sleeping (usecs)" json:"page acquire time sleeping (usecs)"`
}

// Transaction contains stats about Transaction
type Transaction struct {
	NumberOfNamedSnapshotsCreated                                             int `json:"number of named snapshots created" bson:"number of named snapshots created"`
	NumberOfNamedSnapshotsDropped                                             int `json:"number of named snapshots dropped" bson:"number of named snapshots dropped"`
	TransactionBegins                                                         int `json:"transaction begins" bson:"transaction begins"`
	TransactionCheckpointCurrentlyRunning                                     int `json:"transaction checkpoint currently running" bson:"transaction checkpoint currently running"`
	TransactionCheckpointGeneration                                           int `json:"transaction checkpoint generation" bson:"transaction checkpoint generation"`
	TransactionCheckpointMaxTimeMsecs                                         int `json:"transaction checkpoint max time (msecs)" bson:"transaction checkpoint max time (msecs)"`
	TransactionCheckpointMinTimeMsecs                                         int `json:"transaction checkpoint min time (msecs)" bson:"transaction checkpoint min time (msecs)"`
	TransactionCheckpointMostRecentTimeMsecs                                  int `json:"transaction checkpoint most recent time (msecs)" bson:"transaction checkpoint most recent time (msecs)"`
	TransactionCheckpointScrubDirtyTarget                                     int `json:"transaction checkpoint scrub dirty target" bson:"transaction checkpoint scrub dirty target"`
	TransactionCheckpointScrubTimeMsecs                                       int `json:"transaction checkpoint scrub time (msecs)" bson:"transaction checkpoint scrub time (msecs)"`
	TransactionCheckpointTotalTimeMsecs                                       int `json:"transaction checkpoint total time (msecs)" bson:"transaction checkpoint total time (msecs)"`
	TransactionCheckpoints                                                    int `json:"transaction checkpoints" bson:"transaction checkpoints"`
	TransactionFailuresDueToCacheOverflow                                     int `json:"transaction failures due to cache overflow" bson:"transaction failures due to cache overflow"`
	TransactionFsyncCallsForCheckpointAfterAllocatingTheTransactionID         int `json:"transaction fsync calls for checkpoint after allocating the transaction ID" bson:"transaction fsync calls for checkpoint after allocating the transaction ID"`
	TransactionFsyncDurationForCheckpointAfterAllocatingTheTransactionIDUsecs int `json:"transaction fsync duration for checkpoint after allocating the transaction ID (usecs)" bson:"transaction fsync duration for checkpoint after allocating the transaction ID (usecs)"`
	TransactionRangeOfIDsCurrentlyPinned                                      int `json:"transaction range of IDs currently pinned" bson:"transaction range of IDs currently pinned"`
	TransactionRangeOfIDsCurrentlyPinnedByACheckpoint                         int `json:"transaction range of IDs currently pinned by a checkpoint" bson:"transaction range of IDs currently pinned by a checkpoint"`
	TransactionRangeOfIDsCurrentlyPinnedByNamedSnapshots                      int `json:"transaction range of IDs currently pinned by named snapshots" bson:"transaction range of IDs currently pinned by named snapshots"`
	TransactionSyncCalls                                                      int `json:"transaction sync calls" bson:"transaction sync calls"`
	TransactionsCommitted                                                     int `json:"transactions committed" bson:"transactions committed"`
	TransactionsRolledBack                                                    int `json:"transactions rolled back" bson:"transactions rolled back"`
}

// Write contains stats about Write
type Write struct {
	Out          int `json:"out" bson:"out"`
	Available    int `json:"available" bson:"available"`
	TotalTickets int `json:"totalTickets" bson:"totalTickets"`
}

// Read contains stats about Read
type Read struct {
	Out          int `json:"out" bson:"out"`
	Available    int `json:"available" bson:"available"`
	TotalTickets int `json:"totalTickets" bson:"totalTickets"`
}

// ConcurrentTransactions contains stats about ConcurrentTransactions
type ConcurrentTransactions struct {
	Write *Write `json:"write" bson:"write"`
	Read  *Read  `json:"read" bson:"read"`
}

// WiredTiger contains stats about WiredTiger
type WiredTiger struct {
	URI                    string                  `json:"uri" bson:"uri"`
	LSM                    *LSM                    `json:"LSM" bson:"LSM"`
	Async                  *Async                  `json:"async" bson:"async"`
	BlockManager           *BlockManager           `json:"block-manager" bson:"block-manager"`
	Cache                  *Cache                  `json:"cache" bson:"cache"`
	Connection             *Connection             `json:"connection" bson:"connection"`
	Cursor                 *WireCursor             `json:"cursor" bson:"cursor"`
	DataHandle             *DataHandle             `json:"data-handle" bson:"data-handle"`
	Log                    *Log                    `json:"log" bson:"log"`
	Reconciliation         *Reconciliation         `json:"reconciliation" bson:"reconciliation"`
	Session                *Session                `json:"session" bson:"session"`
	ThreadState            *ThreadState            `json:"thread-state" bson:"thread-state"`
	ThreadYield            *ThreadYield            `json:"thread-yield" bson:"thread-yield"`
	Transaction            *Transaction            `json:"transaction" bson:"transaction"`
	ConcurrentTransactions *ConcurrentTransactions `json:"concurrentTransactions" bson:"concurrentTransactions"`
}

// BuildInfo contains stats about BuildInfo
type BuildInfo struct {
	Failed int `json:"failed" bson:"failed"`
	Total  int `json:"total" bson:"total"`
}

// GetLog contains stats about GetLog
type GetLog struct {
	Failed int `json:"failed" bson:"failed"`
	Total  int `json:"total" bson:"total"`
}

// IsMaster contains stats about IsMaster
type IsMaster struct {
	Failed int `json:"failed" bson:"failed"`
	Total  int `json:"total" bson:"total"`
}

// ReplSetGetStatus contains stats about ReplSetGetStatus
type ReplSetGetStatus struct {
	Failed int `json:"failed" bson:"failed"`
	Total  int `json:"total" bson:"total"`
}

// ServerStatusR contains stats about ServerStatusR
type ServerStatusR struct {
	Failed int `json:"failed" bson:"failed"`
	Total  int `json:"total" bson:"total"`
}

// Whatsmyuri contains stats about Whatsmyuri
type Whatsmyuri struct {
	Failed int `json:"failed" bson:"failed"`
	Total  int `json:"total" bson:"total"`
}

// Commands contains stats about Commands
type Commands struct {
	BuildInfo        *BuildInfo        `json:"buildInfo" bson:"buildInfo"`
	GetLog           *GetLog           `json:"getLog" bson:"getLog"`
	IsMaster         *IsMaster         `json:"isMaster" bson:"isMaster"`
	ReplSetGetStatus *ReplSetGetStatus `json:"replSetGetStatus" bson:"replSetGetStatus"`
	ServerStatus     *ServerStatusR    `json:"serverStatus" bson:"serverStatus"`
	Whatsmyuri       *Whatsmyuri       `json:"whatsmyuri" bson:"whatsmyuri"`
}

// Cursor contains stats about Cursor
type Cursor struct {
	TimedOut int   `json:"timedOut" bson:"timedOut"`
	Open     *Open `json:"open" bson:"open"`
}

// Open contains stats about Open
type Open struct {
	NoTimeout int `json:"noTimeout" bson:"noTimeout"`
	Pinned    int `json:"pinned" bson:"pinned"`
	Total     int `json:"total" bson:"total"`
}

// Document contains stats about Document
type Document struct {
	Deleted  int `json:"deleted" bson:"deleted"`
	Inserted int `json:"inserted" bson:"inserted"`
	Returned int `json:"returned" bson:"returned"`
	Updated  int `json:"updated" bson:"updated"`
}

// Wtime contains stats about Wtime
type Wtime struct {
	Num         int `json:"num" bson:"num"`
	TotalMillis int `json:"totalMillis" bson:"totalMillis"`
}

// GetLastError contains stats about GetLastError
type GetLastError struct {
	Wtime     *Wtime `json:"wtime" bson:"wtime"`
	Wtimeouts int    `json:"wtimeouts" bson:"wtimeouts"`
}

// Operation contains stats about Operation
type Operation struct {
	Fastmod        int `json:"fastmod" bson:"fastmod"`
	Idhack         int `json:"idhack" bson:"idhack"`
	ScanAndOrder   int `json:"scanAndOrder" bson:"scanAndOrder"`
	WriteConflicts int `json:"writeConflicts" bson:"writeConflicts"`
}

// QueryExecutor contains stats about QueryExecutor
type QueryExecutor struct {
	Scanned        int `json:"scanned" bson:"scanned"`
	ScannedObjects int `json:"scannedObjects" bson:"scannedObjects"`
}

// Record contains stats about Record
type Record struct {
	Moves int `json:"moves" bson:"moves"`
}

// Counters contains stats about Counters
type Counters struct {
	EventCreated       int `json:"eventCreated" bson:"eventCreated"`
	EventWait          int `json:"eventWait" bson:"eventWait"`
	Cancels            int `json:"cancels" bson:"cancels"`
	Waits              int `json:"waits" bson:"waits"`
	ScheduledNetCmd    int `json:"scheduledNetCmd" bson:"scheduledNetCmd"`
	ScheduledDBWork    int `json:"scheduledDBWork" bson:"scheduledDBWork"`
	ScheduledXclWork   int `json:"scheduledXclWork" bson:"scheduledXclWork"`
	ScheduledWorkAt    int `json:"scheduledWorkAt" bson:"scheduledWorkAt"`
	ScheduledWork      int `json:"scheduledWork" bson:"scheduledWork"`
	SchedulingFailures int `json:"schedulingFailures" bson:"schedulingFailures"`
}

// Queues contains stats about Queues
type Queues struct {
	NetworkInProgress   int `json:"networkInProgress" bson:"networkInProgress"`
	DbWorkInProgress    int `json:"dbWorkInProgress" bson:"dbWorkInProgress"`
	ExclusiveInProgress int `json:"exclusiveInProgress" bson:"exclusiveInProgress"`
	Sleepers            int `json:"sleepers" bson:"sleepers"`
	Ready               int `json:"ready" bson:"ready"`
	Free                int `json:"free" bson:"free"`
}

// Executor contains stats about Executor
type Executor struct {
	Counters         *Counters `json:"counters" bson:"counters"`
	Queues           *Queues   `json:"queues" bson:"queues"`
	UnsignaledEvents int       `json:"unsignaledEvents" bson:"unsignaledEvents"`
	EventWaiters     int       `json:"eventWaiters" bson:"eventWaiters"`
	ShuttingDown     bool      `json:"shuttingDown" bson:"shuttingDown"`
	NetworkInterface string    `json:"networkInterface" bson:"networkInterface"`
}

// Batches contains stats about Batches
type Batches struct {
	Num         int `json:"num" bson:"num"`
	TotalMillis int `json:"totalMillis" bson:"totalMillis"`
}

// Apply contains stats about Apply
type Apply struct {
	Batches *Batches `json:"batches" bson:"batches"`
	Ops     int      `json:"ops" bson:"ops"`
}

// Buffer contains stats about Buffer
type Buffer struct {
	Count        int `json:"count" bson:"count"`
	MaxSizeBytes int `json:"maxSizeBytes" bson:"maxSizeBytes"`
	SizeBytes    int `json:"sizeBytes" bson:"sizeBytes"`
}

// Getmores contains stats about Getmores
type Getmores struct {
	Num         int `json:"num" bson:"num"`
	TotalMillis int `json:"totalMillis" bson:"totalMillis"`
}

// ReplNetwork contains stats about ReplNetwork
type ReplNetwork struct {
	Bytes          int       `json:"bytes" bson:"bytes"`
	Getmores       *Getmores `json:"getmores" bson:"getmores"`
	Ops            int       `json:"ops" bson:"ops"`
	ReadersCreated int       `json:"readersCreated" bson:"readersCreated"`
}

// Docs contains stats about Docs
type Docs struct {
	Num         int `json:"num" bson:"num"`
	TotalMillis int `json:"totalMillis" bson:"totalMillis"`
}

// Indexes contains stats about Indexes
type Indexes struct {
	Num         int `json:"num" bson:"num"`
	TotalMillis int `json:"totalMillis" bson:"totalMillis"`
}

// Preload contains stats about Preload
type Preload struct {
	Docs    *Docs    `json:"docs" bson:"docs"`
	Indexes *Indexes `json:"indexes" bson:"indexes"`
}

// Repl contains stats about Repl
type Repl struct {
	Executor *Executor    `json:"executor" bson:"executor"`
	Apply    *Apply       `json:"apply" bson:"apply"`
	Buffer   *Buffer      `json:"buffer" bson:"buffer"`
	Network  *ReplNetwork `json:"network" bson:"network"`
	Preload  *Preload     `json:"preload" bson:"preload"`
}

// Search contains stats about Search
type Search struct {
	BucketExhausted int `json:"bucketExhausted" bson:"bucketExhausted"`
	Requests        int `json:"requests" bson:"requests"`
	Scanned         int `json:"scanned" bson:"scanned"`
}

// Freelist contains stats about Freelist
type Freelist struct {
	Search *Search `json:"search" bson:"search"`
}

// Storage contains stats about Storage
type Storage struct {
	Freelist *Freelist `json:"freelist" bson:"freelist"`
}

// TTL contains stats about TTL
type TTL struct {
	DeletedDocuments int `json:"deletedDocuments" bson:"deletedDocuments"`
	Passes           int `json:"passes" bson:"passes"`
}

// Metrics contains stats about Metrics
type Metrics struct {
	Commands      *Commands      `json:"commands" bson:"commands"`
	Cursor        *Cursor        `json:"cursor" bson:"cursor"`
	Document      *Document      `json:"document" bson:"document"`
	GetLastError  *GetLastError  `json:"getLastError" bson:"getLastError"`
	Operation     *Operation     `json:"operation" bson:"operation"`
	QueryExecutor *QueryExecutor `json:"queryExecutor" bson:"queryExecutor"`
	Record        *Record        `json:"record" bson:"record"`
	Repl          *Repl          `json:"repl" bson:"repl"`
	Storage       *Storage       `json:"storage" bson:"storage"`
	TTL           *TTL           `json:"ttl" bson:"ttl"`
}

// ServerStatus contains stats about database
type ServerStatus struct {
	Host              string          `json:"host" bson:"host"`
	AdvisoryHostFQDNs []interface{}   `json:"advisoryHostFQDNs" bson:"advisoryHostFQDNs"`
	Version           string          `json:"version" bson:"version"`
	Process           string          `json:"process" bson:"process"`
	Pid               int             `json:"pid" bson:"pid"`
	Uptime            int             `json:"uptime" bson:"uptime"`
	UptimeMillis      int             `json:"uptimeMillis" bson:"uptimeMillis"`
	UptimeEstimate    int             `json:"uptimeEstimate" bson:"uptimeEstimate"`
	LocalTime         time.Time       `json:"localTime" bson:"localTime"`
	Asserts           *Asserts        `json:"asserts" bson:"asserts"`
	Connections       *Connections    `json:"connections" bson:"connections"`
	Extra_Info        *Extra_Info     `json:"extra_info" bson:"extra_info"`
	GlobalLock        *GlobalLock     `json:"globalLock" bson:"globalLock"`
	Locks             *Locks          `json:"locks" bson:"locks"`
	Network           *Network        `json:"network" bson:"network"`
	Opcounters        *Opcounters     `json:"opcounters" bson:"opcounters"`
	OpcountersRepl    *OpcountersRepl `json:"opcountersRepl" bson:"opcountersRepl"`
	StorageEngine     *StorageEngine  `json:"storageEngine" bson:"storageEngine"`
	Tcmalloc          *Tcmalloc       `json:"tcmalloc" bson:"tcmalloc"`
	WiredTiger        *WiredTiger     `json:"wiredTiger" bson:"wiredTiger"`
	WriteBacksQueued  bool            `json:"writeBacksQueued" bson:"writeBacksQueued"`
	Mem               *Mem            `json:"mem" bson:"mem"`
	Metrics           *Metrics        `json:"metrics" bson:"metrics"`
	Ok                int             `json:"ok" bson:"ok"`
}
