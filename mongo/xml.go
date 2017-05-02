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
	Regular   int `bson:"regular"`
	Warning   int `bson:"warning"`
	Msg       int `bson:"msg"`
	User      int `bson:"user"`
	Rollovers int `bson:"rollovers"`
}

//Connections contains info about Connections
type Connections struct {
	Current      int `bson:"current"`
	Available    int `bson:"available"`
	TotalCreated int `bson:"totalCreated"`
}

//Extra_Info contains info about ExtraInfo
type Extra_Info struct {
	Note           string `bson:"note"`
	HeapUsageBytes int    `bson:"heap_usage_bytes"`
	PageFaults     int    `bson:"page_faults"`
}

//GlobalLock contains info about GlobalLock
type GlobalLock struct {
	TotalTime     int            `bson:"totalTime"`
	CurrentQueue  *CurrentQueue  `bson:"currentQueue"`
	ActiveClients *ActiveClients `bson:"activeClients"`
}

//ActiveClients contains info about ActiveClients
type ActiveClients struct {
	Total   int `bson:"total"`
	Readers int `bson:"readers"`
	Writers int `bson:"writers"`
}

//CurrentQueue contains info about CurrentQueue
type CurrentQueue struct {
	Total   int `bson:"total"`
	Readers int `bson:"readers"`
	Writers int `bson:"writers"`
}

//Opcounters contains info about Opcounters
type Opcounters struct {
	Insert  int `bson:"insert"`
	Query   int `bson:"query"`
	Update  int `bson:"update"`
	Delete  int `bson:"delete"`
	Getmore int `bson:"getmore"`
	Command int `bson:"command"`
}

// AcquireCount contains stats about AcquireCount
type AcquireCount struct {
	R int `bson:"r"`
	W int `bson:"w"`
}

// Global contains stats about Global
type Global struct {
	AcquireCount *AcquireCount `bson:"acquireCount"`
}

// Database contains stats about Database
type Database struct {
	AcquireCount *AcquireCount `bson:"acquireCount"`
}

// Collection contains stats about Collection
type Collection struct {
	AcquireCount *AcquireCount `bson:"acquireCount"`
}

// Metadata contains stats about Metadata
type Metadata struct {
	AcquireCount *AcquireCount `bson:"acquireCount"`
}

// Locks contains stats about Locks
type Locks struct {
	Global     *Global     `bson:"Global"`
	Database   *Database   `bson:"Database"`
	Collection *Collection `bson:"Collection"`
	Metadata   *Metadata   `bson:"Metadata"`
}

// OpcountersRepl contains stats about OpcountersRepl
type OpcountersRepl struct {
	Insert  int `bson:"insert"`
	Query   int `bson:"query"`
	Update  int `bson:"update"`
	Delete  int `bson:"delete"`
	Getmore int `bson:"getmore"`
	Command int `bson:"command"`
}

// Network contains stats about Network
type Network struct {
	BytesIn     int `bson:"bytesIn"`
	BytesOut    int `bson:"bytesOut"`
	NumRequests int `bson:"numRequests"`
}

// StorageEngine contains stats about StorageEngine
type StorageEngine struct {
	Name                   string `bson:"name"`
	SupportsCommittedReads bool   `bson:"supportsCommittedReads"`
	Persistent             bool   `bson:"persistent"`
}

// Generic contains stats about Generic
type Generic struct {
	CurrentAllocatedBytes int `bson:"current_allocated_bytes"`
	HeapSize              int `bson:"heap_size"`
}

// TcmallocStat contains stats about TcmallocStat
type TcmallocStat struct {
	PageheapFreeBytes            int    `bson:"pageheap_free_bytes"`
	PageheapUnmappedBytes        int    `bson:"pageheap_unmapped_bytes"`
	MaxTotalThreadCacheBytes     int    `bson:"max_total_thread_cache_bytes"`
	CurrentTotalThreadCacheBytes int    `bson:"current_total_thread_cache_bytes"`
	TotalFreeBytes               int    `bson:"total_free_bytes"`
	CentralCacheFreeBytes        int    `bson:"central_cache_free_bytes"`
	TransferCacheFreeBytes       int    `bson:"transfer_cache_free_bytes"`
	ThreadCacheFreeBytes         int    `bson:"thread_cache_free_bytes"`
	AggressiveMemoryDecommit     int    `bson:"aggressive_memory_decommit"`
	FormattedString              string `bson:"formattedString"`
}

