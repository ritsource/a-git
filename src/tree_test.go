package src_test

import (
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ritwik310/a-git/src"
)

func TestTreeWriteAndRead(t *testing.T) {
	// "../tests/fixtures/tree1"

	// Fixture tree path
	fixturepath := "../tests/fixtures/tree1"

	// Base Object (Read from fixtures)
	bObj, err := src.ReadObject(fixturepath)
	if err != nil {
		t.Error(err)
	}

	// Base Tree
	bTree := src.ParseTree(bObj.Data)

	// Writing new Tree Object
	// nFilePath = new file path
	shaStr, err := bTree.Write("git")

	// Reading newly created object
	rObject, err := src.ReadObject(path.Join("git", "objects", shaStr[:2], shaStr[2:])) // "../tests/fixtures/tree1"
	if err != nil {
		t.Error(filepath.Abs(path.Join("git", "objects", shaStr[:2], shaStr[2:])))
	}

	// Extracting Tree from new tree object
	rTree := src.ParseTree(rObject.Data)

	// Comparing read tree with base tree
	if !reflect.DeepEqual(rTree, bTree) {
		t.Error("Not Match!", "\n\n", rTree, "\n\n", bTree)
	}
}
