package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

var (
	Host            string
	Port            string
	RESTPort        string
	MatchingTimeout time.Duration
	BoardLen        int
	DBName          string
	DBHost          string
	DBUser          string
	DBPassword      string
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// Find the project root by walking upward until .infra/config.json exists.
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error getting working directory: %s", err))
	}

	for {
		configPath := filepath.Join(dir, ".infra", "config.json")

		if _, err := os.Stat(configPath); err == nil {
			viper.AddConfigPath(filepath.Join(dir, ".infra"))
			break
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			panic(fmt.Errorf("fatal error config file: .infra/config.json not found"))
		}

		dir = parent
	}

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	Host = viper.GetString("host.address")
	Port = viper.GetString("host.game_server_port")
	RESTPort = viper.GetString("host.rest_server_port")

	MatchingTimeout = time.Duration(viper.GetInt("game.matching_timeout")) * time.Second

	BoardLen = 8

	DBName = viper.GetString("database.name")
	DBHost = viper.GetString("database.host")
	DBUser = viper.GetString("database.user")
	DBPassword = viper.GetString("database.password")

	if value := os.Getenv("DB_NAME"); value != "" {
		DBName = value
	}

	if value := os.Getenv("DB_HOST"); value != "" {
		DBHost = value
	}

	if value := os.Getenv("DB_USER"); value != "" {
		DBUser = value
	}

	if value := os.Getenv("DB_PASSWORD"); value != "" {
		DBPassword = value
	}
}
