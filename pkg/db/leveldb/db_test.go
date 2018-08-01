package leveldb

import (
	"fmt"
	"os"
	"testing"
)

var testLevelDB *LevelDB

func TestMain(m *testing.M) {
	dbFile := "./test"
	var err error
	testLevelDB, err = NewLevelDB(dbFile)
	if err != nil {
		fmt.Printf("NewLevelDB err:%s\n", err)
		return
	}
	m.Run()
	testLevelDB.Close()
	os.RemoveAll(dbFile)
}

func TestLevelDB(t *testing.T) {
	key := "foo"
	value := "bar"
	err := testLevelDB.Put([]byte(key), []byte(value))
	if err != nil {
		t.Errorf("Put err:%v\n", err)
		return
	}

	v, err := testLevelDB.Get([]byte(key))
	if err != nil {
		t.Errorf("Get err:%v\n", err)
		return
	}

	if string(v) != value {
		t.Errorf("Get err have:%v want:%v\n", v, value)
		return
	}

	err = testLevelDB.Delete([]byte(key))
	if err != nil {
		t.Errorf("Delete err:%v", err)
		return
	}
}
