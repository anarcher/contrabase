package types

import (
	"io"
)

type Account struct {
	Address Address
}

func NewAccount(address Address) *Account {
	a := &Account{
		Address: address,
	}

	return a
}

func (a Account) String() string {
	return ""
}

func (a *Account) Serialize(w io.Writer) error {
	return nil
}

func (a *Account) Deserialize(r io.Reader) error {
	return nil
}
