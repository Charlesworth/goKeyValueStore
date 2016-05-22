package charDB

import (
	"os"
  "bufio"
	"bytes"
	"strings"
	"errors"
)

type charDB struct {
	file *os.File
}

func Open(dbFileName string) (*charDB, error) {
	//check if db file exists already
	newDBFile := false
	_, err := os.Stat(dbFileName)
	if err != nil && strings.Contains(err.Error(), "cannot find the file"){
		newDBFile = true
	}

	//open or create the db file
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		file.Close()
		return nil, err
	}

	// if the db file already exists then memory map the keys to their offsets
	if !newDBFile {
		err = validateDBFile(file)
		//map in here
	}

	if err != nil {
		file.Close()
		return nil, err
	}
	return &charDB{file}, err
}

func validateDBFile(file *os.File) error {
	// check that
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return errors.New("invalid db file: file is empty")
	}

	return nil
}

func (db *charDB) Close() error {
  return db.file.Close()
}

func (db *charDB) Put(key []byte, value []byte) error {
	fileInfo, _ := db.file.Stat()
	writeToPosition := fileInfo.Size()
	newline := []byte("\n")
	data := append(key, newline...)
	data = append(data, value...)
	data = append(data, newline...)

	db.file.WriteAt(data, writeToPosition)
	return nil
}

func (db *charDB) Get(key []byte) []byte {
  fileScanner := bufio.NewScanner(db.file)

	keyFound := findKey(fileScanner, key)
	if keyFound {
		fileScanner.Scan()
		return fileScanner.Bytes()
	}

	return []byte{}
}

func findKey(fileScanner *bufio.Scanner, key []byte) bool {
	for fileScanner.Scan() {
		if (bytes.Compare(fileScanner.Bytes(), key) == 0) {
			return true
    }
		fileScanner.Scan()
  }
	return false
}