// Tcmalloc contains stats about Tcmalloc
type Tcmalloc struct {
	Generic  *Generic      `bson:"generic"`
	Tcmalloc *TcmallocStat `bson:"tcmalloc"`
}

// Mem contains stats about Mem
type Mem struct {
	Bits              int  `bson:"bits"`
	Resident          int  `bson:"resident"`
	Virtual           int  `bson:"virtual"`
	Supported         bool `bson:"supported"`
	Mapped            int  `bson:"mapped"`
	MappedWithJournal int  `bson:"mappedWithJournal"`
}

// LSM contains stats about LSM
type LSM struct {
	ApplicationWorkUnitsCurrentlyQueued int `bson:"application work units currently queued"`
	MergeWorkUnitsCurrentlyQueued       int `bson:"merge work units currently queued"`
	RowsMergedInAnLSMTree               int `bson:"rows merged in an LSM tree"`
	SleepForLSMCheckpointThrottle       int `bson:"sleep for LSM checkpoint throttle"`
	SleepForLSMMergeThrottle            int `bson:"sleep for LSM merge throttle"`
	SwitchWorkUnitsCurrentlyQueued      int `bson:"switch work units currently queued"`
	TreeMaintenanceOperationsDiscarded  int `bson:"tree maintenance operations discarded"`
	TreeMaintenanceOperationsExecuted   int `bson:"tree maintenance operations executed"`
	TreeMaintenanceOperationsScheduled  int `bson:"tree maintenance operations scheduled"`
	TreeQueueHitMaximum                 int `bson:"tree queue hit maximum"`
}

// Async contains stats about Async
type Async struct {
	CurrentWorkQueueLength                    int `bson:"current work queue length"`
	MaximumWorkQueueLength                    int `bson:"maximum work queue length"`
	NumberOfAllocationStateRaces              int `bson:"number of allocation state races"`
	NumberOfFlushCalls                        int `bson:"number of flush calls"`
	NumberOfOperationSlotsViewedForAllocation int `bson:"number of operation slots viewed for allocation"`
	NumberOfTimesOperationAllocationFailed    int `bson:"number of times operation allocation failed"`
	NumberOfTimesWorkerFoundNoWork            int `bson:"number of times worker found no work"`
	TotalAllocations                          int `bson:"total allocations"`
	TotalCompactCalls                         int `bson:"total compact calls"`
	TotalInsertCalls                          int `bson:"total insert calls"`
	TotalRemoveCalls                          int `bson:"total remove calls"`
	TotalSearchCalls                          int `bson:"total search calls"`
	TotalUpdateCalls                          int `bson:"total update calls"`
}

// BlockManager contains stats about BlockManager
type BlockManager struct {
	BlocksPreLoaded           int `bson:"blocks pre-loaded"`
	BlocksRead                int `bson:"blocks read"`
	BlocksWritten             int `bson:"blocks written"`
	BytesRead                 int `bson:"bytes read"`
	BytesWritten              int `bson:"bytes written"`
	BytesWrittenForCheckpoint int `bson:"bytes written for checkpoint"`
	MappedBlocksRead          int `bson:"mapped blocks read"`
	MappedBytesRead           int `bson:"mapped bytes read"`
}

