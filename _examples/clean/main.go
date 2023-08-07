package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// Open the repository.
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Get the worktree.
	w, err := r.Worktree()
	CheckIfError(err)

	// Fetch the upstream master branch.
	Info("git fetch -v")

	err = r.Fetch(&git.FetchOptions{
		RemoteName: "origin",
	})
	if err == git.NoErrAlreadyUpToDate {
		fmt.Println(err)
	} else if err != nil {
		CheckIfError(err)
	}

	// Reset the local master branch to the upstream master branch.
	Info("git reset --hard @{upstream}")

	// Get the reference for remote origin/master
	ref, err := r.Reference(plumbing.NewRemoteReferenceName("origin", "master"), true)
	CheckIfError(err)

	// Reset the worktree to origin/master
	err = w.Reset(&git.ResetOptions{
		Commit: ref.Hash(),
		Mode:   git.HardReset,
	})
	CheckIfError(err)

	// Clean all untracked files from the worktree.
	Info("git clean -xfd")

	cleanOpts := &git.CleanOptions{
		Dir: true,
	}
	err = w.Clean(cleanOpts)
	CheckIfError(err)

	// Print a message indicating that the reset and clean was successful.
	fmt.Println("Fetch, Reset, and clean to upstream origin/master branch successful.")
}
