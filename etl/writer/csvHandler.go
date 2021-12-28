package writer

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
)

func WriteMetaData(location string, id string, columns []string, values []string) error {
	var pathToFile = filepath.Join("data", location)
	// create directory
	err := createDirectory(pathToFile)
	if err != nil {
		return err
	}

	// check if file exists and create it
	pathToFile = filepath.Join(pathToFile, id)
	if _, err := os.Stat(pathToFile); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(pathToFile)
		var csvWriter = csv.NewWriter(f)

		err = csvWriter.Write(columns)
		if err != nil {
			return err
		}
		err = csvWriter.Write(values)
		if err != nil {
			return err
		}

		csvWriter.Flush()
		defer f.Close()
	}
	// file exists
	return nil
}

func WriteValueData(location string, filename string, columns []string, values [][]string) error {
	var pathToFile = filepath.Join("data", location)
	// create directory
	err := createDirectory(pathToFile)
	if err != nil {
		return err
	}

	pathToFile = filepath.Join(pathToFile, filename)

	// check if file exists and create it
	f, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	csvWriter := csv.NewWriter(f)

	if res, err := os.Stat(pathToFile); err == nil && res.Size() == 0 {
		err = csvWriter.Write(columns)
		if err != nil {
			return err
		}
	}

	err = csvWriter.WriteAll(values)
	if err != nil {
		return err
	}

	csvWriter.Flush()
	defer f.Close()
	return nil
}

func createDirectory(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return os.MkdirAll(folder, 0755)
	}
	return nil
}
