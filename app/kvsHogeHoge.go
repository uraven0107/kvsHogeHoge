package core

import (
	"fmt"

	"github.com/uraven0107/kvsHogeHoge/core"
)

type Application struct {
	ds_manager *core.DatastoreManager
}

func Run(config *Configure) (*Application, error) {
	ds_manager, err := core.NewDatastoreManager(config.file_path)
	if err != nil {
		return nil, err
	}

	err = ds_manager.Restore()
	if err != nil {
		return nil, err
	}

	return &Application{
		ds_manager: ds_manager,
	}, nil
}

type AppContext struct {
	app        *Application
	current_ds *core.Datastore
}

func NewAppContext(app *Application) *AppContext {
	return &AppContext{
		app:        app,
		current_ds: nil,
	}
}

func (ctx *AppContext) Current_ds_name() string {
	if ctx.current_ds == nil {
		return "none"
	}
	return ctx.current_ds.GetName()
}

func (ctx *AppContext) Switch_ds(ds_name string) error {
	ds := ctx.app.ds_manager.GetDatastore(ds_name)
	if ds == nil {
		return fmt.Errorf("Datastore named '%v' dosen't exist", ds_name)
	}
	ctx.current_ds = ds
	return nil
}
