package goKeyValueStore

import (
	"testing"
)

func TestHashKey_validKey(t *testing.T) {
	testKey := []byte("abcd1234")

	hash := hashKey(testKey)

	if len(hash) != 20 {
		t.Error("hashKey should return a hash of length 20, ",
			"hashKey returned length:", len(hash))
	}
}

func TestHashKey_zeroLengthKey(t *testing.T) {
	testZeroLenghtKey := []byte{}

	hash := hashKey(testZeroLenghtKey)

	if len(hash) != 20 {
		t.Error("hashKey should return a hash of length 20, ",
			"hashKey returned length:", len(hash))
	}
}
