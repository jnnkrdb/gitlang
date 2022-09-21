package gitlabv4

import (
	"bytes"
	"io"
	"net/http"

	cordb "github.com/jnnkrdb/cordb/f"
	"github.com/jnnkrdb/jlog"
)

/*

Find ore information about the GitLab API:

https://docs.gitlab.com/ee/api/repository_files.html
https://docs.gitlab.com/ee/api/personal_access_tokens.html

*/

// project information
type Project struct {
	AccessToken string `json:"accesstoken"`
	ID          string `json:"projectid"`
}

// return the url for the specific project
//
// looks similar to "https://url.to.gitlab/api/v4/projects/_projectid_/repository/files/"
//
// Parameters:
//   - `apiv4` : string > api url of the v4 version, something like "https://url.to.gitlab/api/v4/"
func (p Project) BaseURL(apiv4 string) string {

	return apiv4 + "/projects/" + p.ID + "/repository/files/"
}

// send an v4request to the gitlab project, if the returned int value
// equals "1" there was an error with the request
//
// Parameters:
//   - `api` : string > contains the base v4 path of the api
//   - `req` : V4Request > request information, send to the gitlab api
func (p Project) Send(api string, req V4Request) int {

	req.CheckFile(api, p)

	if request, err := http.NewRequest(req.HTTPMethod, p.BaseURL(api)+req.FilePath, bytes.NewReader(req.File.JSON())); err != nil {

		jlog.PrintObject(api, req, p, request, err)

	} else {

		request.Header.Add("PRIVATE-TOKEN", cordb.UnencodeB64(p.AccessToken))

		request.Header.Add("Content-Type", "application/json")

		if result, err := http.DefaultClient.Do(request); err == nil {

			defer result.Body.Close()

			jlog.PrintObject(io.ReadAll(result.Body))

			return result.StatusCode

		} else {

			jlog.PrintObject(api, req, p, request, result, err)
		}
	}

	return 1
}
