package leveldb


type DB interface {

	Put(key , value []byte)error

	Get(key []byte) ([]byte, error)

	Delete(key []byte) error

}

type Iterator interface {

    Valid() bool

    Key()interface{}

    Value()interface{}

    Next()

    Prev()

    Seek([]byte)

    SeekToFirst()

    SeekToLast()
}





