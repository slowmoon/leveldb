package block

import (
	"bytes"
	"encoding/binary"
	"github.com/slowmoon/leveldb/internal"
)

type Block struct {
	items []internal.InternalKey
}


func New(p []byte)*Block {
    var block Block
    buf := bytes.NewBuffer(p)
    counter := binary.BigEndian.Uint32(p[len(p)-4:])

    for i :=0; i < int(counter);i ++ {
        var item internal.InternalKey
        err := item.DecodeFrom(buf)
        if err != nil {
        	return  nil
		}

        block.items = append(block.items, item)
	}
    return  &block
}

func (b *Block)Iterator()*Iterator  {
      return  &Iterator{block: b}
}
