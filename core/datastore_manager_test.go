package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const db_file_dir = ".." + string(os.PathSeparator) + "test_data"
const db_file = "test.hogedb"
const db_file_path = db_file_dir + string(os.PathSeparator) + db_file

func TestNewDatastoreManager(t *testing.T) {
	t.Run("blankFilePathShouldReturnError", func(t *testing.T) {
		assert := assert.New(t)
		_, err := NewDatastoreManager("")
		assert.EqualError(err, "Configure.file_path is blank")
	})

	t.Run("canInitializeDatastoreManager", func(t *testing.T) {
		assert := assert.New(t)
		dm, _ := NewDatastoreManager(db_file_path)
		assert.NotNil(dm)
		assert.NotNil(dm.ds_list)
		assert.Equal(db_file_path, dm.db_file_path)
	})
}

func TestDatastoreManager_Persist(t *testing.T) {
	t.Run("canPersistDatabase", func(t *testing.T) {
		assert := assert.New(t)

		ds1 := NewDatastore("test1")
		ds1.Write("hoge", "fuga")
		ds1.Write("foo", "bar")
		ds2 := NewDatastore("test2")
		ds2.Write("baka", "aho")
		ds2.Write("unko", "brbr")
		ds2.Write("aaa", "zzz")

		dm, err := NewDatastoreManager(db_file_path)
		assert.Nil(err, "NewDatastoreManager() shouldn't return error")

		dm.AddDatastore(ds1)
		dm.AddDatastore(ds2)
		defer os.Remove(db_file_path)
		err = dm.Persist()
		assert.Nil(err, "DatastoreManager.Persist() shouldn't return error")

		bytes, err := ioutil.ReadFile(db_file_path)
		assert.Nil(err, "ioutil.ReadFile has return error")

		actual := string(bytes)
		expected := ds1.Persisted() + ds2.Persisted()
		assert.Equal(expected, actual, ":( File contents has not expected")
	})
}

func TestDatastoreManager_Restore(t *testing.T) {
	t.Run("canRestoreDatabaseFromFile", func(t *testing.T) {
		assert := assert.New(t)
		defer os.Remove(db_file_path)

		buf := []byte("test={hoge=fuga;foo=bar;};test2={baka=aho;unko=brbr;aaa=zzz;};")
		err := ioutil.WriteFile(db_file_path, buf, os.ModePerm)
		assert.Nil(err, fmt.Sprintf(":( Error has occured at ioutil.WriteFile(), error = %v", err))

		dm, err := NewDatastoreManager(db_file_path)
		assert.Nil(err, fmt.Sprintf(":( Error has occured at NewDatastoreManage(), error = %v", err))

		err2 := dm.Restore()
		assert.Nil(err2, fmt.Sprintf(":( Error has occured at DatastoreManager().Restore(), error = %v", err))

		err_msg := "Value of saved in Datastore (key = %v) hasn't equal"

		ds_test := dm.GetDatastore("test")
		assert.NotNil(ds_test, "dm.GetDatastore() shouldn't return nil")
		assert.Equal("fuga", ds_test.Read("hoge"), fmt.Sprintf(err_msg, "test"))
		assert.Equal("bar", ds_test.Read("foo"), fmt.Sprintf(err_msg, "test"))

		ds_test2 := dm.GetDatastore("test2")
		assert.NotNil(ds_test2, "dm.GetDatastore() shouldn't return nil")
		assert.Equal("aho", ds_test2.Read("baka"), fmt.Sprintf(err_msg, "test2"))
		assert.Equal("brbr", ds_test2.Read("unko"), fmt.Sprintf(err_msg, "test2"))
		assert.Equal("zzz", ds_test2.Read("aaa"), fmt.Sprintf(err_msg, "test2"))

		assert.Nil(dm.GetDatastore("test3"), "dm.GetDatastore() should return nil")

	})
}

func TestDatastoreManager_GetDatastore(t *testing.T) {

	dm := DatastoreManager{}

	t.Run("notInitializedDatastoreManagerShouldPanic", func(t *testing.T) {
		assert := assert.New(t)
		assert.Panics(func() {
			dm.GetDatastore("hoge")
		})
	})

	ds := NewDatastore("test")
	dm.ds_list = []*Datastore{ds}

	t.Run("getDatastoreShouldReturnDatastore", func(t *testing.T) {
		assert := assert.New(t)
		assert.NotNil(dm.GetDatastore("test"), ":( Datastore.getDatastore() shouldn't return nil")
	})

	t.Run("getDatastoreShouldReturnNil", func(t *testing.T) {
		assert := assert.New(t)
		assert.Nil(dm.GetDatastore("fuga"), ":( Datastore.getDatastore() should return nil")
	})
}

func Test_NewDatastoreManager(t *testing.T) {
	t.Run("couldNewDatastoreManager", func(t *testing.T) {
		assert := assert.New(t)
		dm := NewDatastore(db_file_path)
		assert.NotNil(dm, ":( NewDatastoreManager() shouldn't return nil")
	})
}
