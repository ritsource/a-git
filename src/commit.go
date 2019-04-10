package src

import "bytes"

// GitCommit ...
type GitCommit struct {
	Tree      []byte
	Parent    []byte
	Author    string
	Committer string
	Message   string
}

// Write - Writes a new Git Commit Object to the filesystem
// We are ignoring ---- PGP SIGNATURE ----
// It's just Tree, Parent (Optional), Author and Commiter (just Default), Message (Optional)
func (commit GitCommit) Write(gitdir string) (string, error) {
	// Raw commit data (empty)
	raw := []byte{}

	// Append tree to record
	raw = bytes.Join([][]byte{raw, []byte("tree"), []byte(" "), []byte(commit.Tree), []byte("\n")}, []byte(""))

	// Append Parent Record
	if len(commit.Parent) > 0 {
		raw = bytes.Join([][]byte{raw, []byte("parent"), []byte(" "), []byte(commit.Parent), []byte("\n")}, []byte(""))
	}
	// Append Author and Commit respectively
	raw = bytes.Join([][]byte{raw, []byte("author"), []byte(" "), []byte(commit.Author), []byte("\n")}, []byte(""))
	raw = bytes.Join([][]byte{raw, []byte("committer"), []byte(" "), []byte(commit.Committer), []byte("\n")}, []byte(""))

	// Append Message to the raw data
	if commit.Message != "" {
		raw = bytes.Join([][]byte{raw, []byte(commit.Message), []byte("\n")}, []byte(""))
	}

	// Writetable Object from commit (GitCOmmit)
	wObj := GitObject{Kind: "commit", Data: raw}

	// Writing the Object
	shaStr, err := wObj.Write("git")
	return shaStr, err
}

// ParseCommit - Parses GitObject.Data and Return GitCommit Object
func ParseCommit(commitdata []byte) GitCommit {
	// Initializing Commit Object
	var commit = GitCommit{}

	// Split each record (Split by []byte(\n))
	entry := bytes.Split(commitdata, []byte("\n"))

	// Itetate over the data
	for i, chunk := range entry {
		// If not Empty line (len(chunk) > 1)
		if len(chunk) > 1 {
			// Index of space (byte(' '))
			sp := bytes.IndexByte(chunk, byte(' '))
			key := chunk[:sp]     // Key of record
			value := chunk[sp+1:] // Value

			// Check Tree Type (So Order and Presence Doesn't Matter)
			switch string(key) {
			case "tree":
				commit.Tree = value
				continue
			case "parent":
				commit.Parent = value
				continue
			case "author":
				commit.Author = string(value)
				continue
			case "committer":
				commit.Committer = string(value)
				continue
			}

			// If First 4 main lines are done, then Message there
			// Assuming that PGP SIGNATURE doesn't exist
			if i > 3 {
				commit.Message = string(chunk)
				continue
			}
		}
	}

	// Return GitCommit
	return commit
}
