package core

import (
	"errors"

	"github.com/uraven0107/kvsHogeHoge/process"
	"github.com/uraven0107/kvsHogeHoge/util"
)

type DatastoreManager struct {
	db_file_path string
	ds_list      []*Datastore
}

func NewDatastoreManager(file_path string) (*DatastoreManager, error) {
	if file_path == "" {
		return nil, errors.New("Configure.file_path is blank")
	}

	dm := &DatastoreManager{}
	dm.db_file_path = file_path
	dm.ds_list = []*Datastore{}
	return dm, nil
}

func (dm *DatastoreManager) GetDatastore(name string) *Datastore {
	if dm.ds_list == nil {
		panic("wooooooop!! DatastoreManager.ds_list doesn't initialized!")
	}
	for _, ds := range dm.ds_list {
		if ds == nil {
			continue
		}
		if ds.name == name {
			return ds
		}
	}
	return nil
}

func (dm *DatastoreManager) AddDatastore(ds *Datastore) {
	dm.ds_list = append(dm.ds_list, ds)
}

func (dm *DatastoreManager) Persist() error {
	persisted := ""
	for _, ds := range dm.ds_list {
		persisted = persisted + ds.Persisted()
	}

	err := util.WriteFile(&persisted, dm.db_file_path)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DatastoreManager) Restore() error {
	file_contents, err := util.ReadFile(dm.db_file_path)
	if err != nil {
		return err
	}
	tokenizer, err := process.NewTokenizer(process.Type_DS, file_contents)
	if err != nil {
		return err
	}
	parser := process.NewParser(tokenizer)
	ds_sources, err := parser.Expr()

	if err != nil {
		return err
	}

	for _, ds_source := range ds_sources {
		ds := NewDatastore(ds_source.Name)
		for k, v := range ds_source.K_V_map {
			ds.Write(k, v)
		}
		dm.ds_list = append(dm.ds_list, ds)
	}

	return nil
}
