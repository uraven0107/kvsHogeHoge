package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDatastoreManager_Persist(t *testing.T) {
	file_path := "test_data/test.hogedb"
	ds1 := NewDatastore("test1")
	ds1.Write("hoge", "fuga")
	ds1.Write("foo", "bar")
	ds2 := NewDatastore("test2")
	ds2.Write("baka", "aho")
	ds2.Write("unko", "brbr")
	ds2.Write("aaa", "zzz")
	type fields struct {
		file_path string
		ds_list   []*Datastore
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "canPersistDatabase",
			fields: fields{
				file_path: file_path,
				ds_list:   []*Datastore{ds1, ds2},
			},
			want: ds1.Persisted() + ds2.Persisted(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := DatastoreManager{
				file_path: tt.fields.file_path,
				ds_list:   tt.fields.ds_list,
			}
			if err := dm.Persist(); err != nil {
				t.Errorf(":( Error has occured at DatastoreManager().Persist(), error = %v", err)
				return
			}
			bytes, err := ioutil.ReadFile(file_path)
			if err != nil {
				t.Errorf(":( Error has occured at ioutil.ReadFile(), error = %v", err)
				return
			}
			str := string(bytes)
			if str != tt.want {
				t.Errorf(":( File contents has not expected, want = %v, but got = %v", tt.want, str)
			}
			os.Remove(file_path)
		})
	}
}
