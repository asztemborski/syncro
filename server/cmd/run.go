package cmd

import (
	"github.com/asztemborski/syncro/internal"
	"github.com/spf13/cobra"
)

const (
	DefaultConfigPath = "/config/config-development.yml"
	DefaultEnvPath    = ".env"
	DefaultEnvPrefix  = ""
)

var (
	configFilePath string
	envFilePath    string
	envPrefix      string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", DefaultConfigPath, "config file path")
	rootCmd.PersistentFlags().StringVar(&envFilePath, "env", DefaultEnvPath, ".env file path")
	rootCmd.PersistentFlags().StringVar(&envPrefix, "env-prefix", DefaultEnvPrefix, "env variables prefix")

	rootCmd.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run application",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run(cmd.Context(), internal.BootstrapArgs{
			ConfigFile: configFilePath,
			EnvFile:    envFilePath,
		})
	},
}