// Cache contains stats about Cache
type Cache struct {
	BytesBelongingToPageImagesInTheCache                       int `bson:"bytes belonging to page images in the cache"`
	BytesCurrentlyInTheCache                                   int `bson:"bytes currently in the cache"`
	BytesNotBelongingToPageImagesInTheCache                    int `bson:"bytes not belonging to page images in the cache"`
	BytesReadIntoCache                                         int `bson:"bytes read into cache"`
	BytesWrittenFromCache                                      int `bson:"bytes written from cache"`
	CheckpointBlockedPageEviction                              int `bson:"checkpoint blocked page eviction"`
	EvictionCallsToGetAPage                                    int `bson:"eviction calls to get a page"`
	EvictionCallsToGetAPageFoundQueueEmpty                     int `bson:"eviction calls to get a page found queue empty"`
	EvictionCallsToGetAPageFoundQueueEmptyAfterLocking         int `bson:"eviction calls to get a page found queue empty after locking"`
	EvictionCurrentlyOperatingInAggressiveMode                 int `bson:"eviction currently operating in aggressive mode"`
	EvictionEmptyScore                                         int `bson:"eviction empty score"`
	EvictionServerCandidateQueueEmptyWhenToppingUp             int `bson:"eviction server candidate queue empty when topping up"`
	EvictionServerCandidateQueueNotEmptyWhenToppingUp          int `bson:"eviction server candidate queue not empty when topping up"`
	EvictionServerEvictingPages                                int `bson:"eviction server evicting pages"`
	EvictionServerSleptBecauseWeDidNotMakeProgressWithEviction int `bson:"eviction server slept, because we did not make progress with eviction"`
	EvictionServerUnableToReachEvictionGoal                    int `bson:"eviction server unable to reach eviction goal"`
	EvictionState                                              int `bson:"eviction state"`
	EvictionWalksAbandoned                                     int `bson:"eviction walks abandoned"`
	EvictionWorkerThreadEvictingPages                          int `bson:"eviction worker thread evicting pages"`
	FailedEvictionOfPagesThatExceededTheInMemoryMaximum        int `bson:"failed eviction of pages that exceeded the in-memory maximum"`
	FilesWithActiveEvictionWalks                               int `bson:"files with active eviction walks"`
	FilesWithNewEvictionWalksStarted                           int `bson:"files with new eviction walks started"`
	HazardPointerBlockedPageEviction                           int `bson:"hazard pointer blocked page eviction"`
	HazardPointerCheckCalls                                    int `bson:"hazard pointer check calls"`
	HazardPointerCheckEntriesWalked                            int `bson:"hazard pointer check entries walked"`
	HazardPointerMaximumArrayLength                            int `bson:"hazard pointer maximum array length"`
	InMemoryPagePassedCriteriaToBeSplit                        int `bson:"in-memory page passed criteria to be split"`
	InMemoryPageSplits                                         int `bson:"in-memory page splits"`
	InternalPagesEvicted                                       int `bson:"internal pages evicted"`
	InternalPagesSplitDuringEviction                           int `bson:"internal pages split during eviction"`
	LeafPagesSplitDuringEviction                               int `bson:"leaf pages split during eviction"`
	LookasideTableInsertCalls                                  int `bson:"lookaside table insert calls"`
	LookasideTableRemoveCalls                                  int `bson:"lookaside table remove calls"`
	MaximumBytesConfigured                                     int `bson:"maximum bytes configured"`
	MaximumPageSizeAtEviction                                  int `bson:"maximum page size at eviction"`
	ModifiedPagesEvicted                                       int `bson:"modified pages evicted"`
	ModifiedPagesEvictedByApplicationThreads                   int `bson:"modified pages evicted by application threads"`
	OverflowPagesReadIntoCache                                 int `bson:"overflow pages read into cache"`
	OverflowValuesCachedInMemory                               int `bson:"overflow values cached in memory"`
	PageSplitDuringEvictionDeepenedTheTree                     int `bson:"page split during eviction deepened the tree"`
	PageWrittenRequiringLookasideRecords                       int `bson:"page written requiring lookaside records"`
	PagesCurrentlyHeldInTheCache                               int `bson:"pages currently held in the cache"`
	PagesEvictedBecauseTheyExceededTheInMemoryMaximum          int `bson:"pages evicted because they exceeded the in-memory maximum"`
	PagesEvictedBecauseTheyHadChainsOfDeletedItems             int `bson:"pages evicted because they had chains of deleted items"`
	PagesEvictedByApplicationThreads                           int `bson:"pages evicted by application threads"`
	PagesQueuedForEviction                                     int `bson:"pages queued for eviction"`
	PagesQueuedForUrgentEviction                               int `bson:"pages queued for urgent eviction"`
	PagesQueuedForUrgentEvictionDuringWalk                     int `bson:"pages queued for urgent eviction during walk"`
	PagesReadIntoCache                                         int `bson:"pages read into cache"`
	PagesReadIntoCacheRequiringLookasideEntries                int `bson:"pages read into cache requiring lookaside entries"`
	PagesRequestedFromTheCache                                 int `bson:"pages requested from the cache"`
	PagesSeenByEvictionWalk                                    int `bson:"pages seen by eviction walk"`
	PagesSelectedForEvictionUnableToBeEvicted                  int `bson:"pages selected for eviction unable to be evicted"`
	PagesWalkedForEviction                                     int `bson:"pages walked for eviction"`
	PagesWrittenFromCache                                      int `bson:"pages written from cache"`
	PagesWrittenRequiringInMemoryRestoration                   int `bson:"pages written requiring in-memory restoration"`
	PercentageOverhead                                         int `bson:"percentage overhead"`
	TrackedBytesBelongingToInternalPagesInTheCache             int `bson:"tracked bytes belonging to internal pages in the cache"`
	TrackedBytesBelongingToLeafPagesInTheCache                 int `bson:"tracked bytes belonging to leaf pages in the cache"`
	TrackedDirtyBytesInTheCache                                int `bson:"tracked dirty bytes in the cache"`
	TrackedDirtyPagesInTheCache                                int `bson:"tracked dirty pages in the cache"`
	UnmodifiedPagesEvicted                                     int `bson:"unmodified pages evicted"`
}

