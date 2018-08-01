package statestore

import (
	"fmt"
	"os"
	"testing"

	"github.com/anarcher/contrabase/pkg/db"
	"github.com/anarcher/contrabase/pkg/db/leveldb"
)

var (
	testDB         db.DB
	testStateStore *StateStore
)

func TestMain(m *testing.M) {
	var err error
	var path = "./test"

	testDB, err = leveldb.NewLevelDB(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}

	testStateStore = NewStateStore(testDB)

	m.Run()

	err = testStateStore.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}

	err = os.RemoveAll(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}
}
