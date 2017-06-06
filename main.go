package main

import (
	"flag"
	"os"

	"github.com/hyperpilotio/slow_cooker_wrapper/config"
	"github.com/hyperpilotio/slow_cooker_wrapper/glue"
)

func main() {
	// load config
	configPath := flag.String("c", "", "path to config file, please find example in config.json")
	flag.Parse()
	if *configPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	var conf config.Config = config.ViperConfig{}
	conf.LoadConfig(configPath)

	// run slow cooker
	loaded := glue.LoadParameters(conf)
	// convert to Parameters
	params := glue.Convert(&loaded)
	params.Runs()
}
