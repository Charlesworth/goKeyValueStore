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

type kVStore struct {
	file *os.File
}

/*
ideal struct at the end
type kVStore struct {
	dataFile 	*os.File
	freedMem 	(ordered list of size and location pairs, sortable by size)
	keyMap		BTreeMap of keys to locations
}

need to find a way to give variable key sizes or split the key so there
cannot be any collisions
*/

func Open(dbFileName string) (*kVStore, error) {
	// check if db file exists already
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
	return &kVStore{file}, err
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

func (db *kVStore) Close() error {
	return db.file.Close()
}

func (db *kVStore) Put(key []byte, value []byte) error {
	currentDBValue := db.Get(key)
	if currentDBValue == nil {
		// zero current key memory location
	}

	fileInfo, _ := db.file.Stat()
	writeToPosition := fileInfo.Size()
	data := append(key, valueDelimiter)
	data = append(data, value...)
	data = append(data, keyValueDelimiter)

	db.file.WriteAt(data, writeToPosition)
	return nil
}

func (db *kVStore) Get(key []byte) []byte {

	fileScanner := bufio.NewScanner(db.file)
	fileScanner.Split(keyValueSplitFunc)

	keyFound := findKey(fileScanner, key)
	if keyFound {
		kV := fileScanner.Bytes()
		value := getValueFromKVBytes(kV)
		return value
	}

	return nil
}

func getValueFromKVBytes(kV []byte) []byte {
	i := bytes.IndexByte(kV, valueDelimiter)

	//returns -1 when valueDelimiter is not present, need an error in here
	if i == -1 {
		return []byte{}
	}

	//use i here to split and give the value
	return kV[i+1:]
}

func keyValueSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// if atEOF && len(data) == 0 {

	// if at EOF, as the key vals have a trailing delimiter, return nothing
	if atEOF || len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, keyValueDelimiter); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// if atEOF {
	// 	return
	// }

	return
}

func findKey(fileScanner *bufio.Scanner, key []byte) bool {
	keyAndDelimiter := append(key, valueDelimiter)
	for fileScanner.Scan() {
		if bytes.Contains(fileScanner.Bytes(), keyAndDelimiter) {
			return true
		}
	}
	return false
}
