package shell

import (
  "gopkg.in/yaml.v2"
)

type Room struct {
  Body string
}

func (g *Game) generateRooms() error {
  // get rooms from embedded yaml
  b, err := fs.ReadFile("rooms.yaml")
  if err != nil {
    return err
  }

  if err = yaml.Unmarshal(b, g.Rooms); err != nil {
    return err
  }

  return nil
}