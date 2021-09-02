package util

import (
	"io/ioutil"
	"os"
)

func ReadFile(file_path string) (string, error) {
	contents, err := ioutil.ReadFile(file_path)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func WriteFile(str *string, file_path string) error {
	fp, err := os.Create(file_path)
	if err != nil {
		return err
	}

	bytes := []byte(*str)
	if _, err := fp.Write(bytes); err != nil {
		return err
	}

	return nil
}
