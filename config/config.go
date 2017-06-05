package config

// Config : config interface
type Config interface {
	LoadConfig(path *string) error
	Get(key string) interface{}
}
