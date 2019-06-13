package internal

import (
    "bytes"
    "encoding/binary"
    "github.com/slowmoon/leveldb/comp"
    "io"
    "math"
)

const (
   TypeDeletion ValueType = iota
   TypeValue
)


type ValueType int8

type InternalKey struct {

   Seq uint64
   Type  ValueType
   UserKey []byte
   UserValue []byte
}

func NewInternalKey(seq uint64, valueType ValueType, key, value []byte) *InternalKey {
    var internal InternalKey
    internal.Seq = seq
    internal.Type = valueType
    internal.UserKey = make([]byte, len(key))
    copy(internal.UserKey, key)
    internal.UserValue = make([]byte, len(value))
    copy(internal.UserValue, value)
    return  &internal
}

func (i *InternalKey)EncodeTo(w io.Writer) error {
	binary.Write(w, binary.BigEndian, i.Seq)
	binary.Write(w, binary.BigEndian, i.Type)
	binary.Write(w, binary.BigEndian, int32(len(i.UserKey)))
    binary.Write(w, binary.BigEndian, i.UserKey)
	binary.Write(w, binary.BigEndian, int32(len(i.UserValue)))
   return binary.Write(w, binary.BigEndian, i.UserValue)
}
func (i *InternalKey)DecodeFrom(r io.Reader) error {
    var temp int32
    binary.Read(r, binary.BigEndian, &i.Seq)
    binary.Read(r, binary.BigEndian, &i.Type)
    binary.Read(r, binary.BigEndian, &temp)
    i.UserKey = make([]byte, temp)
    binary.Read(r, binary.BigEndian, i.UserKey)

    binary.Read(r, binary.BigEndian, &temp)
    i.UserValue = make([]byte, temp)
    return binary.Read(r, binary.BigEndian, i.UserValue)
}

func LookupKey(key []byte) *InternalKey {
    return  NewInternalKey(math.MaxUint64, TypeValue, key, nil)
}


var InternalKeyComparator comp.ComparatorFunc = func(a, b interface{}) int {

    aKey := a.(*InternalKey)
    bKey := b.(*InternalKey)
    res := bytes.Compare(aKey.UserKey, bKey.UserKey)

    switch  {
    case res ==0:
        if aKey.Seq > bKey.Seq {
            //seq bigger and then res is bigger
            res++
        }else {
            res --
        }
    }
    return  res
}

func UserKeyCompare(a, b interface{}) int {

     aKey := a.([]byte)
     bKey := b.([]byte)

     return  bytes.Compare(aKey, bKey)
}