package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const file_path = "test_data/test.hogedb"

func TestDatastoreManager_Persist(t *testing.T) {
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

func TestDatastoreManager_Restore(t *testing.T) {
	ds1 := NewDatastore("test")
	ds1.Write("hoge", "fuga")
	ds1.Write("foo", "bar")
	ds2 := NewDatastore("test2")
	ds2.Write("baka", "aho")
	ds2.Write("unko", "brbr")
	ds2.Write("aaa", "zzz")
	type field struct {
		file_path string
	}
	tests := []struct {
		name  string
		field field
		wants []*Datastore
	}{
		{
			name: "canRestoreDatabaseFromFile",
			field: field{
				file_path: file_path,
			},
			wants: []*Datastore{ds1, ds2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := []byte("test={hoge=fuga;foo=bar;};test2={baka=aho;unko=brbr;aaa=zzz;};")
			err := ioutil.WriteFile(file_path, buf, os.ModePerm)
			if err != nil {
				t.Errorf(":( Error has occured at ioutil.WriteFile(), error = %v", err)
				return
			}
			dm := DatastoreManager{
				file_path: tt.field.file_path,
				ds_list:   []*Datastore{},
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
			os.Remove(file_path)
		})
	}
}
