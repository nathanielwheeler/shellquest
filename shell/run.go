package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Run will start a game loop, outputting to stdout and saving locally.
func Run() error {
	fmt.Println("Shell Quest!  (alpha)")

	save, err := LoadSave()
	if err != nil {
		return err
	}

	fmt.Println(RunCmd(save))
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

		res := RunCmd(save, strings.Split(input, " ")...)
    if res == nil {
      break
    }
	}
  fmt.Println("Exiting...")
	return nil
}

// RunCmd will take in a save pointer and variadic commands and make a response, automatically changing the save as needed.
func RunCmd(save *Save, cmds ...string) (res interface{}) {
	return nil
}
