package goKeyValueStore

import (
	"bytes"
	"testing"
)

func TestKVByte_splitKVByte_validKVByte(t *testing.T) {
	var inputKey, inputValue = []byte("testKey"), []byte("testValue")
	var delimiter byte
	delimiter = 31

	// make a valid kVByte
	keyDelim := append(inputKey, delimiter)
	kVByte := append(keyDelim, inputValue...)

	outputKey, outputValue, err := splitKVByte(kVByte)
	if err != nil {
		t.Error("spliting a valid kVByte returned error:", err)
	}

	if !bytes.Equal(outputKey, inputKey) {
		t.Error("key returned incorrectly, input: ", inputKey, ", output: ", outputKey)
	}

	if !bytes.Equal(outputValue, inputValue) {
		t.Error("value returned incorrectly, input: ", inputValue, ", output: ", outputValue)
	}
}

func TestKVByte_splitKVByte_invalidKVByte(t *testing.T) {
	var inputKey, inputValue = []byte("testKey"), []byte("testValue")

	// make a invalid kVByte without a delimiter
	kVByte := append(inputKey, inputValue...)

	_, _, err := splitKVByte(kVByte)
	if err == nil {
		t.Error("spliting a invalid kVByte didn't return error")
	}
}

func TestKVByte_makeKVByte(t *testing.T) {
	var inputKey, inputValue = []byte("testKey"), []byte("testValue")
	var delimiter byte
	delimiter = 31

	// make a valid kVByte
	keyDelim := append(inputKey, delimiter)
	validKVByte := append(keyDelim, inputValue...)

	testKVByte := makeKVByte(inputKey, inputValue)

	if !bytes.Equal(validKVByte, testKVByte) {
		t.Error("kVByte returned incorrectly, valid output: ", validKVByte, ", test output: ", testKVByte)
	}
}
