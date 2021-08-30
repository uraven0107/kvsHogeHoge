package main

type Application struct {
	ds_manager *DatastoreManager
	current_ds *Datastore
}

func (app Application) Current_ds_name() string {
	if app.current_ds == nil {
		return "none"
	}
	return app.current_ds.name
}
