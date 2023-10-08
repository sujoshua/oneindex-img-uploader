package uploader

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"oneindex-img-uploader/pkg/fileutil"
	"oneindex-img-uploader/pkg/uploaderFactory"
	"os"
	"path/filepath"
)

var Config = &Option{}

type Option struct {
	Url        string // image url
	Platform   string // image platform
	Uploader   uploaderFactory.Uploader
	ConfigFile string
}

func bindFlag(flag *pflag.FlagSet, key string) {
	viper.BindPFlag(key, flag.Lookup(key))
}

func InitConfig(rootCmd *cobra.Command) {
	cobra.OnInitialize(initConfigFile, bindConfigFile)

	flag := rootCmd.PersistentFlags()

	flag.StringVar(&Config.ConfigFile, "config", "", "config file (default is $HOME/.config/img-uploader.yaml)")
	bindFlag(flag, "config")

	flag.StringVar(&Config.Platform, "platform", "", "image platform")
	_ = rootCmd.MarkFlagRequired("platform") // mark required
	bindFlag(flag, "platform")

	flag.StringVar(&Config.Url, "url", "", "image url")
	bindFlag(flag, "url")
}

func bindConfigFile() {
	Config.Platform = viper.GetString("platform")
	Config.Url = viper.GetString("url")
}

func initConfigFile() {
	if Config.ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(Config.ConfigFile)
	} else {
		home := fileutil.UserHomeDir()

		configDir := filepath.Join(home, ".config")
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			err = os.Mkdir(configDir, 0755)
			if err != nil {
				log.Printf("create config dir: %s failed: %s", configDir, err.Error())
				os.Exit(1)
			}
		}

		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("img-uploader")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
