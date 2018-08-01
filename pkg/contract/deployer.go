package contract

type Deployer interface {
	Deploy(code []byte) error
}
