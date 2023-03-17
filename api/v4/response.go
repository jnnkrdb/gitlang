package v4

import (
	"encoding/base64"
	"fmt"
)

type V4Response struct {
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Size            int    `json:"size"`
	Encoding        string `json:"encoding"`
	Content         string `json:"content"`
	Content_SHA256  string `json:"content_sha256"`
	Ref             string `json:"ref"`
	BlobID          string `json:"blob_id"`
	CommitID        string `json:"commit_id"`
	LastCommitID    string `json:"last_commit_id"`
	ExecuteFilemode bool   `json:"execute_filemode"`
}

// translate the base64 encoded content into unencoded string
func (v4r V4Response) UnencodedContent() (string, error) {
	if b, err := base64.StdEncoding.DecodeString(v4r.Content); err != nil {
		return "", fmt.Errorf("error unencoding response content: %v", err)
	} else {
		return string(b), nil
	}
}
