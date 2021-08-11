package server

import "errors"

var (
	errCfgWrongEnv = errors.New("wrong config for production")
	errCfgNoFile   = errors.New("no production config provided")
	errPublic      = errors.New("something went wrong, our bad")
)
