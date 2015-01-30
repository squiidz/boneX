/********************************
*** Multiplexer for Go        ***
*** Bonex is under MIT license ***
*** Code by CodingFerret      ***
*** github.com/go-zoo         ***
*********************************/

package bonex

import "net/http"

// Mux have routes and a notFound handler
// Route: all the registred route
// notFound: 404 handler, default http.NotFound if not provided
type Mux struct {
	Routes   map[string][]*Route
	Static   map[string]*Route
	notFound http.HandlerFunc
}

var (
	method = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS"}
)

// New create a pointer to a Mux instance
func New() *Mux {
	return &Mux{
		Routes: make(map[string][]*Route),
		Static: make(map[string]*Route),
	}
}

// Serve http request
func (m *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Check if the request path doesn't end with /

	if !m.valid(req.URL.Path) {
		if key, ok := m.isStatic(req.URL.Path); ok {
			m.Static[key].Handler(rw, req, nil)
			return
		}
		for !m.valid(req.URL.Path) {
			req.URL.Path = req.URL.Path[:len(req.URL.Path)-1]
		}

		rw.Header().Set("Location", req.URL.Path)
		rw.WriteHeader(http.StatusFound)
	}

	// Loop over all the registred route.
	for _, r := range m.Routes[req.Method] {
		// If the route is equal to the request path.
		if req.URL.Path == r.Path && !r.Params {
			r.Handler(rw, req, nil)
			return
		} else if r.Params {
			if v, ok := r.Match(req); ok {
				r.Handler(rw, req, v)
				return
			}
		}
	}

	// If no valid Route found, check for static file
	if key, ok := m.isStatic(req.URL.Path); ok {
		m.Static[key].Handler(rw, req, nil)
		return
	}
	m.HandleNotFound(rw, req)

}

// Wrapper for standard http.Handler
func (m *Mux) HandleFunc(path string, handler http.Handler) {
	r := NewRoute(path, func(rw http.ResponseWriter, req *http.Request, args Args) {
		handler.ServeHTTP(rw, req)
	})
	for _, mt := range method {
		m.register(mt, r)
	}
}

// Handle add a new route to the Mux without a HTTP method
func (m *Mux) Handle(path string, handler BoneHandler) {
	r := NewRoute(path, handler)
	for _, mt := range method {
		m.register(mt, r)
	}
}

// Get add a new route to the Mux with the Get method
func (m *Mux) Get(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("GET", r)
	return r
}

// Post add a new route to the Mux with the Post method
func (m *Mux) Post(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("POST", r)
	return r
}

// Put add a new route to the Mux with the Put method
func (m *Mux) Put(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("PUT", r)
	return r
}

// Delete add a new route to the Mux with the Delete method
func (m *Mux) Delete(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("DELETE", r)
	return r
}

// Head add a new route to the Mux with the Head method
func (m *Mux) Head(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("HEAD", r)
	return r
}

// Patch add a new route to the Mux with the Patch method
func (m *Mux) Patch(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("PATCH", r)
	return r
}

// Options add a new route to the Mux with the Options method
func (m *Mux) Options(path string, handler BoneHandler) *Route {
	r := NewRoute(path, handler)
	m.register("OPTIONS", r)
	return r
}

// NotFound the mux custom 404 handler
func (m *Mux) NotFound(handler http.HandlerFunc) {
	m.notFound = handler
}

// Register the new route in the router with the provided method and handler
func (m *Mux) register(method string, r *Route) {
	if m.valid(r.Path) {
		m.Routes[method] = append(m.Routes[method], r)
		return
	}
	m.Static[r.Path] = r
}
