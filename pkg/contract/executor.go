package contract

import (
	"fmt"

	"github.com/anarcher/contrabase/pkg/contract/code"
	"github.com/anarcher/contrabase/pkg/contract/context"
	"github.com/anarcher/contrabase/pkg/contract/native"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

type Executor interface {
	Execute(*code.ExecCode) (*code.ReturnCode, error)
}

func NewExecutor(ctx *context.Context, contractAddress types.Address) (Executor, error) {
	var ex Executor
	if native.HasContract(contractAddress) {
		ex = native.NewNativeExecutor(ctx)
	}
	return ex, nil
}

func Execute(ctx *context.Context, execCode *code.ExecCode) (*code.ReturnCode, error) {
	ex, err := NewExecutor(ctx, execCode.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("not found")
	}
	if ex == nil {
		return nil, fmt.Errorf("not found")
	}
	ret, err := ex.Execute(execCode)
	return ret, err
}
