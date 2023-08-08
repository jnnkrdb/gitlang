package f

import "strings"

var _encodings = [][]string{
	{` `, `%20`},
	{`!`, `%21`},
	{`"`, `%22`},
	{`#`, `%23`},
	{`$`, `%24`},
	{`%`, `%25`},
	{`&`, `%26`},
	{`'`, `%27`},
	{`(`, `%28`},
	{`)`, `%29`},
	{`*`, `%2A`},
	{`+`, `%2B`},
	{`,`, `%2C`},
	{`-`, `%2D`},
	{`.`, `%2E`},
	{`/`, `%2F`},
	{`:`, `%3A`},
	{`;`, `%3B`},
	{`<`, `%3C`},
	{`=`, `%3D`},
	{`>`, `%3E`},
	{`?`, `%3F`},
	{`@`, `%40`},
	{`[`, `%5B`},
	{`\`, `%5C`},
	{`]`, `%5D`},
	{`{`, `%7B`},
	{`}`, `%7D`},
	{`|`, `%7C`},
}

// encode the path to an url readable path
func EncodeURL(path string) string {
	for index := range _encodings {
		path = strings.ReplaceAll(path, _encodings[index][0], _encodings[index][1])
	}
	return path
}

// unencode the url to an path
func UnencodeURL(url string) string {
	for index := range _encodings {
		url = strings.ReplaceAll(url, _encodings[index][1], _encodings[index][0])
	}
	return url
}
