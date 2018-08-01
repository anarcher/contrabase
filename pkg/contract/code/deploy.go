package code

import "github.com/anarcher/contrabase/pkg/contract/types"

type DeployCode struct {
	ContractAddress types.Address
	VM              string // wasm,otto,...
	Name            string
	Code            []byte
	Description     string
}
