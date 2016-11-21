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
	Regular   int `json:"regular"`
	Warning   int `json:"warning"`
	Msg       int `json:"msg"`
	User      int `json:"user"`
	Rollovers int `json:"rollovers"`
}

//Connections contains info about Connections
type Connections struct {
	Current      int `json:"current"`
	Available    int `json:"available"`
	TotalCreated int `json:"totalCreated"`
}

//Extra_Info contains info about ExtraInfo
type Extra_Info struct {
	Note           string `json:"note"`
	HeapUsageBytes int    `json:"heap_usage_bytes"`
	PageFaults     int    `json:"page_faults"`
}

//GlobalLock contains info about GlobalLock
type GlobalLock struct {
	TotalTime     int            `json:"totalTime"`
	CurrentQueue  *CurrentQueue  `json:"currentQueue"`
	ActiveClients *ActiveClients `json:"activeClients"`
}

//ActiveClients contains info about ActiveClients
type ActiveClients struct {
	Total   int `json:"total"`
	Readers int `json:"readers"`
	Writers int `json:"writers"`
}

//CurrentQueue contains info about CurrentQueue
type CurrentQueue struct {
	Total   int `json:"total"`
	Readers int `json:"readers"`
	Writers int `json:"writers"`
}

//Opcounters contains info about Opcounters
type Opcounters struct {
	Insert  int `json:"insert"`
	Query   int `json:"query"`
	Update  int `json:"update"`
	Delete  int `json:"delete"`
	Getmore int `json:"getmore"`
	Command int `json:"command"`
}

// AcquireCount contains stats about AcquireCount
type AcquireCount struct {
	R int `json:"r"`
	W int `json:"w"`
}

// Global contains stats about Global
type Global struct {
	AcquireCount *AcquireCount `json:"acquireCount"`
}

// Database contains stats about Database
type Database struct {
	AcquireCount *AcquireCount `json:"acquireCount"`
}

// Collection contains stats about Collection
type Collection struct {
	AcquireCount *AcquireCount `json:"acquireCount"`
}

// Metadata contains stats about Metadata
type Metadata struct {
	AcquireCount *AcquireCount `json:"acquireCount"`
}

// Locks contains stats about Locks
type Locks struct {
	Global     *Global     `json:"Global"`
	Database   *Database   `json:"Database"`
	Collection *Collection `json:"Collection"`
	Metadata   *Metadata   `json:"Metadata"`
}

// OpcountersRepl contains stats about OpcountersRepl
type OpcountersRepl struct {
	Insert  int `json:"insert"`
	Query   int `json:"query"`
	Update  int `json:"update"`
	Delete  int `json:"delete"`
	Getmore int `json:"getmore"`
	Command int `json:"command"`
}

// Network contains stats about Network
type Network struct {
	BytesIn     int `json:"bytesIn"`
	BytesOut    int `json:"bytesOut"`
	NumRequests int `json:"numRequests"`
}

// StorageEngine contains stats about StorageEngine
type StorageEngine struct {
	Name                   string `json:"name"`
	SupportsCommittedReads bool   `json:"supportsCommittedReads"`
	Persistent             bool   `json:"persistent"`
}

// Generic contains stats about Generic
type Generic struct {
	CurrentAllocatedBytes int `json:"current_allocated_bytes"`
	HeapSize              int `json:"heap_size"`
}

// TcmallocStat contains stats about TcmallocStat
type TcmallocStat struct {
	PageheapFreeBytes            int    `json:"pageheap_free_bytes"`
	PageheapUnmappedBytes        int    `json:"pageheap_unmapped_bytes"`
	MaxTotalThreadCacheBytes     int    `json:"max_total_thread_cache_bytes"`
	CurrentTotalThreadCacheBytes int    `json:"current_total_thread_cache_bytes"`
	TotalFreeBytes               int    `json:"total_free_bytes"`
	CentralCacheFreeBytes        int    `json:"central_cache_free_bytes"`
	TransferCacheFreeBytes       int    `json:"transfer_cache_free_bytes"`
	ThreadCacheFreeBytes         int    `json:"thread_cache_free_bytes"`
	AggressiveMemoryDecommit     int    `json:"aggressive_memory_decommit"`
	FormattedString              string `json:"formattedString"`
}

