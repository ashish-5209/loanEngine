package config

import (
	"fmt"
	"os"
	"time"
)

var (
	allEnvironments map[string]struct{}
	Environment     string
	ISTLoc          *time.Location
	Domain          string
	AppPath         string
)

func init() {
	allEnvironments = map[string]struct{}{"dev": {}, "pp": {}, "prod": {}}
	Environment = getEnv()
	fmt.Println("env :", Environment)

	var err error
	ISTLoc, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	// for domain name of app
	Domain = "amamrthaloan.in"
	if Environment == "dev" {
		Domain = "localhost"
	}
	// for pdf generation path
	AppPath = os.Getenv("loanengine_path")
}

func getEnv() string {
	env := os.Getenv("ENV")
	if _, found := allEnvironments[env]; found {
		return env
	}
	// default env
	return "dev"
}
