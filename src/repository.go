// This file contains functions related to creating a new git repository
// Woruld probably be used in "a-git init" command

package src

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// GitRepository - Repository type
type GitRepository struct {
	workdir string
	gitdir  string
}

// CreateRepository - creates directories and files necessary for a git repository
func CreateRepository(workdir string) error {
	var err error

	gitdir := path.Join(workdir, ".git")      // ".git" directory
	newRepo := GitRepository{workdir, gitdir} // New repo instance

	// wDir - the specified worktree (directory), by default it's "." (pwd)
	wDir, err := os.Stat(newRepo.workdir)

	if os.IsNotExist(err) {
		// If wDir doesn't exist
		cErr := os.Mkdir(newRepo.workdir, 0777)
		if cErr != nil {
			// If we failed to create wDir
			fmt.Println("Specified directory doesn't exist, Unable to even cretae new one,\n", err)
			return cErr
		}
		fmt.Println("Specified directory doesn't exist, Created!")

	} else {
		// If wDir exist
		isDir := wDir.IsDir()
		if !isDir {
			// If wDir is not a directory
			return errors.New("The specified path is not a directory")
		}
	}

	os.Mkdir(newRepo.gitdir, 0755) // Create a ".git" directory
	err = CreateRepoConf(newRepo.gitdir)
	err = CreateRepoDirs(newRepo.gitdir)

	if err != nil {
		// Remove the whole ".git" folder if some error
		os.RemoveAll(path.Join(newRepo.gitdir))
		fmt.Println("Failed to initialize repository")
		fmt.Println(err)
	}

	return nil
}

// CreateRepoDirs - creates required git directories and files (inside ".git")
func CreateRepoDirs(gitdir string) error {
	var err error

	// Created Directories "branches", "into", "refs" ect.
	err = os.MkdirAll(path.Join(gitdir, "branches"), 0777)
	err = os.MkdirAll(path.Join(gitdir, "info"), 0777)
	err = os.MkdirAll(path.Join(gitdir, "objects", "info"), 0777)
	err = os.MkdirAll(path.Join(gitdir, "objects", "pack"), 0777)
	err = os.MkdirAll(path.Join(gitdir, "refs", "heads"), 0777)
	err = os.MkdirAll(path.Join(gitdir, "refs", "tags"), 0777)

	// Create a "description" file
	fd, err := os.Create(path.Join(gitdir, "description"))
	_, err = fd.Write([]byte("Unnamed repository; edit this file 'description' to name the repository.\n")) // Default

	// Create a "HEAD" fiel
	fh, err := os.Create(path.Join(gitdir, "HEAD"))
	_, err = fh.Write([]byte("ref: refs/heads/master\n")) // Default

	return err
}

// CreateRepoConf - creates required "config" file (Microsoft INI type file)
func CreateRepoConf(gitdir string) error {
	var err error

	// Filename of configuration ("config")
	fname := path.Join(gitdir, "config")

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

// FindRepository - Finds repository (".git" folder)
// Sometimes users current directory might not be the repo workdir
// For those cases this function will find recursively find the .git directory going up on
// on the directory tree
func FindRepository(currentdir string) (GitRepository, error) {
	// Absolute current path
	absCurrent, err := filepath.Abs(currentdir)
	if err != nil {
		return GitRepository{}, err
	}

	// Absolute git directory path
	gitdir := filepath.Join(absCurrent, ".git")

	// Check if gitdir exists or not
	gDir, err := os.Stat(gitdir)
	if err == nil {
		isGitDir := gDir.IsDir()
		if isGitDir {
			return GitRepository{workdir: currentdir, gitdir: gitdir}, nil
		}
	}

	// Relative parent directory
	absParent := filepath.Join(absCurrent, "..")
	// Check if root or not
	if absParent == absCurrent {
		return GitRepository{}, errors.New("No git directory")
	}

	// Recursively run for the parent
	return FindRepository(absParent)
}
