package v4

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jnnkrdb/corerdb/fnc"
	"github.com/jnnkrdb/corerdb/prtcl"
	"github.com/jnnkrdb/gitlang/f"
)

// struct to describe a file request for the gtlab v4 api
type V4Request struct {
	FilePath string
	File     FileInformation
}

// check current file state, if the file exists in the specific filepath of the request
// the request method will be put, to update the file, else its post, to create the file
//
// Parameters:
//   - `proj` : Project > destination project
func (v4 *V4Request) checkFile(proj Project) string {

	if request, err := http.NewRequest("GET", proj.BaseURL()+f.EncodeURL(v4.FilePath)+"?ref="+v4.File.Branch, nil); err != nil {

		prtcl.PrintObject(v4, proj, request, err)

	} else {

		request.Header.Add("PRIVATE-TOKEN", fnc.UnencodeB64(proj.AccessToken))

		if result, err := http.DefaultClient.Do(request); err == nil {

			switch result.StatusCode {

			case 200:

				return "PUT"

			case 404:

				return "POST"

			default:

				prtcl.PrintObject(v4, proj, request, result, err)

			}
		}
	}

	return "POST"
}

// send a post or put request
//
// Parameters:
//   - `proj` : Project > destination project
func (v4 V4Request) Push(proj Project) error {

	req := v4.checkFile(proj)

	if request, err := http.NewRequest(req, proj.BaseURL()+f.EncodeURL(v4.FilePath), bytes.NewReader(v4.File.JSON())); err != nil {

		prtcl.PrintObject(v4, proj, request, err)

		return err

	} else {

		request.Header.Add("PRIVATE-TOKEN", fnc.UnencodeB64(proj.AccessToken))

		request.Header.Add("Content-Type", "application/json")

		if result, err := http.DefaultClient.Do(request); err == nil {

			defer result.Body.Close()

			response, err := io.ReadAll(result.Body)

			prtcl.PrintObject(result, string(response), err)

			return nil

		} else {

			prtcl.PrintObject(v4, proj, request, result, err)

			return err
		}
	}
}

// receive a file via get
//
// Parameters:
//   - `proj` : Project > destination project
func (v4 V4Request) Get(proj Project) (gr GetResponse, err error) {

	if v4.File.Branch == "" {
		v4.File.Branch = "master"
	}

	var request *http.Request
	if request, err = http.NewRequest("GET", proj.BaseURL()+f.EncodeURL(v4.FilePath)+"?ref="+v4.File.Branch, nil); err != nil {

		prtcl.PrintObject(v4, proj, request, gr, err)

	} else {

		request.Header.Add("PRIVATE-TOKEN", fnc.UnencodeB64(proj.AccessToken))

		request.Header.Add("Content-Type", "application/json")

		var result *http.Response
		if result, err = http.DefaultClient.Do(request); err != nil {

			prtcl.PrintObject(v4, proj, request, result, gr, err)

		} else {

			defer result.Body.Close()

			var bytes []byte
			if bytes, err = io.ReadAll(result.Body); err != nil {

				prtcl.PrintObject(v4, proj, request, result, bytes, gr, err)

			} else {

				err = json.Unmarshal(bytes, &gr)
			}
		}
	}

	return
}

// delete a file via delete
//
// Parameters:
//   - `proj` : Project > destination project
func (v4 V4Request) Delete(proj Project) (err error) {

	var request *http.Request
	if request, err = http.NewRequest("DELETE", proj.BaseURL()+f.EncodeURL(v4.FilePath), bytes.NewReader(v4.File.JSON())); err != nil {

		prtcl.PrintObject(v4, proj, request, err)

	} else {

		request.Header.Add("PRIVATE-TOKEN", fnc.UnencodeB64(proj.AccessToken))

		request.Header.Add("Content-Type", "application/json")

		if result, err := http.DefaultClient.Do(request); err == nil {

			defer result.Body.Close()

			response, err := io.ReadAll(result.Body)

			prtcl.PrintObject(result, string(response), err)

		} else {

			prtcl.PrintObject(v4, proj, request, result, err)

		}
	}

	return
}
