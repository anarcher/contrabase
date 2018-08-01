package contract

import (
	"fmt"
	"os"
	"testing"

	"github.com/anarcher/contrabase/pkg/db"
	"github.com/anarcher/contrabase/pkg/db/leveldb"
	"github.com/anarcher/contrabase/pkg/store/statestore"
)

var (
	testDB         db.DB
	testStateStore *statestore.StateStore
	testStateClone *statestore.StateCloneBatch
)

func TestMain(m *testing.M) {
	var err error
	var path = "./test"

	testDB, err = leveldb.NewLevelDB(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}

	testStateStore = statestore.NewStateStore(testDB)
	testStateClone = statestore.NewStateCloneBatch(testStateStore)

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
