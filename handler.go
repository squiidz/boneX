/********************************
*** Multiplexer for Go        ***
*** Bonex is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/go-zoo         ***
*********************************/

package bonex

import (
	"net/http"
)

// Arg content the key and value of a URL Params
type Arg struct {
	Key   string
	Value string
}

type Args []Arg

// GetValue return the value of the provided key
func (a *Args) GetValue(key string) string {
	for _, v := range *a {
		if v.Key == key {
			return v.Value
		}
	}
	return ""
}

// BoneHandler is a custom type of handler, to deal with URL params with a decent speed.
type BoneHandler func(http.ResponseWriter, *http.Request, Args)
