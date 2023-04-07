package configs

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"
)

type Config struct {
	Address        string        `mapstructure:"ADDRESS"`
	PollInterval   time.Duration `mapstructure:"POLL_INTERVAL"`
	ReportInterval time.Duration `mapstructure:"REPORT_INTERVAL"`
	StoreInterval  time.Duration `mapstructure:"STORE_INTERVAL"`
	StoreFile      string        `mapstructure:"STORE_FILE"`
	Restore        bool          `mapstructure:"RESTORE"`
	Key            string        `mapstructure:"KEY"`
}

var CFG = &Config{}

// ReadOs reads config from environment variables
// This func will replace Config parameters if any presented in os environment vars
func (config *Config) ReadOs() {
	// load config from environment variables
	v := viper.New()
	v.AutomaticEnv()
	if v.Get("ADDRESS") != nil {
		config.Address = v.GetString("ADDRESS")
	}
	if v.Get("POLL_INTERVAL") != nil {
		config.PollInterval = v.GetDuration("POLL_INTERVAL")
	}
	if v.Get("REPORT_INTERVAL") != nil {
		config.ReportInterval = v.GetDuration("REPORT_INTERVAL")
	}
	if v.Get("STORE_INTERVAL") != nil {
		config.StoreInterval = v.GetDuration("STORE_INTERVAL")
	}
	if v.Get("STORE_FILE") != nil {
		config.StoreFile = v.GetString("STORE_FILE")
	}
	if v.Get("RESTORE") != nil {
		config.Restore = v.GetBool("RESTORE")
	}
	if v.Get("KEY") != nil {
		config.Key = v.GetString("KEY")
	}
}

// InitFiles creates all necessary files and folders for server storage
func (config *Config) InitFiles() {
	// get dir of the file
	dr := path.Dir(config.StoreFile)
	// check if dir exists
	if _, err := os.Stat(dr); os.IsNotExist(err) {
		// create dir
		err = os.MkdirAll(dr, os.ModePerm)
		if err != nil {
			log.Fatal("error creating dir: ", err)
		}
	}
}

// ReadServerFlags reads config from flags Run this first
func (config *Config) ReadServerFlags() {
	// read flags
	flag.StringVar(&config.Address, "a", "localhost:8080", "server address")
	flag.DurationVar(&config.StoreInterval, "i", 300*time.Second, "store interval")
	flag.StringVar(&config.StoreFile, "f", "/tmp/devops-metrics-db.json", "store file")
	flag.StringVar(&config.Key, "k", "123", "hash key")
	flag.BoolVar(&config.Restore, "r", true, "restore")
	flag.Parse()
}

// ReadAgentFlags separate function required bec of similar variable names required for agent and server
func (config *Config) ReadAgentFlags() {
	// read flags
	flag.StringVar(&config.Address, "a", "localhost:8080", "server address")
	flag.StringVar(&config.Key, "k", "", "hash key")
	flag.DurationVar(&config.PollInterval, "p", 1*time.Second, "poll interval")
	flag.DurationVar(&config.ReportInterval, "r", 2*time.Second, "report interval")
	flag.Parse()
}
