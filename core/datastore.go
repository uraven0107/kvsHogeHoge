package core

type Datastore struct {
	name  string
	store map[string]string
}

func NewDatastore(name string) *Datastore {
	ds := new(Datastore)
	ds.name = name
	ds.store = make(map[string]string)
	return ds
}

func (ds *Datastore) Write(key string, value string) {
	ds.store[key] = value
}

func (ds *Datastore) Read(key string) string {
	return ds.store[key]
}

func (ds *Datastore) Size() int {
	return len(ds.store)
}

func (ds *Datastore) Delete(key string) {
	delete(ds.store, key)
}

func (ds *Datastore) Persisted() string {
	str := ""
	for k, v := range ds.store {
		str = str + k + "=" + v + ";"
	}
	str = ds.name + "={" + str + "};"
	return str
}
