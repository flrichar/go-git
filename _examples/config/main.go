package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

// Open an existing repository in a specific folder.
func main() {
	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Get the .git/config information
	Info("git config --get remote.origin.url")

	// Get the remote origin.
	origin, err := r.Remote("origin")
	CheckIfError(err)

	// Get the remote origin URL.
	originURL := origin.Config().URLs[0]

	// Print the remote origin URL.
	fmt.Println(originURL)
}
