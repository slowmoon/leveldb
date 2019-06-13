package version

import "github.com/slowmoon/leveldb/internal"

type FileMetaData struct {

	allowSeeks uint64
	number uint64
	fileSize uint64
    smallest *internal.InternalKey
    biggest *internal.InternalKey
}


