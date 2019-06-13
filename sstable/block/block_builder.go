package block

import (
	"bytes"
	"encoding/binary"
	"github.com/slowmoon/leveldb/internal"
)

type BlockBuilder struct {
	 buf bytes.Buffer
	 count  uint32
}

func (b *BlockBuilder)Reset()  {
      b.count = 0
      b.buf.Reset()
}

func (b *BlockBuilder)Add(key *internal.InternalKey)  {

     b.count++
     key.EncodeTo(&b.buf)
}

func (b *BlockBuilder)Finish()[]byte  {
	 defer b.Reset()

     binary.Write( &b.buf, binary.BigEndian, b.count)

     return  b.buf.Bytes()

}

func (b *BlockBuilder)Empty()bool  {

     return  b.buf.Len() == 0
}

func (b  *BlockBuilder)CurrentSizeEstimate() int {

    return  b.buf.Len()
}
