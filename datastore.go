package main

import (
	"fmt"
	"os"
)

type Datastore struct {
	store map[string]string
}

func NewDatastore() *Datastore {
	ds := new(Datastore)
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

func (ds *Datastore) Persist(file_path string) {
	f, err := os.Create(file_path)

	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		for k, v := range ds.store {
			str := k + "=" + v + ";"
			if _, err2 := f.Write([]byte(str)); err2 != nil {
				fmt.Errorf(err.Error())
			}
		}
	}
	defer f.Close()
}
