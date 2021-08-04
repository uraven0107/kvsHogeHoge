package main

import (
	"os"
	"testing"
)

func TestDatastore_Write(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name      string
		ds        *Datastore
		args_list []args
		want      int
	}{
		{
			name: "canWrite",
			ds:   NewDatastore(),
			args_list: []args{
				{
					key:   "hoge",
					value: "fuga",
				},
				{
					key:   "foo",
					value: "bar",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := tt.ds
			for _, args := range tt.args_list {
				ds.Write(args.key, args.value)
			}
			if got := ds.Size(); got != tt.want {
				t.Errorf("Datastore.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatastore_Read(t *testing.T) {
	type argAndWant struct {
		key  string
		want string
	}
	tests := []struct {
		name         string
		ds           *Datastore
		argsAndWants []argAndWant
	}{
		{
			name: "canRead",
			ds:   NewDatastore(),
			argsAndWants: []argAndWant{
				{
					key:  "hoge",
					want: "fuga",
				},
				{
					key:  "foo",
					want: "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := tt.ds
			ds.Write("hoge", "fuga")
			ds.Write("foo", "bar")
			for _, argAndWant := range tt.argsAndWants {
				if got := ds.Read(argAndWant.key); got != argAndWant.want {
					t.Errorf("Datastore.Read() = %v, want %v", got, argAndWant.want)
				}
			}
		})
	}
}

func TestDatastore_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		ds        *Datastore
		args      args
		want_size int
	}{
		{
			name: "canDelete",
			ds:   NewDatastore(),
			args: args{
				key: "hoge",
			},
			want_size: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := tt.ds
			ds.Write("hoge", "fuga")
			ds.Write("foo", "bar")
			ds.Delete("hoge")
			if got := ds.Read(tt.args.key); got != "" {
				t.Errorf("Datastore.Delete() = %v, want %v", got, nil)
			}
			if got := ds.Size(); got != tt.want_size {
				t.Errorf("Datastore.Size() = %v, want %v", got, tt.want_size)
			}
		})
	}
}

func TestDatastore_Persist(t *testing.T) {
	const file_path = "test_data/test.hogedb"
	tests := []struct {
		name string
		ds   *Datastore
	}{
		{
			name: "canPersist",
			ds:   NewDatastore(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := tt.ds
			ds.Write("hoge", "fuga")
			ds.Write("foo", "bar")
			ds.Persist(file_path)
			if f, err := os.Open(file_path); err != nil {
				t.Errorf("Error has ocuured in Datastore.Persist(); error = %v", err.Error())
			} else {
				f.Read([]byte{})
			}
		})
	}
}