// Connection contains stats about Connection
type Connection struct {
	AutoAdjustingConditionResets         int `bson:"auto adjusting condition resets"`
	AutoAdjustingConditionWaitCalls      int `bson:"auto adjusting condition wait calls"`
	FilesCurrentlyOpen                   int `bson:"files currently open"`
	MemoryAllocations                    int `bson:"memory allocations"`
	MemoryFrees                          int `bson:"memory frees"`
	MemoryReAllocations                  int `bson:"memory re-allocations"`
	PthreadMutexConditionWaitCalls       int `bson:"pthread mutex condition wait calls"`
	PthreadMutexSharedLockReadLockCalls  int `bson:"pthread mutex shared lock read-lock calls"`
	PthreadMutexSharedLockWriteLockCalls int `bson:"pthread mutex shared lock write-lock calls"`
	TotalFsyncIOs                        int `bson:"total fsync I/Os"`
	TotalReadIOs                         int `bson:"total read I/Os"`
	TotalWriteIOs                        int `bson:"total write I/Os"`
}

// WireCursor contains stats about WireCursor
type WireCursor struct {
	CursorCreateCalls       int `bson:"cursor create calls"`
	CursorInsertCalls       int `bson:"cursor insert calls"`
	CursorNextCalls         int `bson:"cursor next calls"`
	CursorPrevCalls         int `bson:"cursor prev calls"`
	CursorRemoveCalls       int `bson:"cursor remove calls"`
	CursorResetCalls        int `bson:"cursor reset calls"`
	CursorRestartedSearches int `bson:"cursor restarted searches"`
	CursorSearchCalls       int `bson:"cursor search calls"`
	CursorSearchNearCalls   int `bson:"cursor search near calls"`
	CursorUpdateCalls       int `bson:"cursor update calls"`
	TruncateCalls           int `bson:"truncate calls"`
}

// DataHandle contains stats about DataHandle
type DataHandle struct {
	ConnectionDataHandlesCurrentlyActive       int `bson:"connection data handles currently active"`
	ConnectionSweepCandidateBecameReferenced   int `bson:"connection sweep candidate became referenced"`
	ConnectionSweepDhandlesClosed              int `bson:"connection sweep dhandles closed"`
	ConnectionSweepDhandlesRemovedFromHashList int `bson:"connection sweep dhandles removed from hash list"`
	ConnectionSweepTimeOfDeathSets             int `bson:"connection sweep time-of-death sets"`
	ConnectionSweeps                           int `bson:"connection sweeps"`
	SessionDhandlesSwept                       int `bson:"session dhandles swept"`
	SessionSweepAttempts                       int `bson:"session sweep attempts"`
}

