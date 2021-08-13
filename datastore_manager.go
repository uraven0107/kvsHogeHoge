package main

import "os"

type DatastoreManager struct {
	file_path string
	ds_list   []*Datastore
}

func (dm DatastoreManager) Persist() error {
	f, err := os.Create(dm.file_path)
	if err != nil {
		return err
	} else {
		persisted := ""
		for _, ds := range dm.ds_list {
			persisted = persisted + ds.Persisted()
		}
		if _, err := f.Write([]byte(persisted)); err != nil {
			return err
		}
	}
	return nil
}
