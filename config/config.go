package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fgunawan1995/bcg/util"
)

var MainConfig *Config

type Config struct {
	DB struct {
		Host    string `json:"host"`
		Port    string `json:"port"`
		User    string `json:"user"`
		Pass    string `json:"pass"`
		Name    string `json:"name"`
		SSLMode string `json:"sslmode"`
	} `json:"db"`
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}

func GetConfig() *Config {
	return MainConfig
}

func init() {
	var byteValue []byte
	jsonFile, err := os.Open(fmt.Sprintf("./config/%s.json", util.GetEnv()))
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(byteValue), &MainConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
