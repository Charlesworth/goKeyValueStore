package goKeyValueStore

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
)

type kVStore struct {
	file           *os.File
	keyLocationMap map[[10]byte]int64
}

var nullDelimiter byte = 31 // used to show zeroed memory

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
		// first time file setup
		// make empty map[key]offset
	} else {
		err = validateDBFile(file)
		// memory map to map[key]offset
	}

	if err != nil {
		file.Close()
		return nil, err
	}

	keyLocationMap := make(map[[10]byte]int64)
	return &kVStore{file, keyLocationMap}, err
}

func zeroMemory(memOffset int64, bytesToZero int) error {
	// zero the memory
	// place the nullDelimiter in mem location to indicate zeroing
	return nil
}

func validateDBFile(file *os.File) error {
	// check that file size is not 0
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		// if file size is zero, invalid file
		return errors.New("invalid db file: file has not been initialised")
	}

	// add additional file validation here

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
		kVByte := fileScanner.Bytes()
		_, value, err := splitKVByte(kVByte)
		// this bit needs improved error handling
		if err == nil {
			return value
		}
	}

	return nil
}

func keyValueSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// if at EOF, as the key vals have a trailing delimiter, return nothing
	if atEOF || len(data) == 0 {
		return 0, nil, nil
	}

	// if the keyValueDelimiter is present, return data left set to it
	if i := bytes.IndexByte(data, keyValueDelimiter); i >= 0 {
		return i + 1, data[0:i], nil
	}

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
