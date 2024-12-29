package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Log server successfully")
}
