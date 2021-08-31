package main

import (
	"fmt"
)

type Application struct {
	ds_manager *DatastoreManager
}

type Context struct {
	app        *Application
	current_ds *Datastore
}

func NewContext(app *Application) *Context {
	return &Context{
		app:        app,
		current_ds: nil,
	}
}

func (ctx *Context) Current_ds_name() string {
	if ctx.current_ds == nil {
		return "none"
	}
	return ctx.current_ds.name
}

func (ctx *Context) Switch_ds(ds_name string) error {
	ds := ctx.app.ds_manager.getDatastore(ds_name)
	if ds == nil {
		return fmt.Errorf("Datastore named '%v' dosen't exist", ds_name)
	}
	ctx.current_ds = ds
	return nil
}
