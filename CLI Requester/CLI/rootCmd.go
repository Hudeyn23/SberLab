package CLI

import (
	"CLI/CLI/Configuration"
	"github.com/spf13/cobra"
)

var offset string
var limit string
var configPath string
var config Configuration.Config

var rootCmd = &cobra.Command{
	Use:   "sber",
	Short: "CLI utility for SberCloud",
	Long:  "CLI utility for SberCloud with api requests",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config, _ = Configuration.LoadConfig(configPath)
	},
}

func Start() {
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(cmdECS)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", ".", "Path to the folder with config")
	cmdECS.AddCommand(cmdShowECS)
	cmdShowECS.PersistentFlags().StringVarP(&offset, "offset", "o", "1", "Number of first entity")
	cmdShowECS.PersistentFlags().StringVarP(&limit, "limit", "l", "1", "Number of entities")
}
