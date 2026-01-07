package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defauleHomeDir    = ".miniblog"
	defaultConfigName = "mb-apiserver.yaml"
)

func onInitialize() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		for _, dir := range searchDirs() {
			viper.AddConfigPath(dir)
		}
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName(defaultConfigName)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
	log.Printf("Using config file:%s", viper.ConfigFileUsed())
}

func setupEnviromentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MINIBLOG")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func searchDirs() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defauleHomeDir), "."}
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defauleHomeDir, defaultConfigName)
}
