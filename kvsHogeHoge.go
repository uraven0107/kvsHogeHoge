package main

import (
	"errors"
	"fmt"
)

type Context struct {
	ds_manager *DatastoreManager
	current_ds *Datastore
}

func (ctx *Context) Current_ds_name() string {
	if ctx.current_ds == nil {
		return "none"
	}
	return ctx.current_ds.name
}

func (ctx *Context) Switch_ds(ds_name string) error {
	ds := ctx.ds_manager.getDatastore(ds_name)
	if ds == nil {
		return errors.New(fmt.Sprintf("Datastore named '%v' dosen't exist", ds_name))
	}
	ctx.current_ds = ds
	return nil
}
