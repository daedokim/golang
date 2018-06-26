package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//Config Start
type Config struct {
	Debug    bool
	Database struct {
		Driver     string
		Connection string
	}
	Host string
	Port string
}

//LoadConfig Start
func (Conf *Config) LoadConfig() {
	env := os.Getenv("GOENV")
	var confile string
	if env == "" {
		confile = "./config/config.dev.yml"
	} else if env == "prod" {
		confile = "./config/config.yml"
	}
	file, err := os.Open(confile)
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	defer file.Close()
	viper.MergeConfig(file)
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	fmt.Printf("--- Load config from %s ---\n", confile)
}
