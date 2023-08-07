package v4

import (
	"encoding/json"

	"github.com/jnnkrdb/corerdb/fnc"
	"github.com/jnnkrdb/corerdb/prtcl"
)

// this struct will be parsed into a json-string. the json string can
// then be push to the gitlab api via a POST or PUT http-request.
//
// the encoding will be set to base64 by default. the content itself will be parsed
// to base64, so dont parse it to base64 before...
type FileInformation struct {
	Branch     string `json:"branch"`
	Encoding   string `json:"encoding,omitempty"`
	Content    string `json:"content,omitempty"`
	CommitMSG  string `json:"commit_message"`
	AuthorMail string `json:"author_email,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
}

// check the fileinfo struct for the necessary informations
func (fi *FileInformation) validate() {

	if fi.Content != "" {

		fi.Encoding = "base64"
		fi.Content = fnc.EncodeB64(fi.Content)
	}

	if fi.Branch == "" {

		fi.Branch = "master"
	}

	if fi.CommitMSG == "" {

		fi.CommitMSG = "api commit, " + prtcl.Timestamp()
	}
}

// formats the fileinformation struct into a json-string
// but the type will be of []byte
func (fi FileInformation) JSON() (jsn []byte) {

	fi.validate()

	var err error

	if jsn, err = json.Marshal(fi); err != nil {

		prtcl.PrintObject(fi, jsn, err)
	}

	return
}
