package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const db_file_dir = ".." + string(os.PathSeparator) + "test_data"
const db_file = "test.hogedb"
const db_file_path = db_file_dir + string(os.PathSeparator) + db_file

func TestNewDatastoreManager(t *testing.T) {
	t.Run("nilConfigShouldReturnError", func(t *testing.T) {
		assert := assert.New(t)
		_, err := NewDatastoreManager(nil)
		assert.EqualError(err, "Configure is nil")
	})

	t.Run("blankFilePathShouldReturnError", func(t *testing.T) {
		assert := assert.New(t)
		config := new(Configure)
		config.file_path = ""
		_, err := NewDatastoreManager(config)
		assert.EqualError(err, "Configure.file_path is blank")
	})

	t.Run("canInitializeDatastoreManager", func(t *testing.T) {
		assert := assert.New(t)
		config := new(Configure)
		config.file_path = db_file_path
		dm, _ := NewDatastoreManager(config)
		assert.NotNil(dm)
		assert.NotNil(dm.ds_list)
		assert.Equal(db_file_path, dm.db_file_path)
	})
}

func TestDatastoreManager_Persist(t *testing.T) {
	ds1 := NewDatastore("test1")
	ds1.Write("hoge", "fuga")
	ds1.Write("foo", "bar")
	ds2 := NewDatastore("test2")
	ds2.Write("baka", "aho")
	ds2.Write("unko", "brbr")
	ds2.Write("aaa", "zzz")
	type fields struct {
		ds_list []*Datastore
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "canPersistDatabase",
			fields: fields{
				ds_list: []*Datastore{ds1, ds2},
			},
			want: ds1.Persisted() + ds2.Persisted(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := new(Configure)
			config.file_path = db_file_path
			dm, err := NewDatastoreManager(config)
			if err != nil {
				t.Errorf(":( Error has occured at NewDatastoreManage(), error = %v", err)
			}
			dm.ds_list = tt.fields.ds_list
			if err := dm.Persist(); err != nil {
				t.Errorf(":( Error has occured at DatastoreManager().Persist(), error = %v", err)
				return
			}
			bytes, err := ioutil.ReadFile(db_file_path)
			if err != nil {
				t.Errorf(":( Error has occured at ioutil.ReadFile(), error = %v", err)
				return
			}
			str := string(bytes)
			if str != tt.want {
				t.Errorf(":( File contents has not expected, want = %v, but got = %v", tt.want, str)
			}
			defer os.Remove(db_file_path)
		})
	}
}

func TestDatastoreManager_Restore(t *testing.T) {
	defer os.Remove(db_file_path)
	ds1 := NewDatastore("test")
	ds1.Write("hoge", "fuga")
	ds1.Write("foo", "bar")
	ds2 := NewDatastore("test2")
	ds2.Write("baka", "aho")
	ds2.Write("unko", "brbr")
	ds2.Write("aaa", "zzz")
	tests := []struct {
		name  string
		wants []*Datastore
	}{
		{
			name:  "canRestoreDatabaseFromFile",
			wants: []*Datastore{ds1, ds2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := []byte("test={hoge=fuga;foo=bar;};test2={baka=aho;unko=brbr;aaa=zzz;};")
			err := ioutil.WriteFile(db_file_path, buf, os.ModePerm)
			if err != nil {
				t.Errorf(":( Error has occured at ioutil.WriteFile(), error = %v", err)
				return
			}
			config := new(Configure)
			config.file_path = db_file_path
			dm, err := NewDatastoreManager(config)
			if err != nil {
				t.Errorf(":( Error has occured at NewDatastoreManage(), error = %v", err)
			}
			if err := dm.Restore(); err != nil {
				t.Errorf(":( Error has occured at DatastoreManager().Restore(), error = %v", err)
				return
			}
			for _, want := range tt.wants {
				notMatched := true
				for _, ds := range dm.ds_list {
					if ds.name == want.name {
						notMatched = false
						for k, v := range ds.store {
							if want.store[k] != v {
								t.Errorf(":( DatastoreManager().Rstore() has not expected. key = %v, want = %v, but got = %v", k, want.store[k], v)
							}
						}
					}
				}

				if notMatched {
					t.Error(":( Nothing matched! DatastoreManager().Restore() dosen't work. want.name = " + want.name)
					return
				}
			}
		})
	}
}

func TestDatastoreManager_getDatastore(t *testing.T) {

	dm := DatastoreManager{}

	t.Run("notInitializedDatastoreManagerShouldPanic", func(t *testing.T) {
		assert := assert.New(t)
		assert.Panics(func() {
			dm.getDatastore("hoge")
		})
	})

	ds := NewDatastore("test")
	dm.ds_list = []*Datastore{ds}

	t.Run("getDatastoreShouldReturnDatastore", func(t *testing.T) {
		assert := assert.New(t)
		assert.NotNil(dm.getDatastore("test"), ":( Datastore.getDatastore() shouldn't return nil")
	})

	t.Run("getDatastoreShouldReturnNil", func(t *testing.T) {
		assert := assert.New(t)
		assert.Nil(dm.getDatastore("fuga"), ":( Datastore.getDatastore() should return nil")
	})
}

func Test_NewDatastoreManager(t *testing.T) {
	t.Run("couldNewDatastoreManager", func(t *testing.T) {
		assert := assert.New(t)
		dm := NewDatastore(db_file_path)
		assert.NotNil(dm, ":( NewDatastoreManager() shouldn't return nil")
	})
}
