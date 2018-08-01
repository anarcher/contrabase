package code

type ReturnCode struct {
	Values []*Value
}

func NewReturnCode(values ...*Value) *ReturnCode {
	r := &ReturnCode{
		Values: values,
	}
	return r
}