// Log contains stats about Log
type Log struct {
	BusyReturnsAttemptingToSwitchSlots    int `bson:"busy returns attempting to switch slots"`
	ConsolidatedSlotClosures              int `bson:"consolidated slot closures"`
	ConsolidatedSlotJoinRaces             int `bson:"consolidated slot join races"`
	ConsolidatedSlotJoinTransitions       int `bson:"consolidated slot join transitions"`
	ConsolidatedSlotJoins                 int `bson:"consolidated slot joins"`
	ConsolidatedSlotUnbufferedWrites      int `bson:"consolidated slot unbuffered writes"`
	LogBytesOfPayloadData                 int `bson:"log bytes of payload data"`
	LogBytesWritten                       int `bson:"log bytes written"`
	LogFilesManuallyZeroFilled            int `bson:"log files manually zero-filled"`
	LogFlushOperations                    int `bson:"log flush operations"`
	LogForceWriteOperations               int `bson:"log force write operations"`
	LogForceWriteOperationsSkipped        int `bson:"log force write operations skipped"`
	LogRecordsCompressed                  int `bson:"log records compressed"`
	LogRecordsNotCompressed               int `bson:"log records not compressed"`
	LogRecordsTooSmallToCompress          int `bson:"log records too small to compress"`
	LogReleaseAdvancesWriteLSN            int `bson:"log release advances write LSN"`
	LogScanOperations                     int `bson:"log scan operations"`
	LogScanRecordsRequiringTwoReads       int `bson:"log scan records requiring two reads"`
	LogServerThreadAdvancesWriteLSN       int `bson:"log server thread advances write LSN"`
	LogServerThreadWriteLSNWalkSkipped    int `bson:"log server thread write LSN walk skipped"`
	LogSyncOperations                     int `bson:"log sync operations"`
	LogSyncTimeDurationUsecs              int `bson:"log sync time duration (usecs"`
	LogSyncDirOperations                  int `bson:"log sync_dir operations"`
	LogSyncDirTimeDurationUsecs           int `bson:"log sync_dir time duration (usecs"`
	LogWriteOperations                    int `bson:"log write operations"`
	LoggingBytesConsolidated              int `bson:"logging bytes consolidated"`
	MaximumLogFileSize                    int `bson:"maximum log file size"`
	NumberOfPreAllocatedLogFilesToCreate  int `bson:"number of pre-allocated log files to create"`
	PreAllocatedLogFilesNotReadyAndMissed int `bson:"pre-allocated log files not ready and missed"`
	PreAllocatedLogFilesPrepared          int `bson:"pre-allocated log files prepared"`
	PreAllocatedLogFilesUsed              int `bson:"pre-allocated log files used"`
	RecordsProcessedByLogScan             int `bson:"records processed by log scan"`
	TotalInMemorySizeOfCompressedRecords  int `bson:"total in-memory size of compressed records"`
	TotalLogBufferSize                    int `bson:"total log buffer size"`
	TotalSizeOfCompressedRecords          int `bson:"total size of compressed records"`
	WrittenSlotsCoalesced                 int `bson:"written slots coalesced"`
	YieldsWaitingForPreviousLogFileClose  int `bson:"yields waiting for previous log file close"`
}

// Reconciliation contains stats about Reconciliation
type Reconciliation struct {
	FastPathPagesDeleted               int `bson:"fast-path pages deleted"`
	PageReconciliationCalls            int `bson:"page reconciliation calls"`
	PageReconciliationCallsForEviction int `bson:"page reconciliation calls for eviction"`
	PagesDeleted                       int `bson:"pages deleted"`
	SplitBytesCurrentlyAwaitingFree    int `bson:"split bytes currently awaiting free"`
	SplitObjectsCurrentlyAwaitingFree  int `bson:"split objects currently awaiting free"`
}

// Session contains stats about Session
type Session struct {
	OpenCursorCount               int `bson:"open cursor count"`
	OpenSessionCount              int `bson:"open session count"`
	TableCompactFailedCalls       int `bson:"table compact failed calls"`
	TableCompactSuccessfulCalls   int `bson:"table compact successful calls"`
	TableCreateFailedCalls        int `bson:"table create failed calls"`
	TableCreateSuccessfulCalls    int `bson:"table create successful calls"`
	TableDropFailedCalls          int `bson:"table drop failed calls"`
	TableDropSuccessfulCalls      int `bson:"table drop successful calls"`
	TableRebalanceFailedCalls     int `bson:"table rebalance failed calls"`
	TableRebalanceSuccessfulCalls int `bson:"table rebalance successful calls"`
	TableRenameFailedCalls        int `bson:"table rename failed calls"`
	TableRenameSuccessfulCalls    int `bson:"table rename successful calls"`
	TableSalvageFailedCalls       int `bson:"table salvage failed calls"`
	TableSalvageSuccessfulCalls   int `bson:"table salvage successful calls"`
	TableTruncateFailedCalls      int `bson:"table truncate failed calls"`
	TableTruncateSuccessfulCalls  int `bson:"table truncate successful calls"`
	TableVerifyFailedCalls        int `bson:"table verify failed calls"`
	TableVerifySuccessfulCalls    int `bson:"table verify successful calls"`
}

// ThreadState contains stats about ThreadState
type ThreadState struct {
	ActiveFilesystemFsyncCalls int `bson:"active filesystem fsync calls"`
	ActiveFilesystemReadCalls  int `bson:"active filesystem read calls"`
	ActiveFilesystemWriteCalls int `bson:"active filesystem write calls"`
}

