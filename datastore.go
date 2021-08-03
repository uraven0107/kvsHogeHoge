package main

type Datastore struct {
	store map[string]Value
}

func NewDatastore() *Datastore {
	ds := new(Datastore)
	ds.store = make(map[string]Value)
	return ds
}

func (ds *Datastore) Write(key string, value Value) {
	ds.store[key] = value
}

func (ds *Datastore) Read(key string) Value {
	return ds.store[key]
}

func (ds *Datastore) Size() int {
	return len(ds.store)
}

func (ds *Datastore) Delete(key string) {
	delete(ds.store, key)
}
