# goKeyValueStore
Making a K/V store to learn more about how they, files, file systems and B-trees work

## Release 1 target - initial design, no B-tree

Keys and values are stored in one file with a common separation value. On startup the file is memory-mapped giving a key/offset map.

On a read the map is read to give a key's offset position, the value is then read from that position.

On a write, the map is first checked for the pre-existence of the key. In the case it is a new key, the key and value will be appended to end of the file and key/offset added to the map. If the key is present in the map, that key/value data location will be compared against the new k/v size. If the new k/v will fit in the existing location it will be placed there, else the k/v will be appended to the end of the file, the previous data location zeroed and its map entry updated to reflect the change.

As a small add on I would like to include memory of zeroed space, so that new k/v placements can examine the zeroed space and if suitable fill it.
