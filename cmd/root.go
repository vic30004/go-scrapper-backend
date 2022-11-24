package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ecom-scrapper",
	Short: "An application that will scrap ecom websites",
	Long:  "An application that will scrap ecom websites. At the moment it will only scrape Best Buy and Amazon",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.subscriptions-service.yml")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		rootPath := fmt.Sprintf("%s/go/github.com/vic30004/go-scrapper-backend", home)
		viper.AddConfigPath(rootPath)
		viper.AddConfigPath(".")
		viper.AddConfigPath("..")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-scrapper-backend")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "using config file:", viper.ConfigFileUsed())
	}
}