// Tcmalloc contains stats about Tcmalloc
type Tcmalloc struct {
	Generic  *Generic      `json:"generic"`
	Tcmalloc *TcmallocStat `json:"tcmalloc"`
}

// Mem contains stats about Mem
type Mem struct {
	Bits              int  `json:"bits"`
	Resident          int  `json:"resident"`
	Virtual           int  `json:"virtual"`
	Supported         bool `json:"supported"`
	Mapped            int  `json:"mapped"`
	MappedWithJournal int  `json:"mappedWithJournal"`
}

// LSM contains stats about LSM
type LSM struct {
	ApplicationWorkUnitsCurrentlyQueued int `json:"application work units currently queued"`
	MergeWorkUnitsCurrentlyQueued       int `json:"merge work units currently queued"`
	RowsMergedInAnLSMTree               int `json:"rows merged in an LSM tree"`
	SleepForLSMCheckpointThrottle       int `json:"sleep for LSM checkpoint throttle"`
	SleepForLSMMergeThrottle            int `json:"sleep for LSM merge throttle"`
	SwitchWorkUnitsCurrentlyQueued      int `json:"switch work units currently queued"`
	TreeMaintenanceOperationsDiscarded  int `json:"tree maintenance operations discarded"`
	TreeMaintenanceOperationsExecuted   int `json:"tree maintenance operations executed"`
	TreeMaintenanceOperationsScheduled  int `json:"tree maintenance operations scheduled"`
	TreeQueueHitMaximum                 int `json:"tree queue hit maximum"`
}

// Async contains stats about Async
type Async struct {
	CurrentWorkQueueLength                    int `json:"current work queue length"`
	MaximumWorkQueueLength                    int `json:"maximum work queue length"`
	NumberOfAllocationStateRaces              int `json:"number of allocation state races"`
	NumberOfFlushCalls                        int `json:"number of flush calls"`
	NumberOfOperationSlotsViewedForAllocation int `json:"number of operation slots viewed for allocation"`
	NumberOfTimesOperationAllocationFailed    int `json:"number of times operation allocation failed"`
	NumberOfTimesWorkerFoundNoWork            int `json:"number of times worker found no work"`
	TotalAllocations                          int `json:"total allocations"`
	TotalCompactCalls                         int `json:"total compact calls"`
	TotalInsertCalls                          int `json:"total insert calls"`
	TotalRemoveCalls                          int `json:"total remove calls"`
	TotalSearchCalls                          int `json:"total search calls"`
	TotalUpdateCalls                          int `json:"total update calls"`
}

// BlockManager contains stats about BlockManager
type BlockManager struct {
	BlocksPreLoaded           int `json:"blocks pre-loaded"`
	BlocksRead                int `json:"blocks read"`
	BlocksWritten             int `json:"blocks written"`
	BytesRead                 int `json:"bytes read"`
	BytesWritten              int `json:"bytes written"`
	BytesWrittenForCheckpoint int `json:"bytes written for checkpoint"`
	MappedBlocksRead          int `json:"mapped blocks read"`
	MappedBytesRead           int `json:"mapped bytes read"`
}

