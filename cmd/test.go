package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func promptForWorkload(challenges map[string]interface{}) []string {
	options := []string{}
	for k := range challenges {
		options = append(options, k)
	}

	var result string
	var err error

	prompt := promptui.Select{
		Label: "Select a workload",
		Items: options,
	}

	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	if _, ok := challenges[result].(map[string]interface{}); ok {
		return promptForWorkload(challenges[result].(map[string]interface{}))
	}

	return challenges[result].([]string)
}

func handleRun(cmd *cobra.Command, args []string) {
	workload := promptForWorkload(challenges)
	for {
		ok, _ := ExecuteCommand([]string{"go", "build", "-C", "challenges", "-o", "go-gossip-gloomers"})
		if !ok {
			fmt.Println("Error building go-gossip-gloomers")
			return
		}

		command := []string{"test", "test"}
		command = append(command, workload...)
		ExecuteShell("./maelstrom/maelstrom", command)

		fmt.Print("\nPress 'R' to rerun the test, 'C' to change the workload or 'Q' to quit: ")

		// Put terminal in raw mode
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Failed to set raw mode:", err)
			return
		}

		// Read a single character
		buffer := make([]byte, 1)
		os.Stdin.Read(buffer)

		// Restore terminal
		term.Restore(int(os.Stdin.Fd()), oldState)

		// Print newline for clean output
		fmt.Println()

		// Convert to uppercase for case-insensitive comparison
		input := strings.ToUpper(string(buffer))

		if input == "Q" {
			return
		} else if input == "R" {
			continue
		} else if input == "C" {
			workload = promptForWorkload(challenges)
			continue
		} else {
			fmt.Println("Invalid input. Exiting...")
			return
		}
	}
}

var cmd = &cobra.Command{
	Use:   "test",
	Short: "Run maelstrom test",
	Run:   handleRun,
}

func init() {
	rootCmd.AddCommand(cmd)
}
