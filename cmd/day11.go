package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/unixlab/AoC2023/internal/aoeinput"
	"github.com/unixlab/AoC2023/internal/day11"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use: "day11",
	Run: func(cmd *cobra.Command, args []string) {
		example, err := cmd.Flags().GetBool("example")
		if err != nil {
			panic(err)
		}
		input := aoeinput.Read("", cmd.Use, example)
		fmt.Printf("day11 part 1 => %d\n", day11.Run(input, 2))
		fmt.Printf("day11 part 2 => %d\n", day11.Run(input, 1000000))
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}
