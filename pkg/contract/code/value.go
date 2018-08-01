package code

import "io"

type Value struct {
	Type     Type
	Contents []byte
}

func (v *Value) Serialize(w io.Writer) error {

	return nil
}

func (v *Value) Deserialize(r io.Reader) error {

	return nil
}

func StringValue(s string) *Value {
	v := &Value{
		Type:     String,
		Contents: []byte(s),
	}
	return v
}