// Cache contains stats about Cache
type Cache struct {
	BytesBelongingToPageImagesInTheCache                       int `json:"bytes belonging to page images in the cache"`
	BytesCurrentlyInTheCache                                   int `json:"bytes currently in the cache"`
	BytesNotBelongingToPageImagesInTheCache                    int `json:"bytes not belonging to page images in the cache"`
	BytesReadIntoCache                                         int `json:"bytes read into cache"`
	BytesWrittenFromCache                                      int `json:"bytes written from cache"`
	CheckpointBlockedPageEviction                              int `json:"checkpoint blocked page eviction"`
	EvictionCallsToGetAPage                                    int `json:"eviction calls to get a page"`
	EvictionCallsToGetAPageFoundQueueEmpty                     int `json:"eviction calls to get a page found queue empty"`
	EvictionCallsToGetAPageFoundQueueEmptyAfterLocking         int `json:"eviction calls to get a page found queue empty after locking"`
	EvictionCurrentlyOperatingInAggressiveMode                 int `json:"eviction currently operating in aggressive mode"`
	EvictionEmptyScore                                         int `json:"eviction empty score"`
	EvictionServerCandidateQueueEmptyWhenToppingUp             int `json:"eviction server candidate queue empty when topping up"`
	EvictionServerCandidateQueueNotEmptyWhenToppingUp          int `json:"eviction server candidate queue not empty when topping up"`
	EvictionServerEvictingPages                                int `json:"eviction server evicting pages"`
	EvictionServerSleptBecauseWeDidNotMakeProgressWithEviction int `json:"eviction server slept, because we did not make progress with eviction"`
	EvictionServerUnableToReachEvictionGoal                    int `json:"eviction server unable to reach eviction goal"`
	EvictionState                                              int `json:"eviction state"`
	EvictionWalksAbandoned                                     int `json:"eviction walks abandoned"`
	EvictionWorkerThreadEvictingPages                          int `json:"eviction worker thread evicting pages"`
	FailedEvictionOfPagesThatExceededTheInMemoryMaximum        int `json:"failed eviction of pages that exceeded the in-memory maximum"`
	FilesWithActiveEvictionWalks                               int `json:"files with active eviction walks"`
	FilesWithNewEvictionWalksStarted                           int `json:"files with new eviction walks started"`
	HazardPointerBlockedPageEviction                           int `json:"hazard pointer blocked page eviction"`
	HazardPointerCheckCalls                                    int `json:"hazard pointer check calls"`
	HazardPointerCheckEntriesWalked                            int `json:"hazard pointer check entries walked"`
	HazardPointerMaximumArrayLength                            int `json:"hazard pointer maximum array length"`
	InMemoryPagePassedCriteriaToBeSplit                        int `json:"in-memory page passed criteria to be split"`
	InMemoryPageSplits                                         int `json:"in-memory page splits"`
	InternalPagesEvicted                                       int `json:"internal pages evicted"`
	InternalPagesSplitDuringEviction                           int `json:"internal pages split during eviction"`
	LeafPagesSplitDuringEviction                               int `json:"leaf pages split during eviction"`
	LookasideTableInsertCalls                                  int `json:"lookaside table insert calls"`
	LookasideTableRemoveCalls                                  int `json:"lookaside table remove calls"`
	MaximumBytesConfigured                                     int `json:"maximum bytes configured"`
	MaximumPageSizeAtEviction                                  int `json:"maximum page size at eviction"`
	ModifiedPagesEvicted                                       int `json:"modified pages evicted"`
	ModifiedPagesEvictedByApplicationThreads                   int `json:"modified pages evicted by application threads"`
	OverflowPagesReadIntoCache                                 int `json:"overflow pages read into cache"`
	OverflowValuesCachedInMemory                               int `json:"overflow values cached in memory"`
	PageSplitDuringEvictionDeepenedTheTree                     int `json:"page split during eviction deepened the tree"`
	PageWrittenRequiringLookasideRecords                       int `json:"page written requiring lookaside records"`
	PagesCurrentlyHeldInTheCache                               int `json:"pages currently held in the cache"`
	PagesEvictedBecauseTheyExceededTheInMemoryMaximum          int `json:"pages evicted because they exceeded the in-memory maximum"`
	PagesEvictedBecauseTheyHadChainsOfDeletedItems             int `json:"pages evicted because they had chains of deleted items"`
	PagesEvictedByApplicationThreads                           int `json:"pages evicted by application threads"`
	PagesQueuedForEviction                                     int `json:"pages queued for eviction"`
	PagesQueuedForUrgentEviction                               int `json:"pages queued for urgent eviction"`
	PagesQueuedForUrgentEvictionDuringWalk                     int `json:"pages queued for urgent eviction during walk"`
	PagesReadIntoCache                                         int `json:"pages read into cache"`
	PagesReadIntoCacheRequiringLookasideEntries                int `json:"pages read into cache requiring lookaside entries"`
	PagesRequestedFromTheCache                                 int `json:"pages requested from the cache"`
	PagesSeenByEvictionWalk                                    int `json:"pages seen by eviction walk"`
	PagesSelectedForEvictionUnableToBeEvicted                  int `json:"pages selected for eviction unable to be evicted"`
	PagesWalkedForEviction                                     int `json:"pages walked for eviction"`
	PagesWrittenFromCache                                      int `json:"pages written from cache"`
	PagesWrittenRequiringInMemoryRestoration                   int `json:"pages written requiring in-memory restoration"`
	PercentageOverhead                                         int `json:"percentage overhead"`
	TrackedBytesBelongingToInternalPagesInTheCache             int `json:"tracked bytes belonging to internal pages in the cache"`
	TrackedBytesBelongingToLeafPagesInTheCache                 int `json:"tracked bytes belonging to leaf pages in the cache"`
	TrackedDirtyBytesInTheCache                                int `json:"tracked dirty bytes in the cache"`
	TrackedDirtyPagesInTheCache                                int `json:"tracked dirty pages in the cache"`
	UnmodifiedPagesEvicted                                     int `json:"unmodified pages evicted"`
}

