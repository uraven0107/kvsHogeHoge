package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uraven0107/kvsHogeHoge/core"
)

func Test_current_name(t *testing.T) {
	t.Run("canGetCurrentDatastoreName", func(t *testing.T) {
		assert := assert.New(t)
		ds := core.NewDatastore("hoge")
		ctx := &AppContext{}
		ctx.current_ds = ds
		assert.Equal("hoge", ctx.Current_ds_name(), "current_name() doesn't equal")
	})
	t.Run("nilDatastore", func(t *testing.T) {
		assert := assert.New(t)
		ctx := &AppContext{}
		assert.Equal("none", ctx.Current_ds_name(), "current_name() doesn't 'none'")
	})
}

func Test_switch_ds(t *testing.T) {
	t.Run("canSwitchDatastore", func(t *testing.T) {
		assert := assert.New(t)
		ds1 := core.NewDatastore("hoge")
		ds2 := core.NewDatastore("fuga")
		app := &Application{}
		ds_manager, _ := core.NewDatastoreManager("dummy")
		app.ds_manager = ds_manager
		app.ds_manager.AddDatastore(ds1)
		app.ds_manager.AddDatastore(ds2)
		ctx := &AppContext{}
		ctx.app = app
		err := ctx.Switch_ds("hoge")
		assert.Nil(err, "Switch_ds() shoudn't return error")
		assert.Equal(ds1, ctx.current_ds, "current_ds hasn't expected")
		err = ctx.Switch_ds("fuga")
		assert.Nil(err, "Switch_ds() shouldn't return error")
		assert.Equal(ds2, ctx.current_ds, "current_ds hasn't expected")
		err = ctx.Switch_ds("unko")
		assert.Error(err, "Switch_ds() should return error")
	})
}
