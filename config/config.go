/*
   Copyright 2021 TEAM-A

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration estructura
type Configuration struct {
	APIRestServerHost string `mapstructure:"api_logic_host"`
	APIRestServerPort string `mapstructure:"api_logic_port"`
	APIRestUsername   string `mapstructure:"api_logic_username"`
	APIRestPassword   string `mapstructure:"api_logic_password"`
	EthRpcUrl         string `mapstructure:"eth_rpc_url"`
	KafkaBrokers      string `mapstructure:"kafka_brokers"`
	KafkaTopic        string `mapstructure:"kafka_topic"`
}

// Config is package struct containing conf params
var ConfMap Configuration

func Load(path string, name string, ext string) {

	// name := "parameters"
	// ext := "yml"
	// path := "./config"
	fmt.Printf("Loading configuration %s/%s.%s\n", path, name, ext)
	viper.SetConfigType(ext)
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	// Setting defaults if the config not read
	// API
	viper.SetDefault("api_logic_host", "127.0.0.1")
	viper.SetDefault("api_logic_port", "8080")
	viper.SetDefault("api_logic_username", "admin")
	viper.SetDefault("api_logic_password", "admin")
	// MongoDB
	// viper.SetDefault("api_logic_mongodb_host", "127.0.0.1")
	// viper.SetDefault("api_logic_mongodb_port", "27017")
	// viper.SetDefault("api_logic_mongodb_user", "root")
	// viper.SetDefault("api_logic_mongodb_pw", "cmc")

	// ETH
	viper.SetDefault("eth_rpc_url", "http://127.0.0.1:8545")
	// viper.SetDefault("eth_rpc_user", "")
	// viper.SetDefault("eth_rpc_pw", "")

	// Kafka
	viper.SetDefault("kafka_brokers", "")
	viper.SetDefault("kafka_topic", "")

	if _, err := os.Stat(filepath.Join(path, name+"."+ext)); err == nil {
		err = viper.ReadInConfig()
		if err == nil {
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				// TODO: load new config values ...
				log.Println("Config file changed: ", e.Name)
			})
		} else {
			log.Errorln(err)
		}
	} else {
		log.Warningf("File parameters.yml not found. Working with default config: %s \n", err)
	}

	err := viper.Unmarshal(&ConfMap)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %+v", err)
	}
	fmt.Printf("Load configuration : \n")
	spew.Dump(ConfMap)
}
