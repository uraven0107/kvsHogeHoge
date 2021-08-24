package main

type DatastoreManager struct {
	db_file_path string
	ds_list []*Datastore
}

func NewDatastoreManager(db_file_path string) *DatastoreManager {
	dm := &DatastoreManager{}
	dm.db_file_path = db_file_path
	dm.ds_list = []*Datastore{}
	return dm
}

func (dm *DatastoreManager) getDatastore(name string) *Datastore {
	if dm.ds_list == nil {
		panic("wooooooop!! DatastoreManager.ds_list doesn't initialized!")
	}
	for _, ds := range dm.ds_list {
		if ds.name == name {
			return ds
		}
	}
	return nil
}

func (dm *DatastoreManager) Persist() error {
	persisted := ""
	for _, ds := range dm.ds_list {
		persisted = persisted + ds.Persisted()
	}
	
	err := WriteFile(&persisted, dm.db_file_path)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DatastoreManager) Restore() error {
	file_contents, err := ReadFile(dm.db_file_path)
	if err != nil {
		return err
	}
	tokenizer, err := NewTokenizer(file_contents)
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
