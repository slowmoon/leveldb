package sstable

import (
	"encoding/binary"
	"github.com/slowmoon/leveldb/internal"
	"io"
)

const (
	kTableMagicNumber uint64 = 0xdb4775248b80fb57
)


type BlockHandle struct {
     Offset uint32
     Size uint32
}

func (b *BlockHandle)EncodeToBytes()[]byte  {
     bytes := make([]byte, 8)
     binary.BigEndian.PutUint32(bytes, b.Offset)
     binary.BigEndian.PutUint32(bytes[4:], b.Size)
     return  bytes
}

func (b *BlockHandle)DecodeFromBytes(bytes []byte)  {
	_ = bytes[7]     //bounds check
	b.Offset = binary.BigEndian.Uint32(bytes)
	b.Size = binary.BigEndian.Uint32(bytes[4:])
}

type IndexBlockHandler struct {
	*internal.InternalKey
}

func (i *IndexBlockHandler)SetBlockHandler(handler BlockHandle)  {
      i.UserValue = handler.EncodeToBytes()
}

func (i *IndexBlockHandler)GetBlockHandler()(block BlockHandle) {
      block.DecodeFromBytes(i.UserValue)
      return
}

type Footer struct {
	MetaIndexHandler  BlockHandle
	IndexHandler BlockHandle
}

func (f *Footer)Size() int {
   return binary.Size(f) +8
}

func (f *Footer)EncodeTo(w io.Writer) error  {
      err := binary.Write(w, binary.BigEndian, f)
      if err != nil {
      	return  err
	  }

      return binary.Write(w, binary.BigEndian, kTableMagicNumber)
}

func (f *Footer)DecodeFrom(r io.Reader)error  {
	if err := binary.Read(r, binary.BigEndian, f); err != nil {
		return  err
	}
	var num uint64
	err := binary.Read(r, binary.BigEndian, &num)
	if err != nil {
		return  err
	}
	if kTableMagicNumber != num {
		return   internal.ErrTableFileMagic
	}
	return  nil
}



