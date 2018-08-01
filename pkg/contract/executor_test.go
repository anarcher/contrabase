package contract

import (
	"testing"

	"github.com/anarcher/contrabase/pkg/contract/code"
	"github.com/anarcher/contrabase/pkg/contract/context"
	"github.com/anarcher/contrabase/pkg/contract/native/execfunc"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

func Test_Executor_Native_HelloWorld(t *testing.T) {
	ctx := &context.Context{
		SenderAccount: types.NewAccount(types.AddressFromString("sender1")),
		StateStore:    testStateStore,
		StateClone:    testStateClone,
	}

	exCode := &code.ExecCode{
		ContractAddress: execfunc.HelloWorldAddress,
		Method:          "hello",
	}

	ex, err := NewExecutor(ctx, exCode.ContractAddress)
	if err != nil {
		t.Error(err)
		return
	}

	retCode, err := ex.Execute(exCode)
	if err != nil {
		t.Error(err)
		return
	}

	{
		v := retCode.Values[0]
		if v.Type != code.String {
			t.Errorf("v.Type have:%v want:%v", v.Type, code.String)
			return
		}
		if string(v.Contents) != "world" {
			t.Errorf("v.Contents have:%s want:%v", v.Contents, "world")
			return
		}
	}

	if retCode == nil {
		t.Errorf("retCode have:%v want:%v", retCode, nil)
		return
	}
}
