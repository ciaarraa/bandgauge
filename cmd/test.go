/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run a bandwidth test",
	Long: `Run a bandwith test using iperf3. You must have an iperf3 server running on the remote host and
	ip address configured.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := LoadConfig()
		if err != nil {
			fmt.Println("Error loading configuration:", err)
			return
		}
		run := exec.Command("iperf3", "-c", config.RemoteHost, "-p", config.Port, "-J")
		output, err := run.CombinedOutput()
		if err != nil {
			fmt.Println("Error running iperf3:", err)
			return
		}
		fmt.Println("Output from iperf3:", string(output))
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