// Connection contains stats about Connection
type Connection struct {
	AutoAdjustingConditionResets         int `json:"auto adjusting condition resets"`
	AutoAdjustingConditionWaitCalls      int `json:"auto adjusting condition wait calls"`
	FilesCurrentlyOpen                   int `json:"files currently open"`
	MemoryAllocations                    int `json:"memory allocations"`
	MemoryFrees                          int `json:"memory frees"`
	MemoryReAllocations                  int `json:"memory re-allocations"`
	PthreadMutexConditionWaitCalls       int `json:"pthread mutex condition wait calls"`
	PthreadMutexSharedLockReadLockCalls  int `json:"pthread mutex shared lock read-lock calls"`
	PthreadMutexSharedLockWriteLockCalls int `json:"pthread mutex shared lock write-lock calls"`
	TotalFsyncIOs                        int `json:"total fsync I/Os"`
	TotalReadIOs                         int `json:"total read I/Os"`
	TotalWriteIOs                        int `json:"total write I/Os"`
}

// WireCursor contains stats about WireCursor
type WireCursor struct {
	CursorCreateCalls       int `json:"cursor create calls"`
	CursorInsertCalls       int `json:"cursor insert calls"`
	CursorNextCalls         int `json:"cursor next calls"`
	CursorPrevCalls         int `json:"cursor prev calls"`
	CursorRemoveCalls       int `json:"cursor remove calls"`
	CursorResetCalls        int `json:"cursor reset calls"`
	CursorRestartedSearches int `json:"cursor restarted searches"`
	CursorSearchCalls       int `json:"cursor search calls"`
	CursorSearchNearCalls   int `json:"cursor search near calls"`
	CursorUpdateCalls       int `json:"cursor update calls"`
	TruncateCalls           int `json:"truncate calls"`
}

// DataHandle contains stats about DataHandle
type DataHandle struct {
	ConnectionDataHandlesCurrentlyActive       int `json:"connection data handles currently active"`
	ConnectionSweepCandidateBecameReferenced   int `json:"connection sweep candidate became referenced"`
	ConnectionSweepDhandlesClosed              int `json:"connection sweep dhandles closed"`
	ConnectionSweepDhandlesRemovedFromHashList int `json:"connection sweep dhandles removed from hash list"`
	ConnectionSweepTimeOfDeathSets             int `json:"connection sweep time-of-death sets"`
	ConnectionSweeps                           int `json:"connection sweeps"`
	SessionDhandlesSwept                       int `json:"session dhandles swept"`
	SessionSweepAttempts                       int `json:"session sweep attempts"`
}

