package sstable

import (
	"github.com/slowmoon/leveldb/sstable/block"
	"os"
)

type SsTable struct {
     block *block.Block
     footer Footer    // file structure
     file *os.File
}

