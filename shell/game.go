package shell

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Game struct {
	CurrentRoom *Room
	Rooms       map[string]Room
}

func NewGame() (*Game, error) {
	g := &Game{}
	if err := g.generateRooms(); err != nil {
    return nil, err
  }
	return g, nil
}

func LoadGame() (*Game, error) {
	var (
		game           *Game
		defaultSaveDir = filepath.Join(os.Getenv("HOME"), ".config", "shellquest")
		defaultSave    = filepath.Join(defaultSaveDir, "save.json")
		err            error
	)
	fmt.Println(defaultSaveDir)

	// check if local save directory exists, create one if not
	if _, err = os.Stat(defaultSaveDir); os.IsNotExist(err) {
		if err = os.MkdirAll(defaultSaveDir, 0770); err != nil {
			return nil, err
		}
	}

	// check for save data, make a new game file and return if there is none.
	if _, err = os.Stat(defaultSave); os.IsNotExist(err) {
		game, err = NewGame()
    if err != nil {
      return nil, err
    }

		b, err := json.Marshal(&game)
		if err != nil {
			return nil, err
		}
    
		if err = ioutil.WriteFile(defaultSave, b, 0777); err != nil {
			return nil, err
		}
		return game, nil
	}

	// load data and return
	b, err := ioutil.ReadFile(defaultSave)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &game); err != nil {
		return nil, err
	}
	return game, nil
}
