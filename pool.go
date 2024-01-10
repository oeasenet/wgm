package wgm

import "sync"

// FindPageOptionSyncPool is a sync.Pool that stores FindPageOption objects.
var FindPageOptionSyncPool = sync.Pool{
	New: func() interface{} {
		return new(FindPageOption)
	},
}

// acquireFindPageOption
// Description: acquire a FindPageOption object from the pool.
func acquireFindPageOption() *FindPageOption {
	return FindPageOptionSyncPool.Get().(*FindPageOption)
}

// releaseFindPageOption
// Description: release a FindPageOption object to the pool.
// Param: m
func releaseFindPageOption(m *FindPageOption) {
	m.selector = nil
	m.fields = nil
	FindPageOptionSyncPool.Put(m)
}
