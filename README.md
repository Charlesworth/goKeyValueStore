# goKeyValueStore
Making a K/V store to learn more about how they, files, file systems and B-trees work

- goKeyValueStore.Open(filename string) (\*kVStore)
- kVStore.Get(key []byte) (value []byte)
- kVStore.Put(key []byte, key []byte) (error)
- kVStore.Close() (error)

#### 2nd release goals

- Put() supports key reinsertion
- memory-map against keys to speed up Get() search speed
- reuse k/v location on key re-insertion or zero it if too small
- include some higher level behavioral testing

#### 3rd release goals

- record zeroed memory locations for use with new keys
- use a BTree for the key/memory location map
- expose additional functions to API, including len and position.
