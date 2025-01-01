package config

import "fmt"

type Database struct {
	HostName     string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Driver       string
	DSN          string
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
		User:         "postgres",
		Password:     "5525",
		DatabaseName: "api-laundry",
		Driver:       "postgres",
	}
}

func (c *Config) validateConfig() error {

	c.readConfig()

	if c.Api.Port == "" || c.Database.HostName == "" || c.Database.Port == 0 || c.Database.User == "" || c.Database.Password == "" || c.Database.DatabaseName == "" || c.Database.Driver == "" {
		err := fmt.Errorf("invalid configuration")
		panic(err)
	}

	c.Database.DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Database.HostName, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DatabaseName)
	return nil
}

func NewConfig() (config Config, err error) {
	err = config.validateConfig()
	return
}
