package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kolanos/dwolla-v2-go"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var client *dwolla.Client
var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dwolla",
	Short: "A CLI for the Dwolla v2 API",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	defaultConfigFile := filepath.Join(homeDir(), ".dwolla.yml")

	persistentFlags := rootCmd.PersistentFlags()

	persistentFlags.StringVar(&configFile, "config", "", fmt.Sprintf("config file (default: %s)", defaultConfigFile))

	persistentFlags.String("key", "", "Override configured Dwolla API Key")
	viper.BindPFlag("apiKey", persistentFlags.Lookup("key"))

	persistentFlags.String("secret", "", "Override configured Dwolla API secret")
	viper.BindPFlag("apiSecret", persistentFlags.Lookup("secret"))

	persistentFlags.String("environment", "", "Override configured Dwolla API environment")
	viper.BindPFlag("environment", persistentFlags.Lookup("environment"))

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initClient() {
	if client != nil {
		return
	}

	var environment dwolla.Environment

	if viper.GetString("environment") == "production" {
		environment = dwolla.Production
	}

	if viper.GetString("environment") == "sandbox" {
		environment = dwolla.Sandbox
	}

	client = dwolla.New(viper.GetString("apiKey"), viper.GetString("apiSecret"), environment)
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home := homeDir()
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".dwolla")
	}

	viper.SetDefault("apiKey", "")
	viper.SetDefault("apiSecret", "")
	viper.SetDefault("environment", dwolla.Production)

	viper.SetEnvPrefix("dwolla")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func homeDir() string {
	home, err := homedir.Dir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return home
}