// ThreadYield contains stats about ThreadYield
type ThreadYield struct {
	PageAcquireBusyBlocked       int `bson:"page acquire busy blocked"`
	PageAcquireEvictionBlocked   int `bson:"page acquire eviction blocked"`
	PageAcquireLockedBlocked     int `bson:"page acquire locked blocked"`
	PageAcquireReadBlocked       int `bson:"page acquire read blocked"`
	PageAcquireTimeSleepingUsecs int `bson:"page acquire time sleeping (usecs"`
}

// Transaction contains stats about Transaction
type Transaction struct {
	NumberOfNamedSnapshotsCreated                                             int `bson:"number of named snapshots created"`
	NumberOfNamedSnapshotsDropped                                             int `bson:"number of named snapshots dropped"`
	TransactionBegins                                                         int `bson:"transaction begins"`
	TransactionCheckpointCurrentlyRunning                                     int `bson:"transaction checkpoint currently running"`
	TransactionCheckpointGeneration                                           int `bson:"transaction checkpoint generation"`
	TransactionCheckpointMaxTimeMsecs                                         int `bson:"transaction checkpoint max time (msecs"`
	TransactionCheckpointMinTimeMsecs                                         int `bson:"transaction checkpoint min time (msecs"`
	TransactionCheckpointMostRecentTimeMsecs                                  int `bson:"transaction checkpoint most recent time (msecs"`
	TransactionCheckpointScrubDirtyTarget                                     int `bson:"transaction checkpoint scrub dirty target"`
	TransactionCheckpointScrubTimeMsecs                                       int `bson:"transaction checkpoint scrub time (msecs"`
	TransactionCheckpointTotalTimeMsecs                                       int `bson:"transaction checkpoint total time (msecs"`
	TransactionCheckpoints                                                    int `bson:"transaction checkpoints"`
	TransactionFailuresDueToCacheOverflow                                     int `bson:"transaction failures due to cache overflow"`
	TransactionFsyncCallsForCheckpointAfterAllocatingTheTransactionID         int `bson:"transaction fsync calls for checkpoint after allocating the transaction ID"`
	TransactionFsyncDurationForCheckpointAfterAllocatingTheTransactionIDUsecs int `bson:"transaction fsync duration for checkpoint after allocating the transaction ID (usecs"`
	TransactionRangeOfIDsCurrentlyPinned                                      int `bson:"transaction range of IDs currently pinned"`
	TransactionRangeOfIDsCurrentlyPinnedByACheckpoint                         int `bson:"transaction range of IDs currently pinned by a checkpoint"`
	TransactionRangeOfIDsCurrentlyPinnedByNamedSnapshots                      int `bson:"transaction range of IDs currently pinned by named snapshots"`
	TransactionSyncCalls                                                      int `bson:"transaction sync calls"`
	TransactionsCommitted                                                     int `bson:"transactions committed"`
	TransactionsRolledBack                                                    int `bson:"transactions rolled back"`
}

// Write contains stats about Write
type Write struct {
	Out          int `bson:"out"`
	Available    int `bson:"available"`
	TotalTickets int `bson:"totalTickets"`
}

// Read contains stats about Read
type Read struct {
	Out          int `bson:"out"`
	Available    int `bson:"available"`
	TotalTickets int `bson:"totalTickets"`
}

// ConcurrentTransactions contains stats about ConcurrentTransactions
type ConcurrentTransactions struct {
	Write *Write `bson:"write"`
	Read  *Read  `bson:"read"`
}

// WiredTiger contains stats about WiredTiger
type WiredTiger struct {
	URI                    string                  `bson:"uri"`
	LSM                    *LSM                    `bson:"LSM"`
	Async                  *Async                  `bson:"async"`
	BlockManager           *BlockManager           `bson:"block-manager"`
	Cache                  *Cache                  `bson:"cache"`
	Connection             *Connection             `bson:"connection"`
	Cursor                 *WireCursor             `bson:"cursor"`
	DataHandle             *DataHandle             `bson:"data-handle"`
	Log                    *Log                    `bson:"log"`
	Reconciliation         *Reconciliation         `bson:"reconciliation"`
	Session                *Session                `bson:"session"`
	ThreadState            *ThreadState            `bson:"thread-state"`
	ThreadYield            *ThreadYield            `bson:"thread-yield"`
	Transaction            *Transaction            `bson:"transaction"`
	ConcurrentTransactions *ConcurrentTransactions `bson:"concurrentTransactions"`
}

