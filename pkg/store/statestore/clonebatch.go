package statestore

import (
	"github.com/cbergoon/merkletree"

	"github.com/anarcher/contrabase/pkg/db"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

type StateCloneBatch struct {
	db    db.DB
	batch db.Batch

	store *StateStore

	items map[string]*StateItem
}

func NewStateCloneBatch(store *StateStore) *StateCloneBatch {
	db := store.DB()
	batch := db.NewBatch()

	s := &StateCloneBatch{
		db:    db,
		batch: batch,
		store: store,
		items: make(map[string]*StateItem),
	}

	return s
}

// Commit current changed states to db
func (s *StateCloneBatch) Commit() error {
	for _, v := range s.items {
		switch v.State {
		case ItemStateDeleted:
			s.batch.Delete(v.Key)
		case ItemStateChanged:
			s.batch.Put(v.Key, v.Value)
		}
	}

	if err := s.batch.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *StateCloneBatch) MerkleRoot() ([]byte, error) {
	var ls []merkletree.Content
	for _, i := range s.items {
		ls = append(ls, i)
	}

	t, err := merkletree.NewTree(ls)
	if err != nil {
		return nil, err
	}

	mr := t.MerkleRoot()
	return mr, nil
}

func (s *StateCloneBatch) GetAccount(addr types.Address) (*types.Account, error) {
	if s.isStateDeleted(addr) {
		return nil, nil
	}

	a, err := s.store.GetAccount(addr)
	if err != nil {
		return nil, err
	}
	return a, nil
}

//TODO(anarcher)
func (s *StateCloneBatch) ChangeAccountBalance() {

}

func (s *StateCloneBatch) GetStorageItem(key *types.StorageItemKey) (*types.StorageItem, error) {
	if s.isStateDeleted(key) {
		return nil, nil
	}

	if v, ok := s.items[key.String()]; ok {
		item := new(types.StorageItem)
		if err := LoadStateValue(v.Value, item); err != nil {
			return nil, err
		}
		return item, nil
	}

	item, err := s.store.GetStorageItem(key)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *StateCloneBatch) PutStorageItem(key *types.StorageItemKey, value *types.StorageItem) error {
	keyBytes, err := StateKeyBytes(key)
	if err != nil {
		return err
	}
	valueBytes, err := StateValueBytes(value)
	if err != nil {
		return err
	}

	s.items[key.String()] = &StateItem{
		Key:   keyBytes,
		Value: valueBytes,
		State: ItemStateChanged,
	}
	return nil
}

func (s *StateCloneBatch) DeleteStorageItem(key *types.StorageItemKey) error {
	if v, ok := s.items[key.String()]; ok {
		v.Value = nil
		v.State = ItemStateChanged
	} else {
		keyBytes, err := StateKeyBytes(key)
		if err != nil {
			return err
		}

		s.items[key.String()] = &StateItem{
			Key:   keyBytes,
			State: ItemStateDeleted,
		}
	}
	return nil
}

func (s *StateCloneBatch) isStateDeleted(key StateKey) bool {
	if v, ok := s.items[key.String()]; ok {
		if v.State == ItemStateDeleted {
			return false
		}
		return true
	}
	return false
}
