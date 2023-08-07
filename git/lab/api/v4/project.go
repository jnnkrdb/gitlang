package v4

/*

Find ore information about the GitLab API:

https://docs.gitlab.com/ee/api/repository_files.html
https://docs.gitlab.com/ee/api/personal_access_tokens.html

*/

// project information
type Project struct {
	ApiV4       string `json:"apiv4"`
	AccessToken string `json:"accesstoken"`
	ID          string `json:"projectid"`
}

// return the url for the specific project
//
// looks similar to "https://url.to.gitlab/api/v4/projects/_projectid_/repository/files/"
func (p Project) BaseURL() string {

	return p.ApiV4 + "/projects/" + p.ID + "/repository/files/"
}