// BuildInfo contains stats about BuildInfo
type BuildInfo struct {
	Failed int `bson:"failed"`
	Total  int `bson:"total"`
}

// GetLog contains stats about GetLog
type GetLog struct {
	Failed int `bson:"failed"`
	Total  int `bson:"total"`
}

// IsMaster contains stats about IsMaster
type IsMaster struct {
	Failed int `bson:"failed"`
	Total  int `bson:"total"`
}

// ReplSetGetStatus contains stats about ReplSetGetStatus
type ReplSetGetStatus struct {
	Failed int `bson:"failed"`
	Total  int `bson:"total"`
}

// ServerStatusR contains stats about ServerStatusR
type ServerStatusR struct {
	Failed int `bson:"failed"`
	Total  int `bson:"total"`
}

// Whatsmyuri contains stats about Whatsmyuri
type Whatsmyuri struct {
	Failed int `bson:"failed"`
	Total  int `bson:"total"`
}

// Commands contains stats about Commands
type Commands struct {
	BuildInfo        *BuildInfo        `bson:"buildInfo"`
	GetLog           *GetLog           `bson:"getLog"`
	IsMaster         *IsMaster         `bson:"isMaster"`
	ReplSetGetStatus *ReplSetGetStatus `bson:"replSetGetStatus"`
	ServerStatus     *ServerStatusR    `bson:"serverStatus"`
	Whatsmyuri       *Whatsmyuri       `bson:"whatsmyuri"`
}

// Cursor contains stats about Cursor
type Cursor struct {
	TimedOut int   `bson:"timedOut"`
	Open     *Open `bson:"open"`
}

// Open contains stats about Open
type Open struct {
	NoTimeout int `bson:"noTimeout"`
	Pinned    int `bson:"pinned"`
	Total     int `bson:"total"`
}

// Document contains stats about Document
type Document struct {
	Deleted  int `bson:"deleted"`
	Inserted int `bson:"inserted"`
	Returned int `bson:"returned"`
	Updated  int `bson:"updated"`
}

// Wtime contains stats about Wtime
type Wtime struct {
	Num         int `bson:"num"`
	TotalMillis int `bson:"totalMillis"`
}

// GetLastError contains stats about GetLastError
type GetLastError struct {
	Wtime     *Wtime `bson:"wtime"`
	Wtimeouts int    `bson:"wtimeouts"`
}

// Operation contains stats about Operation
type Operation struct {
	Fastmod        int `bson:"fastmod"`
	Idhack         int `bson:"idhack"`
	ScanAndOrder   int `bson:"scanAndOrder"`
	WriteConflicts int `bson:"writeConflicts"`
}

// QueryExecutor contains stats about QueryExecutor
type QueryExecutor struct {
	Scanned        int `bson:"scanned"`
	ScannedObjects int `bson:"scannedObjects"`
}

// Record contains stats about Record
type Record struct {
	Moves int `bson:"moves"`
}

// Counters contains stats about Counters
type Counters struct {
	EventCreated       int `bson:"eventCreated"`
	EventWait          int `bson:"eventWait"`
	Cancels            int `bson:"cancels"`
	Waits              int `bson:"waits"`
	ScheduledNetCmd    int `bson:"scheduledNetCmd"`
	ScheduledDBWork    int `bson:"scheduledDBWork"`
	ScheduledXclWork   int `bson:"scheduledXclWork"`
	ScheduledWorkAt    int `bson:"scheduledWorkAt"`
	ScheduledWork      int `bson:"scheduledWork"`
	SchedulingFailures int `bson:"schedulingFailures"`
}

// Queues contains stats about Queues
type Queues struct {
	NetworkInProgress   int `bson:"networkInProgress"`
	DbWorkInProgress    int `bson:"dbWorkInProgress"`
	ExclusiveInProgress int `bson:"exclusiveInProgress"`
	Sleepers            int `bson:"sleepers"`
	Ready               int `bson:"ready"`
	Free                int `bson:"free"`
}

// Executor contains stats about Executor
type Executor struct {
	Counters         *Counters `bson:"counters"`
	Queues           *Queues   `bson:"queues"`
	UnsignaledEvents int       `bson:"unsignaledEvents"`
	EventWaiters     int       `bson:"eventWaiters"`
	ShuttingDown     bool      `bson:"shuttingDown"`
	NetworkInterface string    `bson:"networkInterface"`
}

