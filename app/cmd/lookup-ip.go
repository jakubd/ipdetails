package cmd

import (
	"github.com/jakubd/ipdetails"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ip <given ip>",
	Short: "get info on a single ip",
	Long:  "get information on a single ip",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		LookupOne(args[0])
	},
}

func LookupOne(ip string) {
	ipdetails.OutputLookup(ip, false)
}

func Execute() error {
	cobra.MinimumNArgs(1)
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(pipeCmd)
}
