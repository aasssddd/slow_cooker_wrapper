package config

import (
	"fmt"
	"testing"

	"github.com/hyperpilotio/slow_cooker_wrapper/utils"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	var conf Config = ViperConfig{}
	var tfpath = "./test.json"

	err := conf.LoadConfig(&tfpath)
	assert.NoError(t, err, "suppose no error")

	var tfpath2 = "../test.json"
	err2 := conf.LoadConfig(&tfpath2)
	assert.Error(t, err2, "should raise error")

	v := conf.Get("jobs")
	utils.Each(v.([]interface{}), func(arg2 interface{}) interface{} {
		fmt.Printf("Type: %T \nValue: %v\n", arg2, arg2)
		return arg2
	})
}
