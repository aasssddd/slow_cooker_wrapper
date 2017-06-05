package main

import (
	"flag"
	"os"

	"github.com/hyperpilotio/slow_cook_wrapper/config"
	"github.com/hyperpilotio/slow_cook_wrapper/slowcooker"
	"github.com/hyperpilotio/slow_cook_wrapper/utils"
)

func loadParameters(c config.Config) slowcooker.Parameters {
	result := make([]slowcooker.Parameter, 0)
	utils.Each(c.Get("jobs").([]interface{}), func(arg2 interface{}) interface{} {
		var p slowcooker.Parameter
		ins := arg2.(map[string]interface{})
		if v, ok := ins["compress"].(bool); ok {
			p.Compress = v
		}
		if v, ok := ins["concurrency"].(float64); ok {
			p.Concurrency = int(v)
		}
		if v, ok := ins["concurrency"].(float64); ok {
			p.Concurrency = int(v)
		}
		if v, ok := ins["data"].(string); ok {
			p.Data = v
		}
		if v, ok := ins["hashSampleRate"].(float64); ok {
			p.HashSampleRate = v
		}
		if v, ok := ins["HashValue"].(uint64); ok {
			p.HashValue = v
		}
		if v, ok := ins["header"].(map[string]string); ok {
			p.Header = v
		}
		if v, ok := ins["host"].(string); ok {
			p.Host = v
		}
		if v, ok := ins["interval"].(string); ok {
			p.Interval = v
		}
		if v, ok := ins["method"].(string); ok {
			p.Method = v
		}
		if v, ok := ins["metricAddr"].(string); ok {
			p.MetricAddr = v
		}
		if v, ok := ins["noLatencySummary"].(bool); ok {
			p.NoLatencySummary = v
		}
		if v, ok := ins["noreuses"].(bool); ok {
			p.Noreuses = v
		}
		if v, ok := ins["qps"].(float64); ok {
			p.Qps = int(v)
		}
		if v, ok := ins["metricAddr"].(string); ok {
			p.MetricAddr = v
		}
		if v, ok := ins["reportLatenciesCSV"].(string); ok {
			p.ReportLatenciesCSV = v
		}
		if v, ok := ins["runOrder"].(float64); ok {
			p.RunOrder = int(v)
		}
		if v, ok := ins["target"].(string); ok {
			p.Target = v
		}
		if v, ok := ins["title"].(string); ok {
			p.Title = v
		}
		if v, ok := ins["totalRequests"].(float64); ok {
			p.TotalRequests = int(v)
		}

		result = append(result, p)
		return arg2
	})
	return result
}

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
	params := loadParameters(conf)
	params.Runs()
}
