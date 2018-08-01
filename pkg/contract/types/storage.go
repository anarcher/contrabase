package types

import (
	"bytes"
	"io"

	"github.com/anarcher/contrabase/pkg/serialization"
)

// Storage is a storage for contract

// StorageItemKey is a key of Contract Storage Item
type StorageItemKey struct {
	ContractAddress Address
	Key             []byte
}

type StorageItem struct {
	Value []byte
}

func (s *StorageItemKey) Serialize(w io.Writer) error {
	if err := s.ContractAddress.Serialize(w); err != nil {
		return err
	}
	if err := serialization.WriteVarBytes(w, s.Key); err != nil {
		return err
	}
	return nil
}

func (s *StorageItemKey) Deserialize(r io.Reader) error {
	if err := s.ContractAddress.Deserialize(r); err != nil {
		return err
	}
	key, err := serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}
	s.Key = key
	return nil
}

func (s *StorageItemKey) String() string {
	bf := bytes.NewBuffer(nil)
	s.Deserialize(bf)
	return bf.String()
}

func (s *StorageItem) Serialize(w io.Writer) error {
	if err := serialization.WriteVarBytes(w, s.Value); err != nil {
		return err
	}

	return nil
}

func (s *StorageItem) Deserialize(r io.Reader) error {
	v, err := serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}
	s.Value = v
	return nil

}
