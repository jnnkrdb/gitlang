package v2022_11_28

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jnnkrdb/gitlang/f"
)

const BaseApiUrl = "https://api.github.com"

func NewGithubConnection(username, authtoken string) *apiConn {
	return &apiConn{
		User:      username,
		AuthToken: authtoken,
		Headers: map[string]string{
			"X-GitHub-Api-Version": "2022-11-28",
			"Authorization":        fmt.Sprintf("Bearer %s", authtoken),
			"Accept":               "application/vnd.github+json",
		},
	}
}

type apiConn struct {
	User      string // user of the connection
	AuthToken string // ghp_biusbdfiusdf
	Headers   map[string]string
}

func (ac apiConn) newRequest(method string, url string, body io.Reader) (req *http.Request, err error) {
	if req, err = http.NewRequest(method, url, body); err == nil {
		for k, v := range ac.Headers {
			req.Header.Add(k, v)
		}
	}
	return
}

func (ac apiConn) CheckFile(repository, relpath string) (err error) {
	var req *http.Request
	if req, err = ac.newRequest(http.MethodGet, fmt.Sprintf("%s/repos/%s/%s/contents/%s", BaseApiUrl, ac.User, repository, f.EncodeURL(relpath)), nil); err == nil {
		fmt.Print(req)
	}
	return
}
