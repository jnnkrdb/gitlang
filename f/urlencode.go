package f

import "strings"

// encode the path to an url for apis like gitlab v4
//
// Parameters:
//   - `path` : string > path of the file in the repository, relative to the repository root
func EncodeURL(path string) string {

	for _, rplc := range [][]string{
		{"/", "//2F"},
		{".", "//2E"},
		{"//", "%"},
	} {

		path = strings.ReplaceAll(path, rplc[0], rplc[1])
	}

	return path
}
