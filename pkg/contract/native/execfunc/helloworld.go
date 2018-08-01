package execfunc

import (
	"encoding/json"

	"github.com/anarcher/contrabase/pkg/contract/code"
	"github.com/anarcher/contrabase/pkg/contract/native"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

var HelloWorldAddress = types.AddressFromString("HELLOWORLDADDRESS")

func init() {
	native.AddContract(HelloWorldAddress, RegisterHelloWorld)
}

func RegisterHelloWorld(ex *native.NativeExecutor) {
	ex.RegisterFunc("hello", hello)
}

func hello(ex *native.NativeExecutor, execCode *code.ExecCode) (*code.ReturnCode, error) {
	stateClone := ex.Context.StateClone
	sender := ex.Context.SenderAccount

	greeterKey := &types.StorageItemKey{
		ContractAddress: execCode.ContractAddress, // or HelloWorldAddress
		Key:             []byte("greeters"),
	}

	item, err := stateClone.GetStorageItem(greeterKey)
	if err != nil {
		return nil, err
	}

	var greeters []types.Address
	if item == nil {
		item = &types.StorageItem{}
	} else {
		err := json.Unmarshal(item.Value, greeters)
		if err != nil {
			return nil, err
		}
	}

	greeters = append(greeters, sender.Address)

	{
		b, err := json.Marshal(greeters)
		if err != nil {
			return nil, err
		}
		item.Value = b
	}

	if err := stateClone.PutStorageItem(greeterKey, item); err != nil {
		return nil, err
	}

	v := code.StringValue("world")
	rCode := code.NewReturnCode(v)
	return rCode, nil
}
