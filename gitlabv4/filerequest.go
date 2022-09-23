package gitlabv4

import (
	"net/http"

	"github.com/jnnkrdb/corerdb/fnc"
	"github.com/jnnkrdb/corerdb/prtcl"
	"github.com/jnnkrdb/gitlang/f"
)

// struct to describe a file request for the gtlab v4 api
type V4Request struct {
	FilePath   string
	HTTPMethod string
	File       FileInformation
}

// create a new request with predefined variables
// and a relative file path, which will automatically be url-encoded
func CreateRequest(file FileInformation, relativepath string) V4Request {

	return V4Request{
		FilePath:   f.EncodeURL(relativepath),
		File:       file,
		HTTPMethod: "POST",
	}
}

// check current file state, if the file exists in the specific filepath of the request
// the request method will be put, to update the file, else its post, to create the file
func (v4 *V4Request) CheckFile(api string, proj Project) {

	v4.HTTPMethod = "POST"

	if request, err := http.NewRequest("GET", proj.BaseURL(api)+v4.FilePath+"?ref="+v4.File.Branch, nil); err != nil {

		prtcl.PrintObject(api, v4, proj, request, err)

	} else {

		request.Header.Add("PRIVATE-TOKEN", fnc.UnencodeB64(proj.AccessToken))

		if result, err := http.DefaultClient.Do(request); err == nil {

			switch result.StatusCode {

			case 200:

				v4.HTTPMethod = "PUT"

			case 404:

				v4.HTTPMethod = "POST"

			default:

				prtcl.PrintObject(v4, api, proj, request, result, err)
			}
		}
	}
}
