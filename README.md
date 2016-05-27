# goKeyValueStore
Making a K/V store to learn more about how they, files, file systems and B-trees work

## Release 1 - initial design, no B-tree, pretty crappy design but well tested, do not use!

- goKeyValueStore.Open(filename string) (\*kVStore)
- kVStore.Get(key []byte) (value []byte)
- kVStore.Put(key []byte, key []byte) (error)
- kVStore.Close() (error)

#### A solid start with very little going on internally, now this base is finished I can continue with memory zeroing, re-insertion of keys and faster Gets.

Keys and values are stored in one file using ASCII delimiters. On Open(filename) the file is checked that it contains some data, or a new file is created if non-existant on the file system. Close() will safely close the db file.

On a Get(key) the whole DB is scanned to give a key's position, the value is then read from that position.

On a Put(key, value) the k/v pair is appended to the end of the file. if the key is a repeat then you will always get the first value added.
