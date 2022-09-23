package git

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/jnnkrdb/corerdb/fnc"
	"github.com/jnnkrdb/corerdb/prtcl"
)

// struct for authentication for specific git repositories
type Repository struct {
	// protocol, used to connect to the git-repo
	Protocol string `json:"protocol"`
	// the user, accessing the gitlab-repo
	Username string `json:"username"`
	// accesstoken, for gitlab authentication against the v4-api
	// the accesstoken must be base64 encoded
	AccessToken string `json:"accesstoken"`
	// repos-uri, creates the URL with the protocol
	URI string `json:"uri"`
	// storage on the local filesystem
	Path string `json:"path"`
	// default branch
	Branch string `json:"branch"`
}

// load config from json
//
// Parameters:
//   - `path` : string > path to the jsonfile, which contains the settings
func (r *Repository) FromJSON(path string) (err error) {

	prtcl.Log.Println("loading gitlab repository-auth configuration from", path)

	if jsonf, err := os.ReadFile(path); err == nil {

		if err := json.Unmarshal(jsonf, r); err != nil {

			prtcl.PrintObject(jsonf, r, path, err)
		}

	} else {

		prtcl.PrintObject(jsonf, r, path, err)
	}

	return
}

// get the connection uri from the repos
func (r Repository) getURL() (url string) {

	url += r.Protocol + "://"

	if r.AccessToken != "" && r.Username != "" {

		url += r.Username + ":" + fnc.UnencodeB64(r.AccessToken) + "@"
	}

	url += r.URI + ""

	return
}

// clone function for repositories
func (r Repository) Clone() (err error) {

	prtcl.Log.Println("cloning gitlab repository:", r.Protocol+"://"+r.URI)

	var pathcmd string = ""

	if r.Path != "" {

		pathcmd = r.Path + "/"
	}

	cmd := exec.Command("git", "clone", r.getURL(), pathcmd)

	if err = cmd.Run(); err != nil {

		prtcl.PrintObject(r, pathcmd, cmd, err)
	}

	return
}

func (r Repository) Pull() (err error) {

	prtcl.Log.Println("pulling gitlab repository:", r.Protocol+"://"+r.URI)

	if err = os.Chdir(r.Path); err != nil {

		prtcl.PrintObject(r, err)

	} else {

		cmd := exec.Command("git", "pull", r.getURL(), r.Branch)

		if err = cmd.Run(); err != nil {

			prtcl.PrintObject(r, cmd, err)
		}
	}

	return
}
