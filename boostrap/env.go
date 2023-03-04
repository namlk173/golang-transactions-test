package boostrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
}

func NewEnv() *Env {
	var env Env
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Can't find the .env file.")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalln("Environment can't be loaded.")
	}

	return &env
}
