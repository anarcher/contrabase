package merkletree

type MerkleTreeDB interface {
	Set(key, value []byte) error
	Get(key []byte) (value []byte, err error)
	Delete(key []byte) error

	Hash() []byte //current working root hash

	CommitDB() ([]byte, error)
}

//  MPTDB(eth) and IAVLDB(tender)
