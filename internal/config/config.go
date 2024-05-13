package config

import (
	"log"
	"os"
)

type Config struct {
	PORT           string
	HOST           string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
}

var Configuration Config

func (config *Config) LoadConfig() {
	config.PORT = os.Getenv("PORT")
	config.HOST = os.Getenv("HOST")
	config.DBUserName = os.Getenv("DBUserName")
	config.DBUserPassword = os.Getenv("DBUserPassword")
	config.DBName = os.Getenv("DBName")
	config.DBPort = os.Getenv("DBPort")
}

func (config *Config) DisplayConfig() {
	log.Printf("PORT %v, HOST %v, DBUserName %v, DBUserPassword %v, DBName %v , DBPort %v", Configuration.PORT, Configuration.HOST, Configuration.DBUserName, Configuration.DBUserPassword, Configuration.DBName, Configuration.DBPort)
}
