package leveldb

import (
	"github.com/anarcher/contrabase/pkg/db"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const BITSPERKEY = 10

// LevelDB
type LevelDB struct {
	db *leveldb.DB
}

// NewStateDB return LevelDB instance
func NewLevelDB(file string) (*LevelDB, error) {
	o := opt.Options{
		NoSync: false,
		Filter: filter.NewBloomFilter(BITSPERKEY),
	}
	db, err := leveldb.OpenFile(file, &o)

	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}

	if err != nil {
		return nil, err
	}

	return &LevelDB{
		db: db,
	}, nil
}

// Get value from key from leveldb
func (l *LevelDB) Get(key []byte) ([]byte, error) {
	val, err := l.db.Get(key, nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return val, err
}

// Put key value pair to leveldb
func (l *LevelDB) Put(key []byte, value []byte) error {
	return l.db.Put(key, value, nil)
}

// Has return whether the key is exist in leveldb
func (l *LevelDB) Has(key []byte) (bool, error) {
	return l.db.Has(key, nil)
}

// Delete item in leveldb
func (l *LevelDB) Delete(key []byte) error {
	return l.db.Delete(key, nil)
}

func (l *LevelDB) Close() error {
	err := l.db.Close()
	return err
}

func (l *LevelDB) NewBatch() db.Batch {
	b := NewLevelDBBatch(l.db)
	return b
}
