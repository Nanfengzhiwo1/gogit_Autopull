package main

import (
	//inner

	"os"
	"path/filepath"
	"strings"

	// //outter
	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

var GIT_URL string = "https://github.com/Nanfengzhiwo1/gitpython_AutoPull.git"
var PATH, _ = os.Getwd()
var PARENT_PATH string = filepath.Dir(PATH)
var COMMON_PARENT_PATH string = strings.Replace(PARENT_PATH, "\\", "/", -1)
var GIT_BRANCH string = "main"

func main() {
	// CheckArgs("<url>", "<directory>", "<github_username>", "<github_password>")
	git_directory := COMMON_PARENT_PATH + "/" + GIT_BRANCH
	url, directory, branch := GIT_URL, git_directory, GIT_BRANCH

	// Clone the given repository to the given directory
	Info("git clone %s %s", url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		// Auth: &http.BasicAuth{
		// 	Username: username,
		// 	Password: password,
		// },
		URL:        url,
		RemoteName: branch,
		Progress:   os.Stdout,
	})
	CheckIfError(err)

}