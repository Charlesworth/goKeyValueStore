package goKeyValueStore

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
)

var keyValueDelimiter byte = 30
var valueDelimiter byte = 31

type charDB struct {
	file *os.File
}

func Open(dbFileName string) (*charDB, error) {
	// check if db file exists already

	// for whatever reason this is breaking
	// newDB := filePresent(dbFileName)

	newDB := false
	_, err := os.Stat(dbFileName)
	if err != nil && strings.Contains(err.Error(), "cannot find the file") {
		newDB = true
	}

	//open or create the db file
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		file.Close()
		return nil, err
	}

	// if the db file already exists then validate and memory map the
	// keys to their offsets
	if newDB {
		// setup file
		// make empty map
	} else {
		err = validateDBFile(file)
		// map in here
	}

	if err != nil {
		file.Close()
		return nil, err
	}
	return &charDB{file}, err
}

// BROCKEN, DON'T KNOW WHY
// filePresent checks the pwd and returns if file 'filename' is present
func filePresent(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil && strings.Contains(err.Error(), "cannot find the file") {
		return false
	}
	return true
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

	// add data validation here

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
		if bytes.Compare(fileScanner.Bytes(), key) == 0 {
			return true
		}
		fileScanner.Scan()
	}
	return false
}
