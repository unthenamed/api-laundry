package config

import "fmt"

type Database struct {
	HostName     string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Driver       string
}

type Api struct {
	Port string
}

type Config struct {
	Database Database
	Api      Api
}

func (c *Config) readConfig() {
	c.Api.Port = "8080"
	c.Database = Database{
		HostName:     "localhost",
		Port:         5432,
		User:         "postges",
		Password:     "jaliL5525",
		DatabaseName: "api-laundry",
		Driver:       "postgres",
	}
}

func (c *Config) validateConfig() error {
	if c.Api.Port == "" || c.Database.HostName == "" || c.Database.Port == 0 || c.Database.User == "" || c.Database.Password == "" || c.Database.DatabaseName == "" || c.Database.Driver == "" {
		err := fmt.Errorf("invalid configuration")
		panic(err)
	}

	c.readConfig()
	return nil
}

func NewConfig() (config *Config, err error) {
	err = config.validateConfig()
	return
}