// Log contains stats about Log
type Log struct {
	BusyReturnsAttemptingToSwitchSlots    int `json:"busy returns attempting to switch slots"`
	ConsolidatedSlotClosures              int `json:"consolidated slot closures"`
	ConsolidatedSlotJoinRaces             int `json:"consolidated slot join races"`
	ConsolidatedSlotJoinTransitions       int `json:"consolidated slot join transitions"`
	ConsolidatedSlotJoins                 int `json:"consolidated slot joins"`
	ConsolidatedSlotUnbufferedWrites      int `json:"consolidated slot unbuffered writes"`
	LogBytesOfPayloadData                 int `json:"log bytes of payload data"`
	LogBytesWritten                       int `json:"log bytes written"`
	LogFilesManuallyZeroFilled            int `json:"log files manually zero-filled"`
	LogFlushOperations                    int `json:"log flush operations"`
	LogForceWriteOperations               int `json:"log force write operations"`
	LogForceWriteOperationsSkipped        int `json:"log force write operations skipped"`
	LogRecordsCompressed                  int `json:"log records compressed"`
	LogRecordsNotCompressed               int `json:"log records not compressed"`
	LogRecordsTooSmallToCompress          int `json:"log records too small to compress"`
	LogReleaseAdvancesWriteLSN            int `json:"log release advances write LSN"`
	LogScanOperations                     int `json:"log scan operations"`
	LogScanRecordsRequiringTwoReads       int `json:"log scan records requiring two reads"`
	LogServerThreadAdvancesWriteLSN       int `json:"log server thread advances write LSN"`
	LogServerThreadWriteLSNWalkSkipped    int `json:"log server thread write LSN walk skipped"`
	LogSyncOperations                     int `json:"log sync operations"`
	LogSyncTimeDurationUsecs              int `json:"log sync time duration (usecs"`
	LogSyncDirOperations                  int `json:"log sync_dir operations"`
	LogSyncDirTimeDurationUsecs           int `json:"log sync_dir time duration (usecs"`
	LogWriteOperations                    int `json:"log write operations"`
	LoggingBytesConsolidated              int `json:"logging bytes consolidated"`
	MaximumLogFileSize                    int `json:"maximum log file size"`
	NumberOfPreAllocatedLogFilesToCreate  int `json:"number of pre-allocated log files to create"`
	PreAllocatedLogFilesNotReadyAndMissed int `json:"pre-allocated log files not ready and missed"`
	PreAllocatedLogFilesPrepared          int `json:"pre-allocated log files prepared"`
	PreAllocatedLogFilesUsed              int `json:"pre-allocated log files used"`
	RecordsProcessedByLogScan             int `json:"records processed by log scan"`
	TotalInMemorySizeOfCompressedRecords  int `json:"total in-memory size of compressed records"`
	TotalLogBufferSize                    int `json:"total log buffer size"`
	TotalSizeOfCompressedRecords          int `json:"total size of compressed records"`
	WrittenSlotsCoalesced                 int `json:"written slots coalesced"`
	YieldsWaitingForPreviousLogFileClose  int `json:"yields waiting for previous log file close"`
}

// Reconciliation contains stats about Reconciliation
type Reconciliation struct {
	FastPathPagesDeleted               int `json:"fast-path pages deleted"`
	PageReconciliationCalls            int `json:"page reconciliation calls"`
	PageReconciliationCallsForEviction int `json:"page reconciliation calls for eviction"`
	PagesDeleted                       int `json:"pages deleted"`
	SplitBytesCurrentlyAwaitingFree    int `json:"split bytes currently awaiting free"`
	SplitObjectsCurrentlyAwaitingFree  int `json:"split objects currently awaiting free"`
}

// Session contains stats about Session
type Session struct {
	OpenCursorCount               int `json:"open cursor count"`
	OpenSessionCount              int `json:"open session count"`
	TableCompactFailedCalls       int `json:"table compact failed calls"`
	TableCompactSuccessfulCalls   int `json:"table compact successful calls"`
	TableCreateFailedCalls        int `json:"table create failed calls"`
	TableCreateSuccessfulCalls    int `json:"table create successful calls"`
	TableDropFailedCalls          int `json:"table drop failed calls"`
	TableDropSuccessfulCalls      int `json:"table drop successful calls"`
	TableRebalanceFailedCalls     int `json:"table rebalance failed calls"`
	TableRebalanceSuccessfulCalls int `json:"table rebalance successful calls"`
	TableRenameFailedCalls        int `json:"table rename failed calls"`
	TableRenameSuccessfulCalls    int `json:"table rename successful calls"`
	TableSalvageFailedCalls       int `json:"table salvage failed calls"`
	TableSalvageSuccessfulCalls   int `json:"table salvage successful calls"`
	TableTruncateFailedCalls      int `json:"table truncate failed calls"`
	TableTruncateSuccessfulCalls  int `json:"table truncate successful calls"`
	TableVerifyFailedCalls        int `json:"table verify failed calls"`
	TableVerifySuccessfulCalls    int `json:"table verify successful calls"`
}

