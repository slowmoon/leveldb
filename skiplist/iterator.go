package skiplist

type Iterator struct {
	 sk *SkipList
	 node *Node
}

func NewIterator(sk *SkipList) *Iterator {
     return  &Iterator{
     	sk: sk,
	 }
}

func (i *Iterator)Valid()bool  {
   return  i.node != nil
}

func (i *Iterator)Key()interface{}  {
    return  i.node.key
}

func (i *Iterator)Value()interface{}  {
    return  i.node.value
}

func (i *Iterator)Next()  {
    i.sk.lock.RLock()
    defer i.sk.lock.RUnlock()
    i.node = i.node.getNext(0)
}

func (i *Iterator)Prev()  {
   i.sk.lock.RLock()
   defer i.sk.lock.RUnlock()
     i.node , _ = i.sk.findGreaterOrEqual(i.node.key)    //TODO
}

func (i *Iterator)Seek(key []byte)  {
    i.sk.lock.RLock()
    defer i.sk.lock.RUnlock()
    i.node,  _ = i.sk.findGreaterOrEqual(key)
}

func (i *Iterator)SeekToFirst()  {

}

func (i *Iterator)SeekToLast()  {

}
