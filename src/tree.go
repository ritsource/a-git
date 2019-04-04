package src

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
)

// GitTreeLeaf ...
type GitTreeLeaf struct {
	Mode  string
	Fpath string
	Sha   string
}

// GitTree ...
type GitTree []GitTreeLeaf

// ReadTree - Reads tree object data and returns an array of records ([]GitTreeLeaf)
func ReadTree(treedata []byte) GitTree {
	pos := 1             // "pos" is used to track the position of each treeleaf's (record's) start
	counter := 0         // Counter is used to track the index of treeleaf (each record)
	max := len(treedata) // End of treelear data

	var tree GitTree // to hold treeleafs

	// As long as pos < max (as long as parse hasn't completed)
	for pos < max {
		// To parse each record, "counter" keeps track of index
		treeleaf := ParseTreeLeaf(treedata, counter, &pos)
		counter++
		// fmt.Printf("Val: %+v\n", treeleaf)
		tree = append(tree, treeleaf)
	}

	return tree
}

// ParseTreeLeaf - Every tree contains one or more records
// This function parses and returns GitTreeLeaf for each Record
func ParseTreeLeaf(raw []byte, counter int, pos *int) GitTreeLeaf {
	// fmt.Println(string(raw))
	// 0x00 [mode] space [path] 0x00 [sha-1] => Structure of a Tree Leaf (One Record in Tree Object)
	// in this function "x" = "space index", and "y" = "null index (0x00)"
	x := IndexBytesByIndex(raw, byte(' '), counter) // Finding the ' ' (according to counter)
	y := IndexBytesByIndex(raw, 0x00, counter+1)

	// Find "Mode" from start to byte(' ')
	mode := raw[*pos:x]

	// Find "Path" from first byte(' ') after pos (start) to byte(0x00)
	fpath := raw[x+1 : y] // fpath not to confuse with path library

	// Decode "Sha" from 20 byte Binary (Sha is encoded in binary in Git Tree Object)
	shaStr := hex.EncodeToString(raw[y+1 : y+21])

	// Setting posiition to the end of encoded sha, (end of current record)
	*pos = y + 21

	return GitTreeLeaf{
		Mode:  string(mode),
		Fpath: string(fpath),
		Sha:   shaStr,
	}
}

// IndexBytesByIndex - Returns index of n-th byte match in a []byte
// for axample IndexBytesByIndex(raw, byte' ', 2) will return
// the index of 3ed byte(' ') in the buffer
func IndexBytesByIndex(b []byte, k byte, i int) int {
	// A counter to track index of byte match
	c := 0

	// Finding Index
	return bytes.IndexFunc(b, func(r rune) bool {
		if r == rune(k) {
			c++
		}
		if c == i+1 {
			return true
		}

		return false
	})
}

// ReadSingleRecord ...
func ReadSingleRecord(raw []byte, pos int) (string, int) {
	iNull := IndexBytesByIndex(raw, 0x00, 1)
	// fmt.Println("Data", string(raw[iNull:]))

	h := sha1.New()
	h.Write(raw[iNull:])
	shaStr := hex.EncodeToString(h.Sum(nil))
	// fmt.Println("sha:", shaStr)

	return shaStr, iNull + 21
}
