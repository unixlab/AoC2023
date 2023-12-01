package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2023/internal/aoeinput"
	"github.com/unixlab/AoC2023/internal/day01"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use: "day01",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		day := "day01"
		if example {
			day += "p1"
		}
		input := aoeinput.Read("", day, example)
		fmt.Printf("day01 part 1 => %d\n", day01.RunPart1(input))
		if example {
			day = "day01p2"
		}
		input = aoeinput.Read("", day, example)
		fmt.Printf("day01 part 2 => %d\n", day01.RunPart2(input))
	},
}

func init() {
	rootCmd.AddCommand(day01Cmd)
}
