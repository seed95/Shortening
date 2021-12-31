package cmd

import (
	"espad_task/application"
	"github.com/spf13/cobra"
	"log"
)

var (
	configFile = "./build/config/config.yaml"

	rootCmd = &cobra.Command{
		Use:   "shortening",
		Short: "Shortening is a very simple service for shortening original links",
		Long: `A very simple shortening service with MD5 for hash algorithm and base64 for encoding hash to represent to user.

Source code https://github.com/seed95/Shortening`,
		Run: func(cmd *cobra.Command, args []string) {
			opt := &application.Option{
				ConfigFile: configFile,
			}

			if err := application.Run(opt); err != nil {
				log.Fatalln(err)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", configFile, "config file")
}
