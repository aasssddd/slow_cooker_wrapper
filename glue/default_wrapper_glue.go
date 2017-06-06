package glue

import (
	"log"

	"github.com/hyperpilotio/slow_cooker_wrapper/config"
	"github.com/hyperpilotio/slow_cooker_wrapper/slowcooker"
	"github.com/hyperpilotio/slow_cooker_wrapper/utils"
)

type dbconfig struct {
	Name     string
	Host     string
	User     string
	Password string
	Database string
}

// LoadParameters :
func LoadParameters(c config.Config) []interface{} {
	result := make([]interface{}, 0)
	db := make([]interface{}, 0)
	loadDBConfig(c, &db)
	loadJobConfig(c, &db, &result)
	return result
}

func loadDBConfig(c config.Config, db *[]interface{}) {
	utils.Each(c.Get("influxDB").([]interface{}), func(arg2 interface{}) interface{} {
		c := new(dbconfig)
		ins, ok := arg2.(map[string]interface{})
		if !ok {
			log.Fatal("wrong format")
		}
		name, okay := ins["name"].(string)

		if !okay {
			log.Fatal("influxDB config: name must provided")
		}
		c.Name = name

		if host, ok := ins["host"].(string); ok {
			c.Host = host
		}

		if user, ok := ins["user"].(string); ok {
			c.User = user
		}

		if password, ok := ins["password"].(string); ok {
			c.Password = password
		}

		if database, ok := ins["database"].(string); ok {
			c.Database = database
		}

		*db = append(*db, *c)

		return arg2
	})
}

func loadJobConfig(c config.Config, db *[]interface{}, result *[]interface{}) {
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

		if database, ok := ins["database"].(string); ok {
			dbc := utils.First(*db, func(arg2 interface{}) bool {
				return arg2.(dbconfig).Name == database
			}).(dbconfig)
			p.UseInfluxDB = true
			p.InfluxUserName = dbc.User
			p.InfluxPassword = dbc.Password
			p.InfluxDatabase = dbc.Database
			p.MetricAddr = dbc.Host
		}
		if influxDatabase, ok := ins["influxDatabase"].(string); ok {
			p.InfluxDatabase = influxDatabase
		}

		if influxUserName, ok := ins["influxUserName"].(string); ok {
			p.InfluxUserName = influxUserName
		}

		if influxPassword, ok := ins["influxPassword"].(string); ok {
			p.InfluxPassword = influxPassword
		}
		if useInflux, ok := ins["useInflux"].(bool); ok {
			p.UseInfluxDB = useInflux
		}

		*result = append(*result, p)
		return arg2
	})
}

func Convert(i *[]interface{}) slowcooker.Parameters {
	var params slowcooker.Parameters
	for _, p := range *i {
		params = append(params, p.(slowcooker.Parameter))
	}
	return params
}
