package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func Run(fs *embed.FS) error {
	s, err := newServer(fs)
	if err != nil {
		return err
	}

	fmt.Printf("Using %s config...\n", strings.ToUpper(s.c.Env))

	port := fmt.Sprintf(":%d", s.c.Port)
	fmt.Printf("Now listening on %s...\n", port)
	err = http.ListenAndServe(port, s)
	if err != nil {
		return err
	}
	return nil
}

type server struct {
  c  *config
	l  *log.Logger
	fs *embed.FS
	r  *mux.Router
}

func newServer(fs *embed.FS) (*server, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	s := &server{
		c: cfg,
		l: log.New(os.Stdout, "servertemplate: ", log.Lshortfile),
    fs: fs,
		r: mux.NewRouter(),
	}
	s.routes()
	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
