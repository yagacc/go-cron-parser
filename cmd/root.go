package cmd

import (
	"errors"
	"fmt"
	"github.com/mryan321/go-cron-parser/parser"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-cron-parser",
	Short: "A cron expression parser",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("error: requires a single cron expression argument")
		}
		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		p := parser.CronExpressionParser{Expression: args[0]}
		cron, err := p.Parse()
		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			fmt.Printf("%s", cron)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
