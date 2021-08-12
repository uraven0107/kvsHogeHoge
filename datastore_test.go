package main

import (
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
			ds:   NewDatastore("test"),
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
				t.Errorf(":( Datastore.Size() = %v, want %v", got, tt.want)
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
			ds:   NewDatastore("test"),
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
					t.Errorf(":( Datastore.Read() = %v, want %v", got, argAndWant.want)
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
			ds:   NewDatastore("test"),
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
				t.Errorf(":( Datastore.Delete() = %v, want %v", got, nil)
			}
			if got := ds.Size(); got != tt.want_size {
				t.Errorf(":( Datastore.Size() = %v, want %v", got, tt.want_size)
			}
		})
	}
}

func TestDatastore_Persisted(t *testing.T) {
	tests := []struct {
		name string
		ds   *Datastore
		want string
	}{
		{
			name: "canConvertPersisted",
			ds:   NewDatastore("test"),
			want: "test={hoge=fuga;foo=bar;};",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := tt.ds
			ds.Write("hoge", "fuga")
			ds.Write("foo", "bar")
			if got := ds.Persisted(); got != tt.want {
				t.Errorf(":( Datastore.Persisted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatastore_Restore(t *testing.T) {
	const file_path = "test_data/test.hogedb"
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
			name: "canRestore",
			ds:   NewDatastore("test"),
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
			ds := Restore(file_path)
			for _, argAndWant := range tt.argsAndWants {
				if got := ds.Read(argAndWant.key); got != argAndWant.want {
					t.Errorf(":( Datastore.Read() = %v, want %v", got, argAndWant.want)
				}
			}
		})
	}
}
