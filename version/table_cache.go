package version

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/slowmoon/leveldb/internal"
	"github.com/slowmoon/leveldb/sstable"
	"sync"
)

type TableCache struct {
	dbName string
	lock sync.Mutex
	cache *lru.Cache
}

func NewTableCache(name string)*TableCache  {
     var tc TableCache
     tc.dbName = name
     tc.cache , _ = lru.New(internal.MaxOpenFiles - internal.NumNonTableCacheFiles)
     return  &tc
}

func (tc *TableCache)findTable(num uint64) (*sstable.SsTable, error) {
     tc.lock.Lock()
     defer tc.lock.Unlock()

     v , ok:= tc.cache.Get(num)
     if ok {
     	return v.(*sstable.SsTable) , nil
	 }

     return  nil , nil
}

func (tc *TableCache)Evict(num uint64)  {
	 tc.cache.Remove(num)
}

func (tc *TableCache)Get(num uint64, key interface{})([]byte, error)  {
	_ ,err := tc.findTable(num)
	if err != nil {
	   return  nil, err
	}
	return nil , nil
}



