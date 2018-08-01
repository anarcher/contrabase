package statestore

import (
	"bytes"
	"testing"

	"github.com/anarcher/contrabase/pkg/contract/types"
)

func Test_StateStore_StorageGet(t *testing.T) {
	itemKey := &types.StorageItemKey{
		ContractAddress: types.AddressFromString("helloworld"),
		Key:             []byte("key1"),
	}
	item := &types.StorageItem{
		Value: []byte("item1"),
	}

	itemKeyB := bytes.NewBuffer(nil)
	if err := itemKey.Serialize(itemKeyB); err != nil {
		t.Error(err)
		return
	}
	itemB := bytes.NewBuffer(nil)
	if err := item.Serialize(itemB); err != nil {
		t.Error(err)
		return
	}
	if err := testDB.Put(itemKeyB.Bytes(), itemB.Bytes()); err != nil {
		t.Error(err)
		return
	}

	itemR, err := testStateStore.GetStorageItem(itemKey)
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(itemR.Value, item.Value) {
		t.Errorf("item have:%v want:%v", itemR.Value, item.Value)
		return
	}
}
