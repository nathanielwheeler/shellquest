package server

import (
	"net/http"
	"sync"
)

func (s *server) handleView(tpls ...string) http.HandlerFunc {
	var (
		init sync.Once
		v    *view
		err  error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			v = s.newView("home")
			err = s.addViewContent(v, "home")
		})
		if err != nil {
			http.Error(w, errPublic.Error(), http.StatusInternalServerError)
      return
		}
		s.execView(v, w, r)
	}
}
