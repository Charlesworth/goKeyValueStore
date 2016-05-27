package goKeyValueStore

import (
	"bytes"
	"os"
	"testing"
)

// add any test files produced to this array for later cleanup
var testFiles []string

func TestClose(t *testing.T) {
	testDBFileName := "TestClose"
	testFiles = append(testFiles, testDBFileName)
	file := makeEmptyFile(testDBFileName, t)

	testDB := &kVStore{file}

	err := testDB.Close()
	if err != nil {
		t.Log("Close() should os.File.Close() without error")
		t.Error("Error: ", err)
	}

	_, err = file.Write([]byte("some data"))
	if err == nil {
		t.Log("Once Closed, we should expect an error if we try to write")
		t.Error("Error: no error returned from apemting to write to closed file")
	}
}

func TestOpen_OnNewFile(t *testing.T) {
	testDBFileName := "TestOpen_OnNewFile"
	testFiles = append(testFiles, testDBFileName)

	testDB, err := Open(testDBFileName)

	if testDB == nil {
		t.Log("Open() should return a non nil kVStore pointer")
		t.Error("Error: Open() did not return a *kVStore, instead returned nil")
	}

	if err != nil {
		t.Log("Open() should return err == nil when Open() makes a new file")
		t.Error("Error: ", err)
	}

	_, err = os.Stat(testDBFileName)
	if err != nil {
		t.Log("Open() should make a new file if the fileName is not present")
		t.Error("Error:", err)
	}

	testDB.Close()
}

func TestOpen_OnExistingFile_CannotValidateFile(t *testing.T) {
	testDBFileName := "TestOpen_OnExistingFile_CannotValidateFile"
	testFiles = append(testFiles, testDBFileName)

	// make a pre existing file for this test case
	makeEmptyFile(testDBFileName, t).Close()

	testDB, err := Open(testDBFileName)

	if testDB != nil {
		t.Log("Open() should return a nil kVStore pointer on a invalid DB file")
		t.Error("Error: Open() on a invalid db file should return nil")
	}

	if err == nil {
		t.Error("Error: Open() should return err when opening an invalid db file")
	}

	_, err = os.Stat(testDBFileName)
	if err != nil {
		t.Log("Open() when the file already exists shouldn't delete the file")
		t.Error("Error:", err)
	}
}

func TestPut_OnNewKey(t *testing.T) {
	testDB := newTestDB("TestPut_OnNewKey", t)

	key, value := []byte("hello"), []byte("goodbyte")
	err := testDB.Put(key, value)
	if err != nil {
		t.Log(".Put() returned with an error")
		t.Error(err)
	}

	testDB.Close()
}

func TestPut_OnExistingKey(t *testing.T) {
	testDB := newTestDB("TestPut_OnExistingKey", t)

	key, oldValue, newValue := []byte("hello"), []byte("goodbyte"), []byte("byte me")
	testDB.Put(key, oldValue)
	err := testDB.Put(key, newValue)
	if err != nil {
		t.Log(".Put() returned with an error")
		t.Error(err)
	}

	testDB.Close()
}

func TestPut_OnMultipleDifferentKeys(t *testing.T) {
	testDB := newTestDB("TestPut_OnMultipleDifferentKeys", t)

	key, value := []byte("hello"), []byte("goodbyte")
	testDB.Put(key, value)

	differentKey, differentValue := []byte("afternoon"), []byte("good evening")
	err := testDB.Put(differentKey, differentValue)
	if err != nil {
		t.Log(".Put() returned with an error")
		t.Error(err)
	}

	testDB.Close()
}

func TestOpen_OnExistingFile(t *testing.T) {
	testDBFileName := "TestOpen_OnExistingFile"
	testFiles = append(testFiles, testDBFileName)

	// makeEmptyFile(testDBFileName, t).Close()

	// make a pre existing file for this test case
	testDB, err := Open(testDBFileName)
	if err != nil {
		t.Error("Test Error: unable to make the pre-existing test file for this test")
	}
	testDB.Put([]byte("value"), []byte("value"))
	testDB.Close()

	// reopen the file, which should now be a valid db file after initialisation
	testDB, err = Open(testDBFileName)

	if testDB == nil {
		t.Log("Open() should return a non nil kVStore pointer")
		t.Error("Error: Open() did not return a *kVStore, instead returned nil")
	}

	if err != nil {
		t.Log("Open() should return err == nil")
		t.Error("Error: ", err)
	}

	_, err = os.Stat(testDBFileName)
	if err != nil {
		t.Log("Open() when the file already exists shouldn't delete the file")
		t.Error("Error:", err)
	}

	testDB.Close()
}

func TestGet_OnExistingKey(t *testing.T) {
	testDB := newTestDB("TestGet_OnExistingKey", t)
	key, inputValue := []byte("hello"), []byte("goodbyte")
	testDB.Put(key, inputValue)

	outputValue := testDB.Get(key)
	if !(bytes.Compare(inputValue, outputValue) == 0) {
		t.Log("Get() on a existing key did not return matching value")
		t.Error("Error: input value [", inputValue, "] output value: [", outputValue, "]")
	}

	testDB.Close()
}

func TestGet_OnNonExistantKey(t *testing.T) {
	testDB := newTestDB("TestGet_OnNonExistantKey", t)
	key, expectedOutputValue := []byte("hello"), []byte{}

	outputValue := testDB.Get(key)
	if !(bytes.Compare(expectedOutputValue, outputValue) == 0) {
		t.Log("Get() on a non-existant key did not return matching value")
		t.Error("Error: input value [", expectedOutputValue, "] output value: [", outputValue, "]")
	}

	testDB.Close()
}

func TestGet_OnMultipleExistingKeys(t *testing.T) {
	testDB := newTestDB("TestGet_OnMultipleExistingKeys", t)
	falseKey, falseValue := []byte("hi"), []byte("bye")
	testDB.Put(falseKey, falseValue)

	key, inputValue := []byte("hello"), []byte("goodbyte")
	testDB.Put(key, inputValue)

	outputValue := testDB.Get(key)
	if !(bytes.Compare(inputValue, outputValue) == 0) {
		t.Log("Get() on a existing key did not return matching value")
		t.Error("Error: input value [", inputValue, "] output value: [", outputValue, "]")
	}

	testDB.Close()
}

func TestCleanUp(t *testing.T) {
	for _, file := range testFiles {
		err := os.Remove(file)
		if err != nil {
			t.Log("unable to clean up file:", file)
			t.Error(err)
		}
	}
}

func newTestDB(testDBName string, t *testing.T) *kVStore {
	testFiles = append(testFiles, testDBName)

	testDB, err := Open(testDBName)
	if err != nil {
		t.Fatal("unable to open testDB ", testDBName, "with error: ", err)
	}

	return testDB
}

func makeEmptyFile(fileName string, t *testing.T) *os.File {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0777)
	if err != nil {
		t.Fatal("makeEmptyFile testing helper function failed:", err)
	}

	return file
}
