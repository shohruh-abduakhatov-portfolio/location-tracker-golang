package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config loads configugations
type Config struct {
	Server struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"server"`
	Mongo struct {
		Protocol string `json:"protocol"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		DBName   string `json:"db_name"`
	} `json:"mongo"`
	RabbitMQ struct {
		Protocol string `json:"protocol"`
		Username string `json:"username"`
		Password string `json:"password"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
	} `json:"rabbitmq"`
}

var (
	// Cfg config store
	Cfg *Config
)

func init() {
	Cfg = newConfig()
}

func newConfig() *Config {
	conf := new(Config)

	// read file
	var f *os.File

	f, err := os.Open(filepath.Join("config", "prod.json"))
	if err != nil {
		f, err = os.Open(filepath.Join("config", "dev.json"))
	}
	processError(err)
	defer f.Close()

	// parse and unmarshall
	store, _ := ioutil.ReadAll(f)
	json.Unmarshal(store, conf)

	return conf
}

func processError(err error) {
	if err == nil {
		return
	}
	log.Println(err)
	os.Exit(2)
}
