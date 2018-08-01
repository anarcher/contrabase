package statestore

import "io"

type StateKey interface {
	String() string
	Serialize(io.Writer) error
	Deserialize(io.Reader) error
}

type StateValue interface {
	Serialize(io.Writer) error
	Deserialize(io.Reader) error
}