// ThreadState contains stats about ThreadState
type ThreadState struct {
	ActiveFilesystemFsyncCalls int `json:"active filesystem fsync calls"`
	ActiveFilesystemReadCalls  int `json:"active filesystem read calls"`
	ActiveFilesystemWriteCalls int `json:"active filesystem write calls"`
}

// ThreadYield contains stats about ThreadYield
type ThreadYield struct {
	PageAcquireBusyBlocked       int `json:"page acquire busy blocked"`
	PageAcquireEvictionBlocked   int `json:"page acquire eviction blocked"`
	PageAcquireLockedBlocked     int `json:"page acquire locked blocked"`
	PageAcquireReadBlocked       int `json:"page acquire read blocked"`
	PageAcquireTimeSleepingUsecs int `json:"page acquire time sleeping (usecs"`
}

// Transaction contains stats about Transaction
type Transaction struct {
	NumberOfNamedSnapshotsCreated                                             int `json:"number of named snapshots created"`
	NumberOfNamedSnapshotsDropped                                             int `json:"number of named snapshots dropped"`
	TransactionBegins                                                         int `json:"transaction begins"`
	TransactionCheckpointCurrentlyRunning                                     int `json:"transaction checkpoint currently running"`
	TransactionCheckpointGeneration                                           int `json:"transaction checkpoint generation"`
	TransactionCheckpointMaxTimeMsecs                                         int `json:"transaction checkpoint max time (msecs"`
	TransactionCheckpointMinTimeMsecs                                         int `json:"transaction checkpoint min time (msecs"`
	TransactionCheckpointMostRecentTimeMsecs                                  int `json:"transaction checkpoint most recent time (msecs"`
	TransactionCheckpointScrubDirtyTarget                                     int `json:"transaction checkpoint scrub dirty target"`
	TransactionCheckpointScrubTimeMsecs                                       int `json:"transaction checkpoint scrub time (msecs"`
	TransactionCheckpointTotalTimeMsecs                                       int `json:"transaction checkpoint total time (msecs"`
	TransactionCheckpoints                                                    int `json:"transaction checkpoints"`
	TransactionFailuresDueToCacheOverflow                                     int `json:"transaction failures due to cache overflow"`
	TransactionFsyncCallsForCheckpointAfterAllocatingTheTransactionID         int `json:"transaction fsync calls for checkpoint after allocating the transaction ID"`
	TransactionFsyncDurationForCheckpointAfterAllocatingTheTransactionIDUsecs int `json:"transaction fsync duration for checkpoint after allocating the transaction ID (usecs"`
	TransactionRangeOfIDsCurrentlyPinned                                      int `json:"transaction range of IDs currently pinned"`
	TransactionRangeOfIDsCurrentlyPinnedByACheckpoint                         int `json:"transaction range of IDs currently pinned by a checkpoint"`
	TransactionRangeOfIDsCurrentlyPinnedByNamedSnapshots                      int `json:"transaction range of IDs currently pinned by named snapshots"`
	TransactionSyncCalls                                                      int `json:"transaction sync calls"`
	TransactionsCommitted                                                     int `json:"transactions committed"`
	TransactionsRolledBack                                                    int `json:"transactions rolled back"`
}

// Write contains stats about Write
type Write struct {
	Out          int `json:"out"`
	Available    int `json:"available"`
	TotalTickets int `json:"totalTickets"`
}

// Read contains stats about Read
type Read struct {
	Out          int `json:"out"`
	Available    int `json:"available"`
	TotalTickets int `json:"totalTickets"`
}

// ConcurrentTransactions contains stats about ConcurrentTransactions
type ConcurrentTransactions struct {
	Write *Write `json:"write"`
	Read  *Read  `json:"read"`
}

// WiredTiger contains stats about WiredTiger
type WiredTiger struct {
	URI                    string                  `json:"uri"`
	LSM                    *LSM                    `json:"LSM"`
	Async                  *Async                  `json:"async"`
	BlockManager           *BlockManager           `json:"block-manager"`
	Cache                  *Cache                  `json:"cache"`
	Connection             *Connection             `json:"connection"`
	Cursor                 *WireCursor             `json:"cursor"`
	DataHandle             *DataHandle             `json:"data-handle"`
	Log                    *Log                    `json:"log"`
	Reconciliation         *Reconciliation         `json:"reconciliation"`
	Session                *Session                `json:"session"`
	ThreadState            *ThreadState            `json:"thread-state"`
	ThreadYield            *ThreadYield            `json:"thread-yield"`
	Transaction            *Transaction            `json:"transaction"`
	ConcurrentTransactions *ConcurrentTransactions `json:"concurrentTransactions"`
}

