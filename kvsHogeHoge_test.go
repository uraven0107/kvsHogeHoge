package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_current_name(t *testing.T) {
	t.Run("canGetCurrentDatastoreName", func(t *testing.T) {
		assert := assert.New(t)
		ds := NewDatastore("hoge")
		app := Application{}
		app.current_ds = ds
		assert.Equal("hoge", app.Current_ds_name(), "current_name() doesn't equal")
	})
	t.Run("nilDatastore", func(t *testing.T) {
		assert := assert.New(t)
		app := Application{}
		assert.Equal("none", app.Current_ds_name(), "current_name() doesn't 'none'")
	})
}
