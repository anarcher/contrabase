package statestore

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/cbergoon/merkletree"
)

type ItemState byte

const (
	ItemStateNone ItemState = iota
	ItemStateChanged
	ItemStateDeleted
)

type StateItem struct {
	Key   []byte
	Value []byte
	State ItemState
}

func (i *StateItem) CalculateHash() ([]byte, error) {
	h := sha256.New()
	var v []byte
	v = append(v, i.Key...)
	v = append(v, i.Value...)
	v = append(v, byte(i.State))
	if _, err := h.Write(v); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (i *StateItem) Equals(other merkletree.Content) (bool, error) {
	o, ok := other.(*StateItem)
	if !ok {
		return false, fmt.Errorf("not merkletree.Content")
	}
	return bytes.Equal(i.Key, o.Key), nil
}
