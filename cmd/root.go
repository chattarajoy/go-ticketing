package cmd

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/chattarajoy/go-ticketing/cmd/server"
)

var (
	// GitTag stores the tag of current branch, if tag is not present it takes the branch
	GitTag string
	// GitCommit stores the most recent commit id
	GitCommit string

	// base command
	rootCmd = &cobra.Command{
		Use:   "ticketing",
		Short: "REST API for booking cinema tickets",
	}

	// config file path
	cfgFile string
	Logger  log.Logger
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	initLogger()
	// flag for specifying config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file, default is ~/.ticketing.yaml")
	// register your commands here
	server.Init(&server.CMD{RootCmd: rootCmd, Logger: Logger})
}

func initLogger() {
	Logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = Logger.Log("Git Tag: ", GitTag, "Git Commit: ", GitCommit)
}

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory and look for .edge file
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".ticketing")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
