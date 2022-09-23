package gitlabv4

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/jnnkrdb/corerdb/prtcl"
)

// this struct will be parsed into a json-string. the json string can
// then be push to the gitlab api via a POST or PUT http-request.
//
// the encoding will be set to base64 by default. the content itself will be parsed
// to base64, so dont parse it to base64 before...
type FileInformation struct {
	Branch     string `json:"branch"`
	Encoding   string `json:"encoding"`
	Content    string `json:"content"`
	CommitMSG  string `json:"commit_message"`
	AuthorMail string `json:"author_email"`
	AuthorName string `json:"author_name"`
}

// check the fileinfo struct for the necessary informations
func (fi *FileInformation) check() {

	if fi.Content != "" {

		fi.Encoding = "base64"
		fi.Content = base64.StdEncoding.EncodeToString([]byte(fi.Content))
	}

	if fi.Branch == "" {

		fi.Branch = "master"
	}

	if fi.CommitMSG == "" {

		fi.CommitMSG = "api commit, " + time.Now().Format(time.RFC3339)
	}
}

// formats the fileinformation struct into a json-string
// but the type will be of []byte
func (fi FileInformation) JSON() (res []byte) {

	fi.check()

	if res, err := json.Marshal(fi); err != nil {

		prtcl.PrintObject(fi, res, err)
	}

	return
}
