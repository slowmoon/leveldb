package internal

import "errors"

var (

	ErrNotFound = errors.New("error not found")

	ErrDeletion = errors.New("type deletion")

	ErrTableFileMagic = errors.New("not an sstable(bad maginc number)")

	ErrTableFileTooShort = errors.New("file is too short to be an sstable")
)
