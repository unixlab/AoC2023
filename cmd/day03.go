package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2023/internal/aoeinput"
	"github.com/unixlab/AoC2023/internal/day03"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use: "day03",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day03 part 1 => %d\n", day03.RunPart1(input))
		fmt.Printf("day03 part 2 => %d\n", day03.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day03Cmd)
}
