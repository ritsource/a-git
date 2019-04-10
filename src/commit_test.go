package src_test

import (
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ritwik310/a-git/src"
)

func TestCommitWriteAndRead(t *testing.T) {
	// "../tests/fixtures/commit1"

	// Fixture commit object path
	fixturepath := "../tests/fixtures/commit1"

	// Base Object (Read from fixtures)
	bObj, err := src.ReadObject(fixturepath)
	if err != nil {
		t.Error(err)
	}

	// Base Commit
	bCommit := src.ParseCommit(bObj.Data)

	// Writing new Commit Object
	// nFilePath = new file path
	shaStr, err := bCommit.Write("git")

	// Reading newly created object
	rObject, err := src.ReadObject(path.Join("git", "objects", shaStr[:2], shaStr[2:])) // "../tests/fixtures/tree1"
	if err != nil {
		t.Error(filepath.Abs(path.Join("git", "objects", shaStr[:2], shaStr[2:])))
	}

	// Extracting Tree from new tree object
	rCommit := src.ParseCommit(rObject.Data)

	// Comparing read tree with base tree
	if !reflect.DeepEqual(rCommit, bCommit) {
		t.Error("Not Match!", "\n\n", rCommit, "\n\n", bCommit)
	}
}
