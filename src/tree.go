package src

import (
	"bytes"
)

// GitTreeLeaf ...
type GitTreeLeaf struct {
	Mode  string
	Fpath string
	Sha   []byte
}

// GitTree ...
type GitTree []GitTreeLeaf

// Write - Writes a GitObject from a GitTree Struct
func (tree GitTree) Write(gitdir string) (string, error) {
	// Raw tree data (empty)
	raw := []byte{}

	// For each tree-leaf
	for _, t := range tree {
		// Building up data as [mode] space [path] 0x00 [sha-1]
		raw = bytes.Join([][]byte{raw, []byte(t.Mode)}, []byte(""))
		raw = bytes.Join([][]byte{raw, []byte(" ")}, []byte(""))
		raw = bytes.Join([][]byte{raw, []byte(t.Fpath)}, []byte(""))
		raw = bytes.Join([][]byte{raw, []byte{0x00}}, []byte(""))
		raw = bytes.Join([][]byte{raw, []byte(t.Sha)}, []byte(""))
	}

	// Writetable Object from tree (GitTree)
	wObj := GitObject{Kind: "tree", Data: raw}

	// Writing the Object
	nFilePath, err := wObj.Write("git")
	return nFilePath, err
}

// ReadTree - Reads tree object data and returns an array of records ([]GitTreeLeaf)
func ReadTree(treedata []byte) GitTree {
	pos := 0             // "pos" is used to track the position of each treeleaf's (record's) start
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
	// [mode] space [path] 0x00 [sha-1] => Structure of a Tree Leaf (One Record in Tree Object)
	// in this function "x" = "space index", and "y" = "null index (0x00)"
	x := IndexBytesByIndex(raw, byte(' '), counter) // Finding the ' ' (according to counter)
	y := IndexBytesByIndex(raw, 0x00, counter)

	// Find "Mode" from start to byte(' ')
	mode := raw[*pos:x]

	// Find "Path" from first byte(' ') after pos (start) to byte(0x00)
	fpath := raw[x+1 : y] // fpath not to confuse with path library

	// Decode "Sha" from 20 byte Binary (Sha is encoded in binary in Git Tree Object)
	sha := raw[y+1 : y+21]

	// Setting posiition to the end of encoded sha, (end of current record)
	*pos = y + 21

	return GitTreeLeaf{
		Mode:  string(mode),
		Fpath: string(fpath),
		Sha:   sha,
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
