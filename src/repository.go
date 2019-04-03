// This file contains functions related to creating a new git repository
// Woruld probably be used in "a-git init" command

package src

import (
	"errors"
	"fmt"
	"os"
	"path"

	"gopkg.in/ini.v1"
)

// repository

// CreateRepository - creates directories and files necessary for a git repository
func CreateRepository(workdir string) error {
	var err error
	// Present Working Directory if "workdir" not specified
	if workdir == "" {
		workdir = "."
	}

	// wDir - the specified worktree (directory), by default it's "." (pwd)
	wDir, err := os.Stat(workdir)

	if os.IsNotExist(err) {
		// If wDir doesn't exist
		cErr := os.Mkdir(workdir, 0777)
		if cErr != nil {
			// If we failed to create wDir
			fmt.Println("Specified directory doesn't exist, Unable to even cretae new one,\n", err)
			return cErr
		} else {
			fmt.Println("Specified directory doesn't exist, Created!")
		}
	} else {
		// If wDir exist
		isDir := wDir.IsDir()
		if !isDir {
			// If wDir is not a directory
			return errors.New("The specified path is not a directory")
		}
	}

	gitdir := ".git" // Git directory name

	os.Mkdir(path.Join(workdir, gitdir), 0777) // Create a ".git" directory
	err = CreateRepoConf(workdir, gitdir)
	err = CreateRepoDirs(workdir, gitdir)

	if err != nil {
		// Remove the whole ".git" folder if some error
		os.RemoveAll(path.Join(workdir, gitdir))
		fmt.Println("Failed to initialize repository")
		fmt.Println(err)
	}

	return nil
}

// CreateRepoDirs - creates required git directories and files (inside ".git")
func CreateRepoDirs(workdir string, gitdir string) error {
	var err error

	// Created Directories "branches", "into", "refs" ect.
	err = os.MkdirAll(path.Join(workdir, gitdir, "branches"), 0777)
	err = os.MkdirAll(path.Join(workdir, gitdir, "info"), 0777)
	err = os.MkdirAll(path.Join(workdir, gitdir, "objects", "info"), 0777)
	err = os.MkdirAll(path.Join(workdir, gitdir, "objects", "pack"), 0777)
	err = os.MkdirAll(path.Join(workdir, gitdir, "refs", "heads"), 0777)
	err = os.MkdirAll(path.Join(workdir, gitdir, "refs", "tags"), 0777)

	// Create a "description" file
	fd, err := os.Create(path.Join(workdir, gitdir, "description"))
	_, err = fd.Write([]byte("Unnamed repository; edit this file 'description' to name the repository.\n")) // Default

	// Create a "HEAD" fiel
	fh, err := os.Create(path.Join(workdir, gitdir, "HEAD"))
	_, err = fh.Write([]byte("ref: refs/heads/master\n")) // Default

	return err
}

// CreateRepoConf - creates required "config" file (Microsoft INI type file)
func CreateRepoConf(workdir string, gitdir string) error {
	var err error

	// Filename of configuration ("config")
	fname := path.Join(workdir, gitdir, "config")

	// Create a new config file (but check if it is exist or not)
	if _, err = os.Stat(fname); os.IsNotExist(err) {
		_, err = os.Create(fname)
	}
	if err != nil {
		return err
	}

	// Load "config" file
	cfg, err := ini.Load(fname)
	if err != nil {
		return err
	}

	// Adding default git configuration and saving the file
	cfg.Section("core").Key("repositoryformatversion").SetValue("0")
	cfg.Section("core").Key("filemode").SetValue("true")
	cfg.Section("core").Key("bare").SetValue("false")
	cfg.Section("core").Key("logallrefupdates").SetValue("true")
	cfg.SaveTo(fname)

	return nil
}