// BuildInfo contains stats about BuildInfo
type BuildInfo struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}

// GetLog contains stats about GetLog
type GetLog struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}

// IsMaster contains stats about IsMaster
type IsMaster struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}

// ReplSetGetStatus contains stats about ReplSetGetStatus
type ReplSetGetStatus struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}

// ServerStatusR contains stats about ServerStatusR
type ServerStatusR struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}

// Whatsmyuri contains stats about Whatsmyuri
type Whatsmyuri struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}

// Commands contains stats about Commands
type Commands struct {
	BuildInfo        *BuildInfo        `json:"buildInfo"`
	GetLog           *GetLog           `json:"getLog"`
	IsMaster         *IsMaster         `json:"isMaster"`
	ReplSetGetStatus *ReplSetGetStatus `json:"replSetGetStatus"`
	ServerStatus     *ServerStatusR    `json:"serverStatus"`
	Whatsmyuri       *Whatsmyuri       `json:"whatsmyuri"`
}

// Cursor contains stats about Cursor
type Cursor struct {
	TimedOut int   `json:"timedOut"`
	Open     *Open `json:"open"`
}

// Open contains stats about Open
type Open struct {
	NoTimeout int `json:"noTimeout"`
	Pinned    int `json:"pinned"`
	Total     int `json:"total"`
}

// Document contains stats about Document
type Document struct {
	Deleted  int `json:"deleted"`
	Inserted int `json:"inserted"`
	Returned int `json:"returned"`
	Updated  int `json:"updated"`
}

// Wtime contains stats about Wtime
type Wtime struct {
	Num         int `json:"num"`
	TotalMillis int `json:"totalMillis"`
}

// GetLastError contains stats about GetLastError
type GetLastError struct {
	Wtime     *Wtime `json:"wtime"`
	Wtimeouts int    `json:"wtimeouts"`
}

// Operation contains stats about Operation
type Operation struct {
	Fastmod        int `json:"fastmod"`
	Idhack         int `json:"idhack"`
	ScanAndOrder   int `json:"scanAndOrder"`
	WriteConflicts int `json:"writeConflicts"`
}

// QueryExecutor contains stats about QueryExecutor
type QueryExecutor struct {
	Scanned        int `json:"scanned"`
	ScannedObjects int `json:"scannedObjects"`
}

// Record contains stats about Record
type Record struct {
	Moves int `json:"moves"`
}

// Counters contains stats about Counters
type Counters struct {
	EventCreated       int `json:"eventCreated"`
	EventWait          int `json:"eventWait"`
	Cancels            int `json:"cancels"`
	Waits              int `json:"waits"`
	ScheduledNetCmd    int `json:"scheduledNetCmd"`
	ScheduledDBWork    int `json:"scheduledDBWork"`
	ScheduledXclWork   int `json:"scheduledXclWork"`
	ScheduledWorkAt    int `json:"scheduledWorkAt"`
	ScheduledWork      int `json:"scheduledWork"`
	SchedulingFailures int `json:"schedulingFailures"`
}

// Queues contains stats about Queues
type Queues struct {
	NetworkInProgress   int `json:"networkInProgress"`
	DbWorkInProgress    int `json:"dbWorkInProgress"`
	ExclusiveInProgress int `json:"exclusiveInProgress"`
	Sleepers            int `json:"sleepers"`
	Ready               int `json:"ready"`
	Free                int `json:"free"`
}

// Executor contains stats about Executor
type Executor struct {
	Counters         *Counters `json:"counters"`
	Queues           *Queues   `json:"queues"`
	UnsignaledEvents int       `json:"unsignaledEvents"`
	EventWaiters     int       `json:"eventWaiters"`
	ShuttingDown     bool      `json:"shuttingDown"`
	NetworkInterface string    `json:"networkInterface"`
}

