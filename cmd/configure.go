package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the bandgauge cli tool",
	Long: `Configure the bandgauge CLI tool to work with your network setup.
This command allows you to set up the necessary parameters such as the iperf3 server remote host adress, port, and
other configurations required for the tool to function correctly.`,
	Run: func(cmd *cobra.Command, args []string) {
		var remote_host string
		var port string
		fmt.Print("Enter the host (IP address) of the remote iperf3 server:")
		fmt.Scan(&remote_host)
		fmt.Print("Enter the port of the remote iperf3 server: ")
		fmt.Scan(&port)
		cfg := Config{
			RemoteHost: remote_host,
			Port:       port,
		}
		if err := saveConfig(cfg); err != nil {
			fmt.Println("Error saving configuration:", err)
			return
		}
		fmt.Println("Configuration saved successfully.")
	},
}

type Config struct {
	RemoteHost string `json:"remote_host"`
	Port       string `json:"port"`
}

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(configDir, "bandgauge")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "config.json"), nil
}

func saveConfig(cfg Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(cfg)
}

func LoadConfig() (Config, error) {
	var cfg Config
	path, err := getConfigPath()
	if err != nil {
		return cfg, err
	}

	f, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&cfg)
	return cfg, err
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
