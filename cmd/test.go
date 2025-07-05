/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

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
		var result IPerf3Result
		err = json.Unmarshal(output, &result)
		if err != nil {
			fmt.Println("Error running iperf3:", err)
			return
		}
		timestamp := time.Now().Format(time.RFC3339)

		// Open file in append mode, create if it doesn't exist
		file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v", err)
		}
		defer file.Close()

		// Create CSV writer
		writer := csv.NewWriter(file)
		defer writer.Flush()
		Mbits_per_second := int(result.End.Streams[0].Sender.BitsPerSecond / 1_000_000)

		// Write the row: timestamp, directtion, Mbits_per_second
		err = writer.Write([]string{timestamp, "Download", fmt.Sprintf("%d", Mbits_per_second)})
		if err != nil {
			fmt.Printf("Error writing to CSV: %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
