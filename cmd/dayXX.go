package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2023/internal/aoeinput"
	"github.com/unixlab/AoC2023/internal/dayXX"
)

// dayXXCmd represents the dayXX command
var dayXXCmd = &cobra.Command{
	Use: "dayXX",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("dayXX part 1 => %d\n", dayXX.RunPart1(input))
		fmt.Printf("dayXX part 2 => %d\n", dayXX.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(dayXXCmd)
}
