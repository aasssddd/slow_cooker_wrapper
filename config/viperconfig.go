package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// ViperConfig : config obj
type ViperConfig struct {
}

// LoadConfig : load config fire from path
func (c ViperConfig) LoadConfig(vfile *string) error {
	fileInfo, err := os.Stat(*vfile)
	if err != nil {
		return err
	}
	viper.SetConfigName(strings.Replace(fileInfo.Name(), path.Ext(*vfile), "", -1))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath(path.Dir(*vfile))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// Get : get generic property value
func (c ViperConfig) Get(key string) interface{} {
	fmt.Println("all keys", viper.AllKeys())
	return viper.Get(key)
}
