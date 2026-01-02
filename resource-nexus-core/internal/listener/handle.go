package listener

import (
	"fmt"
	"net/http"
)

// AddRoute adds a new route to the listener.
func (l *Listener) AddRoute(method, url string, handler http.HandlerFunc) {
	l.multiplexer.Handle(url, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			l.logger.Warn(fmt.Sprintf("method not allowed: %s %s", r.Method, r.URL.Path))
			
			return
		}

		handler.ServeHTTP(w, r)
	}))
}
