package charDB

import (
	"os"
  "bufio"
	"bytes"
)

type charDB struct {
	file *os.File
}

func Open(dbFileName string) (*charDB, error) {
	file, err := os.OpenFile(dbFileName, os.O_CREATE, 0777)
	return &charDB{file}, err
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
  for fileScanner.Scan() {
		if (bytes.Compare(fileScanner.Bytes(), key) == 0) {
    	fileScanner.Scan()
			return fileScanner.Bytes()
    }
		fileScanner.Scan()
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
