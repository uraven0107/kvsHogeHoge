package main

import (
	"io/ioutil"
	"os"
)

type DatastoreManager struct {
	file_path string
	ds_list   []*Datastore
}

func (dm *DatastoreManager) Persist() error {
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

func (dm *DatastoreManager) Restore() error {
	source, err := ioutil.ReadFile(dm.file_path)
	if err != nil {
		return err
	}
	tokenizer, err := NewTokenizer(string(source))
	if err != nil {
		return err
	}
	parser := NewParser(tokenizer)
	ds_sources, err := parser.Expr()

	if err != nil {
		return err
	}

	for _, ds_source := range ds_sources {
		ds := NewDatastore(ds_source.name)
		for k, v := range ds_source.k_v_map {
			ds.Write(k, v)
		}
		dm.ds_list = append(dm.ds_list, ds)
	}

	return nil
}
