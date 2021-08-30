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
		assert.Equal("hoge", app.Current_name(), "current_name() doesn't equal")
	})
}
