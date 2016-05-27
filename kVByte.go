package goKeyValueStore

import (
	"bytes"
	"errors"
)

var keyValueDelimiter byte = 30 // this shouldn't be in this file
var valueDelimiter byte = 31    // this should!

// splitKVByte takes a key and a value byte slice, finds the value delimiter
// and returns the seperated key and value byte slices.
func splitKVByte(kVByte []byte) (key []byte, value []byte, err error) {
	// index the delimiter byte
	valueDelimiterIndex := bytes.IndexByte(kVByte, valueDelimiter)

	// index = -1 when valueDelimiter is not present
	if valueDelimiterIndex == -1 {
		return nil, nil, errors.New("cannot find value delimiter")
	}

	return kVByte[:valueDelimiterIndex], kVByte[valueDelimiterIndex+1:], nil
}

// makeKVByte takes key and value byte slices and returns them as a single byte
// slice with the value delimiter in the middle.
func makeKVByte(key []byte, value []byte) []byte {
	keyAndDelimiter := append(key, valueDelimiter)
	return append(keyAndDelimiter, value...)
}
