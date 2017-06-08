package glue

import (
	"testing"

	"github.com/hyperpilotio/slow_cooker_wrapper/config"
	"github.com/hyperpilotio/slow_cooker_wrapper/slowcooker"
	"github.com/hyperpilotio/slow_cooker_wrapper/utils"
	"github.com/stretchr/testify/assert"
)

func TestDefaultWrapperGlue(t *testing.T) {
	var c config.Config = config.ViperConfig{}
	configPath := "./testconfig.json"
	c.LoadConfig(&configPath)
	pars := LoadParameters(c)
	i := utils.First(pars, func(arg2 interface{}) bool {
		return arg2.(slowcooker.Parameter).Title == "step 1a"
	})

	assert.Equal(t, "influxdb", i.(slowcooker.Parameter).MetricServerBackend, "not load properly")
	assert.Equal(t, "http://localhost:8086", i.(slowcooker.Parameter).MetricAddr, "not load properly")
	assert.Equal(t, "root", i.(slowcooker.Parameter).InfluxUserName, "not load properly")
	assert.Equal(t, "root", i.(slowcooker.Parameter).InfluxPassword, "not load properly")
	assert.Equal(t, "metrics", i.(slowcooker.Parameter).InfluxDatabase, "not load properly")

	j := utils.First(pars, func(arg2 interface{}) bool {
		return arg2.(slowcooker.Parameter).Title == "step 1b"
	})

	assert.Equal(t, "", j.(slowcooker.Parameter).InfluxDatabase, "not load properly")
	assert.Equal(t, "r1", j.(slowcooker.Parameter).InfluxUserName, "not load properly")
	assert.Equal(t, "r2", j.(slowcooker.Parameter).InfluxPassword, "not load properly")
	assert.Equal(t, "http://172.17.0.3:8086", j.(slowcooker.Parameter).MetricAddr, "not load properly")
}
