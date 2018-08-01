package native

import "github.com/anarcher/contrabase/pkg/contract/types"

var (
	contracts = make(map[string]Register)
)

type (
	Register func(executor *NativeExecutor)
)

func AddContract(address types.Address, r Register) {
	contracts[address.String()] = r
}

func HasContract(address types.Address) bool {
	if _, ok := contracts[address.String()]; ok {
		return true
	}

	return false
}
