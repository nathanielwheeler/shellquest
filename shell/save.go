package shell

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Save struct {
}

func LoadSave() (*Save, error) {
	var (
		save           Save
		defaultSaveDir = filepath.Join(os.Getenv("Home"), ".config", "shellquest")
		defaultSave    = filepath.Join(defaultSaveDir, "save.json")
		err            error
	)

	// check if local save directory exists, create one if not
	if _, err = os.Stat(defaultSaveDir); os.IsNotExist(err) {
		if err = os.MkdirAll(defaultSaveDir, 0770); err != nil {
			return nil, err
		}
	}

	// check for save data, make a new file if there is none.
	if _, err = os.Stat(defaultSave); os.IsNotExist(err) {
		save = Save{}

		b, err := json.Marshal(&save)
    if err != nil {
			return nil, err
		}
		if err = ioutil.WriteFile(defaultSave, b, 0777); err != nil {
      return nil, err
    }
	}

  return &save, nil
}
