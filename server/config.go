package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type config struct {
	Port int
	Env  string
}

func loadConfig() (*config, error) {
	var (
		cfgFilename = ".config.json"
		cfg         = config{}
		isProd      = false
	)
	if len(os.Args) > 1 {
		for i, arg := range os.Args[1:] {
			switch arg {
			case "-c":
			case "--config":
				cfgFilename = os.Args[i+2]
			case "--port":
				cfg.Port, _ = strconv.Atoi(os.Args[i+2])
			case "--prod":
				isProd = true
			}
		}
	}

	b, err := ioutil.ReadFile(cfgFilename)
	if err != nil && !isProd {
		return devConfig(), nil
	} else if err != nil {
		fmt.Println("WARNING: missing PROD config!")
		return nil, errCfgNoFile
	}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		fmt.Fprintln(os.Stdout, "JSON unmarshal error")
		return nil, err
	}

	if cfg.isProd() != isProd {
		return nil, errCfgWrongEnv
	}

	return &cfg, nil
}

func devConfig() *config {
	return &config{
		Port: 9998,
		Env:  "dev",
	}
}

func (c *config) isProd() bool {
	return c.Env == "prod"
}
