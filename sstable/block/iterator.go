package block

import "github.com/slowmoon/leveldb/internal"

type Iterator struct {
    block *Block
    index int
}

func Iter(b *Block)*Iterator  {
    return  &Iterator{block:b }
}

func (i *Iterator)Valid() bool {
   return  i.index >=0 && i.index < len(i.block.items)
}

func (i *Iterator)Key() interface{} {
    return i.block.items[i.index]
}

func (i *Iterator)Prev()  {
     i.index--
}

func (i *Iterator)Next()  {
     i.index++
}

//seek if true
func (i *Iterator)Seek(key interface{})  {
   left := 0
   right := len(i.block.items) -1
   for left < right {
       mid := (left +right) /2
       if internal.UserKeyCompare(i.block.items[mid].UserKey, key) <0 {
           left = mid+1
       }else {
           right = mid
       }
   }
   if left == len(i.block.items) -1 {
       if internal.UserKeyCompare(i.block.items[left], key) <0 {
           left++
       }
   }
   i.index = left
}

func (i *Iterator)SeekToFirst()  {
    i.index=0
}

func (i *Iterator)SeekToLast()  {
    if i.block.items != nil {
        i.index = len(i.block.items) -1
    }
}