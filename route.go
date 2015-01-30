/********************************
*** Multiplexer for Go        ***
*** Bonex is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/go-zoo         ***
*********************************/

package bonex

import (
	"net/http"
	"strings"
)

// Evaluator define the type of function need to evaluate a URL params.
type Evaluator func(string) bool

// Eval bind each evaluating function with the URL params.
func (r *Route) Eval(eva ...Evaluator) {
	if eva != nil {
		var arg []string
		for _, v := range r.Pattern {
			arg = append(arg, v)
		}

		for i, e := range eva {
			if i < len(arg) {
				r.Evaluate[arg[i]] = e
			}
		}
	}
}

// Route content the required information for a valid route
// Path: is the Route URL
// Size: is the length of the path
// Token: is the value of each part of the path, split by /
// Pattern: is content information about the route, if it's have a route variable
// handler: is the handler who handle this route
// Method: define HTTP method on the route
type Route struct {
	Path     string
	Size     int
	Token    Token
	Params   bool
	Pattern  map[int]string
	Handler  BoneHandler
	Evaluate map[string]Evaluator
	Method   string
}

// Token content all value of a spliting route path
// Tokens: string value of each token
// size: number of token
type Token struct {
	raw    []int
	Tokens []string
	Size   int
}

// NewRoute return a pointer to a Route instance and call save() on it
func NewRoute(url string, h BoneHandler) *Route {
	r := &Route{Path: url, Handler: h}
	r.save()
	return r
}

// Save, set automaticly the the Route.Size and Route.Pattern value
func (r *Route) save() {
	r.Size = len(r.Path)
	r.Token.Tokens = strings.Split(r.Path, "/")
	r.Pattern = make(map[int]string)

	for i, s := range r.Token.Tokens {
		if len(s) >= 1 {
			if s[:1] == ":" {
				r.Pattern[i] = s[1:]
				r.Params = true
			} else {
				r.Token.raw = append(r.Token.raw, i)
			}
		}
		r.Token.Size++
	}
	if r.Params {
		r.Evaluate = make(map[string]Evaluator)
	}
}

// Match check if the request match the route Pattern
func (r *Route) Match(req *http.Request) (Args, bool) {
	ss := strings.Split(req.URL.Path, "/")
	args := Args{}

	if len(ss) == r.Token.Size {
		for _, v := range r.Token.raw {
			if ss[v] != r.Token.Tokens[v] {
				return nil, false
			}
		}
		for k, v := range r.Pattern {
			if r.Evaluate[v] != nil {
				if r.Evaluate[v](ss[k]) {
					args = append(args, Arg{Key: v, Value: ss[k]})
				} else {
					return args, false
				}
			} else {
				args = append(args, Arg{Key: v, Value: ss[k]})
			}
		}
		return args, true
	}

	return nil, false
}
