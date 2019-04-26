package search

import "time"

const defaultPageSize = 300
const MaxAllowedSearchHits = 10000

// Only used by RegexpSearcher
const MaxAllowedSearchMessages = 100000

// Paging context around a search hit
const MaxContext = 15

const (
	// max convs to sync in the background
	maxSyncConvsDesktop = 20
	maxSyncConvsMobile  = 5

	// delay before starting SelectiveSync
	startSyncDelayDesktop = 10 * time.Second
	startSyncDelayMobile  = 30 * time.Second
)

// Bumped whenever there are tokenization changes to building the index.
const IndexDataVersion = 5

// Bumped if we change the index key because of a different disk encoding.
type indexDiskVersion int

const latestIndexDiskVersion = indexDiskVersion(3)
