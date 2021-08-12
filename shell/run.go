package shell

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed rooms.yaml
var fs embed.FS

// Run will start a game loop, outputting to stdout and saving locally.
func Run() error {
	fmt.Println("Shell Quest!  (alpha)")

	fmt.Println("Loading save...")
	game, err := LoadGame()
	if err != nil {
		return err
	}

	fmt.Println(game.Cmd("look"))
	reader := bufio.NewReader(os.Stdin)

	for {
		// User input
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("input error: %s", err)
		}

		// Prep input
		input = strings.TrimSuffix(input, "\n")
		if input == "" {
			fmt.Println("Please enter something.")
			continue
		}

		res := game.Cmd(strings.Split(input, " ")...)
		if res == nil {
			break
		}
	}
	fmt.Println("Exiting...")
	return nil
}
