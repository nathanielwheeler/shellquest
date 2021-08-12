package shell

// Cmd will take in variadic commands and make a response, automatically saving the game.
func (g *Game) Cmd(cmds ...string) (res interface{}) {
  switch cmds[0] {
  case "look":
    return g.CurrentRoom.Body
  }

	return nil
}
