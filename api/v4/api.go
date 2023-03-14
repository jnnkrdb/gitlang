package v4

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jnnkrdb/gitlang/f"
)

const _notFoundErrorMsg string = "[Code 404] Not Found"

/*

Find ore information about the GitLab API:

https://docs.gitlab.com/ee/api/repository_files.html
https://docs.gitlab.com/ee/api/personal_access_tokens.html

*/

// base url for the gitlab v4 api
type ApiV4 string

// create the base for an gitlab v4 api request with the accesstoken and the project id
func (v4 ApiV4) Request(projectid int, accesstoken string) *v4request {

	var req *v4request = &v4request{
		api: v4,
		pid: projectid,
		at:  accesstoken,
	}
	req.detailedInfo.Branch = "master"
	req.detailedInfo.Encoding = "base64"
	req.detailedInfo.Commit_Message = fmt.Sprintf("commit by github.com/jnnkrdb/gitlang/api/v4 [%s]", time.Now().Format(time.RFC3339))
	return req
}

// structural type to store information about the actual request
type v4request struct {
	api          ApiV4
	at           string
	pid          int
	detailedInfo struct {
		Branch         string `json:"branch"`
		Encoding       string `json:"encoding,omitempty"`
		Content        string `json:"content,omitempty"`
		Commit_Message string `json:"commit_message"`
		Author_Mail    string `json:"author_email,omitempty"`
		Author_Name    string `json:"author_name,omitempty"`
	}
}

// -------------------------------------------------------------------- helperfunctions for the v4request

// return the url for the specific project
//
// looks similar to "https://url.to.gitlab/api/v4/projects/_projectid_/repository/files/"
func (v4r v4request) filesurl() string {
	return fmt.Sprintf("%s/projects/%d/repository/files/", v4r.api, v4r.pid)
}

// -------------------------------------------------------------------- external functions to configure the request

// set the branch for the request, if notset, the branch will be master
func (v4r *v4request) Branch(branch string) *v4request {
	v4r.detailedInfo.Branch = branch
	return v4r
}

// set the content for the request if neccessary, otherwise the content will remain empty
func (v4r *v4request) Content(content string) *v4request {
	v4r.detailedInfo.Content = content
	return v4r
}

// set the commitmessage for the request
// if not set manually, the default is: "commit by github.com/jnnkrdb/gitlang/api/v4 [2006-01-02T15:04:05Z07:00]"
func (v4r *v4request) Commit_Message(commit string) *v4request {
	v4r.detailedInfo.Commit_Message = commit
	return v4r
}

// set the author mail if required, otherwise the author_mail will remain empty
func (v4r *v4request) Author_Mail(mail string) *v4request {
	v4r.detailedInfo.Author_Mail = mail
	return v4r
}

// set the author name if required, otherwise the author_name will remain empty
func (v4r *v4request) Author_Name(name string) *v4request {
	v4r.detailedInfo.Author_Name = name
	return v4r
}

// -------------------------------------------------------------------- actual requests

// get a specific file from gitlab
func (v4r v4request) Get(file string) (res v4response, err error) {
	var httpreq *http.Request
	if httpreq, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?ref=%s", v4r.filesurl(), f.EncodeURL(file), v4r.detailedInfo.Branch), nil); err == nil {
		httpreq.Header.Add("PRIVATE-TOKEN", v4r.at)
		httpreq.Header.Add("Content-Type", "application/json")
		var httpresp *http.Response
		if httpresp, err = http.DefaultClient.Do(httpreq); err == nil {
			defer httpresp.Body.Close()
			switch httpresp.StatusCode {
			case http.StatusOK:
				err = json.NewDecoder(httpresp.Body).Decode(&res)
			case http.StatusNotFound:
				err = fmt.Errorf(_notFoundErrorMsg)
			default:
				err = fmt.Errorf("[Code %d]could'nt parse response from %s: %v", httpresp.StatusCode, httpresp.Request.URL, httpresp.Body)
			}
		}
	}
	return
}

// push a file, it will create the file if doesnt exist or update the file if exists
func (v4r v4request) Push(file string) (*http.Response, error) {
	v4r.detailedInfo.Content = base64.StdEncoding.EncodeToString([]byte(v4r.detailedInfo.Content))
	// defining a function to execute the file upload
	var upload = func(method string) (httpresp *http.Response, err error) {
		var jsn []byte
		if jsn, err = json.Marshal(v4r.detailedInfo); err == nil {
			var httpreq *http.Request
			if httpreq, err = http.NewRequest(method, fmt.Sprintf("%s%s", v4r.filesurl(), f.EncodeURL(file)), bytes.NewReader(jsn)); err == nil {
				httpreq.Header.Add("PRIVATE-TOKEN", v4r.at)
				httpreq.Header.Add("Content-Type", "application/json")
				httpresp, err = http.DefaultClient.Do(httpreq)
			}
		}
		return
	}
	// check if the file already exists
	_, err := v4r.Get(file)
	switch {
	case err == nil:
		return upload(http.MethodPut)
	case err.Error() == _notFoundErrorMsg:
		return upload(http.MethodPost)
	}
	return nil, err
}

// delete a specific file
func (v4r v4request) Delete(file string) (*http.Response, error) {
	var (
		httpreq *http.Request
		jsn     []byte
		err     error
	)
	if jsn, err = json.Marshal(v4r.detailedInfo); err == nil {
		if httpreq, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", v4r.filesurl(), f.EncodeURL(file)), bytes.NewReader(jsn)); err == nil {
			httpreq.Header.Add("PRIVATE-TOKEN", v4r.at)
			httpreq.Header.Add("Content-Type", "application/json")
			return http.DefaultClient.Do(httpreq)
		}
	}
	return nil, err
}
