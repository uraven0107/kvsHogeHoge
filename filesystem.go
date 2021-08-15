package main

import (
	"io/ioutil"
	"os"
)

const db_file_dir = "test_data"
const db_file = "test.hogedb"
const db_file_path = db_file_dir + string(os.PathSeparator) + db_file

func ReadDBFile() (string, error) {
	contents, err := ioutil.ReadFile(db_file_path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func WriteDBFile(str *string) error {
	fp, err := os.Create(db_file_path)
	if err != nil {
		return err
	}

	bytes := []byte(*str)
	if _, err := fp.Write(bytes); err != nil {
		return err
	}

	return nil
}
