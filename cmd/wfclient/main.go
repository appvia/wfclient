package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/appvia/wfclient/pkg/client"
	"github.com/appvia/wfclient/pkg/client/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Use:   "wfclient",
	Short: "Simple example of using the wfclient library in a CLI tool",
	Long:  `This is a demonstration of how to use the wfclient library.`,
}

var serverInfoCmd = &cobra.Command{
	Use:   "serverinfo",
	Short: "Get server information",
	Long:  `Retrieve information about the Wayfinder server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile := ""
		for i, f := range os.Args {
			if f == "--profile" && len(os.Args) > i+1 {
				profile = os.Args[i+1]
			} else if strings.HasPrefix(f, "--profile=") {
				profile = strings.TrimPrefix(f, "--profile=")
			}
		}

		cfg, err := config.GetConfig()
		if err != nil {
			return err
		}

		// @step: we create an client from the configuration
		wfClient := client.NewClient(cfg, client.UseUpdateHandler(updateClientConfiguration(cfg)))
		if err != nil {
			return err
		}
		if profile != "" {
			wfClient.OverrideProfile(profile)
		}

		// Create a request to get server info
		req := wfClient.
			Request().
			Endpoint("/serverinfo")

		// Execute the request
		var serverInfo map[string]interface{}
		if err := req.
			Result(&serverInfo).
			Get().
			Error(); err != nil {
			return fmt.Errorf("failed to get server info: %w", err)
		}

		// Print the server info
		fmt.Printf("Server Information:\n")
		if err := yaml.NewEncoder(os.Stdout).Encode(serverInfo); err != nil {
			return fmt.Errorf("failed to print server info: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverInfoCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func updateClientConfiguration(cfg *config.Config) client.UpdateHandlerFunc {
	return func() error {
		if config.IsEphemeralConfig() {
			return nil
		}

		return config.UpdateConfig(cfg, config.GetClientConfigurationPath())
	}
}
