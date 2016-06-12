package goKeyValueStore

import (
	"crypto/sha1"
)

func hashKey(key []byte) (hash string) {
	// start with a new hash.
	h := sha1.New()

	// write key into sha1 hasher
	h.Write(key)

	// This gets the finalized hash result as a [20]byte slice
	return string(h.Sum(nil))
}
