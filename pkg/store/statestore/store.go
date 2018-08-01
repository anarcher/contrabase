package statestore

import (
	"github.com/anarcher/contrabase/pkg/contract/types"
	"github.com/anarcher/contrabase/pkg/db"
)

type StateStore struct {
	db db.DB
}

func NewStateStore(db db.DB) *StateStore {
	s := &StateStore{
		db: db,
	}

	return s
}

func (s *StateStore) DB() db.DB {
	return s.db
}

func (s *StateStore) Close() error {
	return s.db.Close()
}

func (s *StateStore) GetAccount(addr types.Address) (*types.Account, error) {
	k, err := StateKeyBytes(addr)
	if err != nil {
		return nil, err
	}

	v, err := s.db.Get(k)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}

	var account *types.Account
	if err := LoadStateValue(v, account); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *StateStore) GetStorageItem(itemKey *types.StorageItemKey) (*types.StorageItem, error) {
	k, err := StateKeyBytes(itemKey)
	if err != nil {
		return nil, err
	}

	v, err := s.db.Get(k)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}

	item := new(types.StorageItem)
	if err := LoadStateValue(v, item); err != nil {
		return nil, err
	}

	return item, nil
}
