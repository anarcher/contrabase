package db

type DB interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error

	Close() error

	NewBatch() Batch
}

type Batch interface {
	Put(key, value []byte)
	Delete(key []byte)
	Commit() error
	Reset()
}
