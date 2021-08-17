package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatastore_Write(t *testing.T) {
	assert := assert.New(t)
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
			got := ds.Size()
			assert.Equal(tt.want, got, fmt.Sprintf(":( Datastore.Size() = %v, want %v", got, tt.want))
		})
	}
}

func TestDatastore_Read(t *testing.T) {
	assert := assert.New(t)
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
				got := ds.Read(argAndWant.key)
				assert.Equal(argAndWant.want, got, fmt.Sprintf(":( Datastore.Read() = %v, want %v", got, argAndWant.want))
			}
		})
	}
}

func TestDatastore_Delete(t *testing.T) {
	assert := assert.New(t)
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
			got := ds.Read(tt.args.key)
			assert.Empty(got, ":( Datastore.Delete() dosen't work!")
			got_size := ds.Size()
			assert.Equal(got_size, tt.want_size, fmt.Sprintf(":( Datastore.Size() = %v, want %v", got_size, tt.want_size))
		})
	}
}

func TestDatastore_Persisted(t *testing.T) {
	assert := assert.New(t)
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
			got := ds.Persisted()
			assert.Equal(got, tt.want, fmt.Sprintf(":( Datastore.Persisted() = %v, want %v", got, tt.want))
		})
	}
}
