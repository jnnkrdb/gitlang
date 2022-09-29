package gitlabv4

import "github.com/jnnkrdb/gitlang/f"

// create a new push request with predefined variables
// and a relative file path, which will automatically be url-encoded
//
// Parameters:
//   - `file` : FileInformation > information about the requested file
//   - `relativepath` : string > path of the requested file, relative to the repository root
func CreatePushRequest(file FileInformation, relativepath string) V4Request {

	return V4Request{
		FilePath: f.EncodeURL(relativepath),
		File:     file,
	}
}

// create a new push request with predefined variables
// and a relative file path, which will automatically be url-encoded
//
// Parameters:
//   - `relativepath` : string > path of the requested file, relative to the repository root
func CreateGetRequest(relativepath string) V4Request {

	return V4Request{
		FilePath: f.EncodeURL(relativepath),
		File: FileInformation{
			Branch: "master",
		},
	}
}
