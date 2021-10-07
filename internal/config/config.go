package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type 	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

// Init populates HTTPConfig struct with values from env file and environment variables.
func Init() (*HTTPConfig, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("no .env file found: %s",err)
	}
	mhm, err := strconv.Atoi(os.Getenv("HTTP_MAX_HEADER_BYTES"))
	if err != nil {
		return nil, fmt.Errorf("error with variable HTTP_MAX_HEADER_BYTES: %s",err)
	}
	wt, err := time.ParseDuration(os.Getenv("HTTP_WRITE_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error with variable HTTP_WRITE_TIMEOUT: %s",err)
	}
	hrt, err := time.ParseDuration(os.Getenv("HTTP_READ_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("error with variable HTTP_READ_TIMEOUT: %s",err)
	}
	httpC := HTTPConfig{
		Host:               os.Getenv("HTTP_HOST"),
		Port:               os.Getenv("HTTP_PORT"),
		ReadTimeout:        hrt,
		WriteTimeout:       wt,
		MaxHeaderMegabytes: mhm,
	}
	fmt.Println("config of server: ")
	fmt.Printf("%+v\n", httpC)
	return &httpC, err
}