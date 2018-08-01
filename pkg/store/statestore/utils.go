package statestore

import "bytes"

func StateKeyBytes(key StateKey) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := key.Serialize(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func StateValueBytes(value StateValue) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := value.Serialize(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func LoadStateValue(value []byte, stateValue StateValue) error {
	buf := bytes.NewBuffer(value)
	if err := stateValue.Deserialize(buf); err != nil {
		return err
	}
	return nil
}
