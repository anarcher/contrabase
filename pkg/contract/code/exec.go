package code

import "github.com/anarcher/contrabase/pkg/contract/types"

type ExecCode struct {
	ContractAddress types.Address
	Method          string
	Args            []*Value
}