// Batches contains stats about Batches
type Batches struct {
	Num         int `bson:"num"`
	TotalMillis int `bson:"totalMillis"`
}

// Apply contains stats about Apply
type Apply struct {
	Batches *Batches `bson:"batches"`
	Ops     int      `bson:"ops"`
}

// Buffer contains stats about Buffer
type Buffer struct {
	Count        int `bson:"count"`
	MaxSizeBytes int `bson:"maxSizeBytes"`
	SizeBytes    int `bson:"sizeBytes"`
}

// Getmores contains stats about Getmores
type Getmores struct {
	Num         int `bson:"num"`
	TotalMillis int `bson:"totalMillis"`
}

// ReplNetwork contains stats about ReplNetwork
type ReplNetwork struct {
	Bytes          int       `bson:"bytes"`
	Getmores       *Getmores `bson:"getmores"`
	Ops            int       `bson:"ops"`
	ReadersCreated int       `bson:"readersCreated"`
}

// Docs contains stats about Docs
type Docs struct {
	Num         int `bson:"num"`
	TotalMillis int `bson:"totalMillis"`
}

// Indexes contains stats about Indexes
type Indexes struct {
	Num         int `bson:"num"`
	TotalMillis int `bson:"totalMillis"`
}

// Preload contains stats about Preload
type Preload struct {
	Docs    *Docs    `bson:"docs"`
	Indexes *Indexes `bson:"indexes"`
}

// Repl contains stats about Repl
type Repl struct {
	Executor *Executor    `bson:"executor"`
	Apply    *Apply       `bson:"apply"`
	Buffer   *Buffer      `bson:"buffer"`
	Network  *ReplNetwork `bson:"network"`
	Preload  *Preload     `bson:"preload"`
}

// Search contains stats about Search
type Search struct {
	BucketExhausted int `bson:"bucketExhausted"`
	Requests        int `bson:"requests"`
	Scanned         int `bson:"scanned"`
}

// Freelist contains stats about Freelist
type Freelist struct {
	Search *Search `bson:"search"`
}

// Storage contains stats about Storage
type Storage struct {
	Freelist *Freelist `bson:"freelist"`
}

// TTL contains stats about TTL
type TTL struct {
	DeletedDocuments int `bson:"deletedDocuments"`
	Passes           int `bson:"passes"`
}

// Metrics contains stats about Metrics
type Metrics struct {
	Commands      *Commands      `bson:"commands"`
	Cursor        *Cursor        `bson:"cursor"`
	Document      *Document      `bson:"document"`
	GetLastError  *GetLastError  `bson:"getLastError"`
	Operation     *Operation     `bson:"operation"`
	QueryExecutor *QueryExecutor `bson:"queryExecutor"`
	Record        *Record        `bson:"record"`
	Repl          *Repl          `bson:"repl"`
	Storage       *Storage       `bson:"storage"`
	TTL           *TTL           `bson:"ttl"`
}

// ServerStatus contains stats about database
type ServerStatus struct {
	Host              string          `bson:"host"`
	AdvisoryHostFQDNs []interface{}   `bson:"advisoryHostFQDNs"`
	Version           string          `bson:"version"`
	Process           string          `bson:"process"`
	Pid               int             `bson:"pid"`
	Uptime            int             `bson:"uptime"`
	UptimeMillis      int             `bson:"uptimeMillis"`
	UptimeEstimate    int             `bson:"uptimeEstimate"`
	LocalTime         time.Time       `bson:"localTime"`
	Asserts           *Asserts        `bson:"asserts"`
	Connections       *Connections    `bson:"connections"`
	Extra_Info        *Extra_Info     `bson:"extra_info"`
	GlobalLock        *GlobalLock     `bson:"globalLock"`
	Locks             *Locks          `bson:"locks"`
	Network           *Network        `bson:"network"`
	Opcounters        *Opcounters     `bson:"opcounters"`
	OpcountersRepl    *OpcountersRepl `bson:"opcountersRepl"`
	StorageEngine     *StorageEngine  `bson:"storageEngine"`
	Tcmalloc          *Tcmalloc       `bson:"tcmalloc"`
	WiredTiger        *WiredTiger     `bson:"wiredTiger"`
	WriteBacksQueued  bool            `bson:"writeBacksQueued"`
	Mem               *Mem            `bson:"mem"`
	Metrics           *Metrics        `bson:"metrics"`
	Ok                int             `bson:"ok"`
}
