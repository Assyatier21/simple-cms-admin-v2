package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ApplicationConfig ApplicationConfig `mapstructure:"APP_CONFIG"`
	PostgresConfig    DBConfig          `mapstructure:"POSTGRESQL"`
	ElasticConfig     ElasticConfig     `mapstructure:"ELASTICSEARCH"`
	RedisConfig       RedisConfig       `mapstructure:"REDIS_CONFIG"`
	JWTSecretKey      string            `mapstructure:"JWT_SECRET_KEY"`
}

type ApplicationConfig struct {
	Host string `mapstructure:"APP_HOST"`
	Port string `mapstructure:"APP_PORT"`
}

type DBConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	Database string `mapstructure:"POSTGRES_DB"`
	Schema   string `mapstructure:"POSTGRES_SCHEMA"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
}

type ElasticConfig struct {
	Address       string `mapstructure:"ESADDRESS"`
	IndexArticle  string `mapstructure:"ES_INDEX_ARTICLE"`
	IndexCategory string `mapstructure:"ES_INDEX_CATEGORY"`
}

type RedisConfig struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     string `mapstructure:"REDIS_PORT"`
	Username string `mapstructure:"REDIS_USERNAME"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

func Load() (conf Config) {
	viper.SetConfigFile("config")
	viper.SetConfigFile("./config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return
}

func (db *DBConfig) GetDSN() (dsn string) {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", db.User, db.Password, db.Host, db.Database, db.Schema)
}
