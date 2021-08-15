package main

type DatastoreManager struct {
	ds_list []*Datastore
}

func (dm *DatastoreManager) Persist() error {
	persisted := ""
	for _, ds := range dm.ds_list {
		persisted = persisted + ds.Persisted()
	}
	err := WriteDBFile(&persisted)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DatastoreManager) Restore() error {
	file_contents, err := ReadDBFile()
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
