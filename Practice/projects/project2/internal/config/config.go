package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-reqiured:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
		flags:= flag.String("config","","path of configuration file")
		flag.Parse()
		configPath= *flags
		if configPath==""{
			log.Fatal("config path is not defined")
		}
	}

	// if _,err:= os.Stat(configPath); os.IsNotExist(err){
	// 	log.Fatalf("config path do not exist: %s",configPath)
	// }
	_,err:= os.Stat(configPath)
	if os.IsNotExist(err){
		log.Fatalf("config path do not exist: %s",configPath)
	}
	var cnf Config
	err= cleanenv.ReadConfig(configPath,&cnf)
	if err!=nil{
		log.Fatalf("could not read config file: %s",err.Error())
	}

	return &cnf
}	