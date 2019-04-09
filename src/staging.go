package src

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"path"
)

// StageDir - Creates trees and blob objects recursively for all sub directories
// and files of a given dir
func StageDir(dir string) (GitTree, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return GitTree{}, nil
	}

	// var fSlice []string
	tree := GitTree{}

	for _, f := range files {
		// "%06o", "o" - in 8 base because of this (visit the url)
		// https://github.com/src-d/go-git/blob/v4.10.0/plumbing/filemode/filemode.go#L141
		modeStr := fmt.Sprintf("%06o", f.Mode()) // Filemode

		// Check if directory of file
		if f.IsDir() {
			// Building Tree for Sub-Directory recursively, returns GitTree
			subTree, err := StageDir(path.Join(dir, f.Name()))
			if err != nil {
				return GitTree{}, nil
			}

			// Writing Tree Object
			shaStr, err := subTree.Write("git")
			if err != nil {
				return GitTree{}, nil
			}

			// Decode shaStr
			sha, err := hex.DecodeString(shaStr)
			if err != nil {
				return GitTree{}, nil
			}

			// Append to tree (Slice of leafs)
			tree = append(tree, GitTreeLeaf{
				Mode:  modeStr,
				Fpath: f.Name(),
				Sha:   sha,
			})

		} else {
			// If not directory, if file

			// Read data from source
			data, err := ioutil.ReadFile(path.Join(dir, f.Name()))
			if err != nil {
				return GitTree{}, nil
			}

			// Initialize a git object
			blob := GitObject{
				Kind: "blob",
				Data: data,
			}

			// Writing Blob Object
			shaStr, err := blob.Write("git")
			if err != nil {
				return GitTree{}, nil
			}

			// Decoding shaStr
			sha, err := hex.DecodeString(shaStr)
			if err != nil {
				log.Fatal(err)
			}

			// Append the tree leaf to tree
			tree = append(tree, GitTreeLeaf{
				Mode:  modeStr,
				Fpath: f.Name(),
				Sha:   sha,
			})
		}
	}

	// Return the tree (GitTree)
	return tree, nil
}
