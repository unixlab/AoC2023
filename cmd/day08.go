package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2023/internal/aoeinput"
	"github.com/unixlab/AoC2023/internal/day08"
)

// day08Cmd represents the day08 command
var day08Cmd = &cobra.Command{
	Use: "day08",
	Run: func(cmd *cobra.Command, _ []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day := "day08"
		if example {
			day += "p1"
		}
		input := aoeinput.Read("", day, example)
		fmt.Printf("day08 part 1 => %d\n", day08.RunPart1(input))
		if example {
			day = "day08p2"
		}
		input = aoeinput.Read("", day, example)
		fmt.Printf("day08 part 2 => %d\n", day08.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day08Cmd)
}
