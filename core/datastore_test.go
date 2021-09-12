package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatastore_Write(t *testing.T) {
	t.Run("canWrite", func(t *testing.T) {
		assert := assert.New(t)
		ds := NewDatastore("test")
		ds.Write("hoge", "fuga")
		ds.Write("foo", "bar")
		expected := 2
		assert.Equal(expected, ds.Size(), "Datastore.Size() hasn't equal")
	})
}

func TestDatastore_Read(t *testing.T) {
	t.Run("canRead", func(t *testing.T) {
		assert := assert.New(t)
		ds := NewDatastore("test")
		ds.Write("hoge", "fuga")
		ds.Write("foo", "bar")

		err_msg := "Datastore.Read() hasn't expected"
		assert.Equal("fuga", ds.Read("hoge"), err_msg)
		assert.Equal("bar", ds.Read("foo"), err_msg)
	})
}

func TestDatastore_Delete(t *testing.T) {
	t.Run("canDelete", func(t *testing.T) {
		assert := assert.New(t)
		ds := NewDatastore("test")
		ds.Write("hoge", "fuga")
		ds.Write("foo", "bar")

		key := "hoge"
		ds.Delete(key)
		assert.Empty(ds.Read(key), fmt.Sprintf("Datastore.Read() should return empty, key = %v", key))

		expected := 1
		assert.Equal(expected, ds.Size(), "Datastore.Size() hasn't equal. Datastore.Delete() might not work")
	})
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
