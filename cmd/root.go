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

	persistentFlags.StringVarP(&configFile, "config", "c", "", fmt.Sprintf("config file (default: %s)", defaultConfigFile))

	persistentFlags.StringP("key", "k", "", "dwolla api key")
	viper.BindPFlag("apiKey", persistentFlags.Lookup("key"))

	persistentFlags.StringP("secret", "s", "", "dwolla api secret")
	viper.BindPFlag("apiSecret", persistentFlags.Lookup("secret"))

	persistentFlags.StringP("environment", "e", "", "dwolla api environment (default: production)")
	viper.BindPFlag("environment", persistentFlags.Lookup("environment"))

	persistentFlags.BoolP("verbose", "v", false, "increase verbosity")
	viper.BindPFlag("verbose", persistentFlags.Lookup("verbose"))
}

func initClient() {
	if client != nil {
		return
	}

	if !viper.IsSet("apiKey") {
		fmt.Println("No dwolla api key configured")
		os.Exit(1)
	}

	if !viper.IsSet("apiSecret") {
		fmt.Println("No dwolla api secret configured")
		os.Exit(1)
	}

	var environment dwolla.Environment

	switch viper.GetString("environment") {
	case "production":
		environment = dwolla.Production
	case "sandbox":
		environment = dwolla.Sandbox
	default:
		fmt.Println("Invalid dwolla environment:", viper.GetString("environment"))
		os.Exit(1)
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
		if viper.GetBool("verbose") {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
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
