package leveldb

import "github.com/syndtr/goleveldb/leveldb"

type LevelDBBatch struct {
	db    *leveldb.DB
	batch *leveldb.Batch
}

func NewLevelDBBatch(db *leveldb.DB) *LevelDBBatch {
	b := &LevelDBBatch{
		db:    db,
		batch: new(leveldb.Batch),
	}
	return b
}

func (b *LevelDBBatch) Put(key, value []byte) {
	b.batch.Put(key, value)
}

func (b *LevelDBBatch) Delete(key []byte) {
	b.batch.Delete(key)
}

func (b *LevelDBBatch) Commit() error {
	if err := b.db.Write(b.batch, nil); err != nil {
		return err
	}
	return nil
}

func (b *LevelDBBatch) Reset() {
	b.batch.Reset()
}
