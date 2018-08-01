package statestore

import (
	"bytes"
	"testing"

	"github.com/anarcher/contrabase/pkg/contract/types"
)

func Test_StateCloneBatch_Commit(t *testing.T) {
	cloneBatch := NewStateCloneBatch(testStateStore)

	itemKey := &types.StorageItemKey{
		ContractAddress: types.AddressFromString("hello"),
		Key:             []byte("key1"),
	}

	item := &types.StorageItem{
		Value: []byte("item1"),
	}

	if err := cloneBatch.PutStorageItem(itemKey, item); err != nil {
		t.Error(err)
		return
	}

	if len(cloneBatch.items) != 1 {
		t.Errorf("cloneBatch.items have:%v want:%v", len(cloneBatch.items), 1)
		return
	}

	if item1, err := testStateStore.GetStorageItem(itemKey); err != nil {
		t.Error(err)
		return
	} else if err == nil {
		if item1 != nil {
			t.Errorf("item1 have:%v want:%v", item1, nil)
			return
		}
	}

	{
		mr1, err := cloneBatch.MerkleRoot()
		if err != nil {
			t.Error(err)
			return
		}
		if err := cloneBatch.PutStorageItem(itemKey, item); err != nil {
			t.Error(err)
			return
		}

		mr2, err := cloneBatch.MerkleRoot()
		if err != nil {
			t.Error(err)
			return
		}

		if !bytes.Equal(mr1, mr2) {
			t.Errorf("merkleroot: mr1:%v mr2:%v", mr1, mr2)
			return
		}
	}

	if err := cloneBatch.Commit(); err != nil {
		t.Error(err)
		return
	}

	if item1, err := testStateStore.GetStorageItem(itemKey); err != nil {
		t.Error(err)
		return
	} else if err == nil {
		if !bytes.Equal(item1.Value, item.Value) {
			t.Errorf("item1.Value have:%v want:%v", item1.Value, item.Value)
			return
		}
	}

}
