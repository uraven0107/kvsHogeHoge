package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_current_name(t *testing.T) {
	t.Run("canGetCurrentDatastoreName", func(t *testing.T) {
		assert := assert.New(t)
		ds := NewDatastore("hoge")
		app := &Application{}
		app.current_ds = ds
		assert.Equal("hoge", app.Current_ds_name(), "current_name() doesn't equal")
	})
	t.Run("nilDatastore", func(t *testing.T) {
		assert := assert.New(t)
		app := &Application{}
		assert.Equal("none", app.Current_ds_name(), "current_name() doesn't 'none'")
	})
}

func Test_switch_ds(t *testing.T) {
	t.Run("canSwitchDatastore", func(t *testing.T) {
		assert := assert.New(t)
		ds1 := NewDatastore("hoge")
		ds2 := NewDatastore("fuga")
		ds_list := []*Datastore{ds1, ds2}
		manager := &DatastoreManager{}
		manager.ds_list = ds_list
		app := &Application{}
		app.ds_manager = manager
		err := app.Switch_ds("hoge")
		assert.Nil(err, "Switch_ds() shoudn't return error")
		assert.Equal(ds1, app.current_ds, "current_ds hasn't expected")
		err = app.Switch_ds("fuga")
		assert.Nil(err, "Switch_ds() shouldn't return error")
		assert.Equal(ds2, app.current_ds, "current_ds hasn't expected")
		err = app.Switch_ds("unko")
		assert.Error(err, "Switch_ds() should return error")
	})
}
