package memtable

import (
   "github.com/slowmoon/leveldb/internal"
   "github.com/slowmoon/leveldb/skiplist"
)


type MemTable struct {
   sk *skiplist.SkipList
   memoryUsage uint64
}

func New()*MemTable  {

   table := new(MemTable)
   table.sk =  skiplist.New(internal.InternalKeyComparator, 4)

   return  table
}

func (m *MemTable)Add(seq uint64, vt internal.ValueType, key, value []byte)  {
      internalKey := internal.NewInternalKey(seq, vt, key, value)
      m.memoryUsage += uint64(17+ len(internalKey.UserKey) + len(internalKey.UserValue))
      m.sk.Insert( internalKey, nil)
}

func (m *MemTable)ApproximateMemoryUsage()uint64  {
    return m.memoryUsage
}

func (m *MemTable)Get(key []byte) ([]byte, error) {
      internalKey := internal.LookupKey(key)
      iter := m.sk.Iterator()
      iter.Seek(key)
      if iter.Valid() {
          lookupkey := iter.Key().(*internal.InternalKey)
          if internal.InternalKeyComparator.Compare(internalKey, lookupkey) ==0 {
                 if lookupkey.Type == internal.TypeValue {
                     return lookupkey.UserValue, nil
                 } else  {
                     return  nil, internal.ErrDeletion
                 }
           }
      }
      return  nil, internal.ErrNotFound
}