package native

import (
	"fmt"

	"github.com/anarcher/contrabase/pkg/contract/code"
	"github.com/anarcher/contrabase/pkg/contract/context"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

type ExecFunc func(e *NativeExecutor, code *code.ExecCode) (*code.ReturnCode, error)

type NativeExecutor struct {
	Context   *context.Context
	execFuncs map[string]ExecFunc
}

func NewNativeExecutor(ctx *context.Context) *NativeExecutor {
	ex := &NativeExecutor{
		Context:   ctx,
		execFuncs: map[string]ExecFunc{},
	}

	return ex
}

func (ex *NativeExecutor) Execute(c *code.ExecCode) (*code.ReturnCode, error) {
	//TODO(anarcher)
	ex.loadFuncs(c.ContractAddress)

	if f, ok := ex.execFuncs[c.Method]; ok {
		returnCode, err := f(ex, c)
		return returnCode, err
	}

	//TODO(anarcher) define error
	return nil, fmt.Errorf("not found")
}

func (ex *NativeExecutor) RegisterFunc(name string, f ExecFunc) {
	ex.execFuncs[name] = f
}

func (ex *NativeExecutor) loadFuncs(addr types.Address) {
	if r, ok := contracts[addr.String()]; ok {
		r(ex)
	}
}
