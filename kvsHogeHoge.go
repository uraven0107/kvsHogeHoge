package main

type Application struct {
	ds_manager *DatastoreManager
	current_ds *Datastore
}

func (app Application) Current_name() string {
	return app.current_ds.name
}
