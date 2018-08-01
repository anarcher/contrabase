package execfunc

import (
	"github.com/anarcher/contrabase/pkg/contract/code"
	"github.com/anarcher/contrabase/pkg/contract/native"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

var TransferAddress = types.AddressFromString("transfer")

func init() {
	//native.AddExecFunc("transfer", transfer)
	native.AddContract(TransferAddress, RegisterTransfer)
}

func RegisterTransfer(ex *native.NativeExecutor) {
	ex.RegisterFunc("transfer", transfer)
}

func transfer(ex *native.NativeExecutor, execCode *code.ExecCode) (*code.ReturnCode, error) {
	//TODO(anarcher)
	/*
		sender := cfg.Transaction.Sender
		receiver := cfg.Transaction.Receiver

		transferForm(sender)
	*/

	return nil, nil
}