// Batches contains stats about Batches
type Batches struct {
	Num         int `json:"num"`
	TotalMillis int `json:"totalMillis"`
}

// Apply contains stats about Apply
type Apply struct {
	Batches *Batches `json:"batches"`
	Ops     int      `json:"ops"`
}

// Buffer contains stats about Buffer
type Buffer struct {
	Count        int `json:"count"`
	MaxSizeBytes int `json:"maxSizeBytes"`
	SizeBytes    int `json:"sizeBytes"`
}

// Getmores contains stats about Getmores
type Getmores struct {
	Num         int `json:"num"`
	TotalMillis int `json:"totalMillis"`
}

// ReplNetwork contains stats about ReplNetwork
type ReplNetwork struct {
	Bytes          int       `json:"bytes"`
	Getmores       *Getmores `json:"getmores"`
	Ops            int       `json:"ops"`
	ReadersCreated int       `json:"readersCreated"`
}

// Docs contains stats about Docs
type Docs struct {
	Num         int `json:"num"`
	TotalMillis int `json:"totalMillis"`
}

// Indexes contains stats about Indexes
type Indexes struct {
	Num         int `json:"num"`
	TotalMillis int `json:"totalMillis"`
}

// Preload contains stats about Preload
type Preload struct {
	Docs    *Docs    `json:"docs"`
	Indexes *Indexes `json:"indexes"`
}

// Repl contains stats about Repl
type Repl struct {
	Executor *Executor    `json:"executor"`
	Apply    *Apply       `json:"apply"`
	Buffer   *Buffer      `json:"buffer"`
	Network  *ReplNetwork `json:"network"`
	Preload  *Preload     `json:"preload"`
}

// Search contains stats about Search
type Search struct {
	BucketExhausted int `json:"bucketExhausted"`
	Requests        int `json:"requests"`
	Scanned         int `json:"scanned"`
}

// Freelist contains stats about Freelist
type Freelist struct {
	Search *Search `json:"search"`
}

// Storage contains stats about Storage
type Storage struct {
	Freelist *Freelist `json:"freelist"`
}

// TTL contains stats about TTL
type TTL struct {
	DeletedDocuments int `json:"deletedDocuments"`
	Passes           int `json:"passes"`
}

// Metrics contains stats about Metrics
type Metrics struct {
	Commands      *Commands      `json:"commands"`
	Cursor        *Cursor        `json:"cursor"`
	Document      *Document      `json:"document"`
	GetLastError  *GetLastError  `json:"getLastError"`
	Operation     *Operation     `json:"operation"`
	QueryExecutor *QueryExecutor `json:"queryExecutor"`
	Record        *Record        `json:"record"`
	Repl          *Repl          `json:"repl"`
	Storage       *Storage       `json:"storage"`
	TTL           *TTL           `json:"ttl"`
}

// ServerStatus contains stats about database
type ServerStatus struct {
	Host              string          `json:"host"`
	AdvisoryHostFQDNs []interface{}   `json:"advisoryHostFQDNs"`
	Version           string          `json:"version"`
	Process           string          `json:"process"`
	Pid               int             `json:"pid"`
	Uptime            int             `json:"uptime"`
	UptimeMillis      int             `json:"uptimeMillis"`
	UptimeEstimate    int             `json:"uptimeEstimate"`
	LocalTime         time.Time       `json:"localTime"`
	Asserts           *Asserts        `json:"asserts"`
	Connections       *Connections    `json:"connections"`
	Extra_Info        *Extra_Info     `json:"extra_info"`
	GlobalLock        *GlobalLock     `json:"globalLock"`
	Locks             *Locks          `json:"locks"`
	Network           *Network        `json:"network"`
	Opcounters        *Opcounters     `json:"opcounters"`
	OpcountersRepl    *OpcountersRepl `json:"opcountersRepl"`
	StorageEngine     *StorageEngine  `json:"storageEngine"`
	Tcmalloc          *Tcmalloc       `json:"tcmalloc"`
	WiredTiger        *WiredTiger     `json:"wiredTiger"`
	WriteBacksQueued  bool            `json:"writeBacksQueued"`
	Mem               *Mem            `json:"mem"`
	Metrics           *Metrics        `json:"metrics"`
	Ok                int             `json:"ok"`
}
