package cmd

import (
	"fmt"
	"nets/pkg"
	"nets/tools"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:     "scan",
	Short:   "scan http(s) port",
	Example: "nets scan 1.1.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		var scanner tools.Scanner
		for _, str := range args {
			ip, port, err := pkg.StringToSocket(str)
			if err == nil {
				scanner.Add(ip, port)
			}
		}
		if scanner.Len() == 0 {
			pkg.ExitWithError("Please enter at least one valid URL, domain name, or IP address")
		}
		results := scanner.Run()
		fmt.Printf("%+v", results)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
