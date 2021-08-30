package main

import (
	"errors"
	"fmt"
)

type Application struct {
	ds_manager *DatastoreManager
	current_ds *Datastore
}

func (app *Application) Current_ds_name() string {
	if app.current_ds == nil {
		return "none"
	}
	return app.current_ds.name
}

func (app *Application) Switch_ds(ds_name string) error {
	ds := app.ds_manager.getDatastore(ds_name)
	if ds == nil {
		return errors.New(fmt.Sprintf("Datastore named '%v' dosen't exist", ds_name))
	}
	app.current_ds = ds
	return nil
}
