package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DbHost         string `mapstructure:"DB_HOST"`
	DbUser         string `mapstructure:"DB_USER"`
	DbPass         string `mapstructure:"DB_PASS"`
	DbPort         string `mapstructure:"DB_PORT"`
	DbName         string `mapstructure:"DB_NAME"`
	MigrationDir   string `mapstructure:"MIGRATION_DIR"`
	RestDbSchema   string `mapstructure:"REST_DB_SCHEMA"`
	RestDbAnonRole string `mapstructure:"REST_DB_ANON_ROLE"`
	RestJwtSecret  string `mapstructure:"REST_JWT_SECRET"`
	RestServerPort string `mapstructure:"REST_SERVER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("gpm")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func CreateEmptyConfigFile(path string) error {
	fp := filepath.Join(path, "gpm.env")
	file, err := os.Create(fp)
	if err != nil {

		return err
	}
	defer file.Close()

	fmt.Println("File created successfully!")
	return nil
}
