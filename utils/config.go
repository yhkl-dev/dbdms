package utils

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Datasource *Datasource
	Token      *TokenConfig
	Server     *ServerConfig
	Gin        *ginConfig
}

type TokenConfig struct {
	Issuer      string `yaml:"issuer"`
	SignKey     string `yaml:"sign-key"`
	ActiveTime  int64  `yaml:"active-time"`
	ExpiredTime int64  `yaml:"expired-time"`
}

type ginConfig struct {
	RunMode                string `yaml:"run-mode"`
	HandleMethodNotAllowed bool   `yaml:"handle-method-not-allowed"`
	MaxMultipartMememory   int64  `yaml:"max-multipart-memory"`
}

// ServerConfig server config
type ServerConfig struct {
	Addr           string        `yaml:"addr"`
	ReadTimeout    time.Duration `yaml:"read-timeout"`
	WriteTimeout   time.Duration `yaml:"write-timeout"`
	IdleTimeout    time.Duration `yaml:"idle-timeout"`
	MaxHeaderBytes int           `yaml:"max-header-bytes"`
}

// Datasource database config
type Datasource struct {
	Driver        string `yaml:"driver"`
	URL           string `yaml:"url"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	MaxOpenConns  int    `yaml:"max-open-conns"`
	MaxIdleConns  int    `yaml:"max-idle-conns"`
	ShowSQL       bool   `yaml:"show-sql"`
	SingularTable bool   `yaml:"singular-table"`
}

var configuration *config

// GetGinConfig get gin config
func GetGinConfig() (ginconfig *ginConfig) {
	return configuration.Gin
}

// GetDatasource get data source
func GetDatasource() *Datasource {
	return configuration.Datasource
}

// LoadDatasourceConfig load config from file path
func LoadDatasourceConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return err
	}
	return err
}

// GetTokenConfig get token config
func GetTokenConfig() (tokenconfig *TokenConfig) {
	return configuration.Token
}

// LoadTokenConfig load token config
func LoadTokenConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return err
	}
	return err
}

// LoadServerConfig load server config
func LoadServerConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return err
	}
	return err
}

// GetServerConfig get server config
func GetServerConfig() (serverconfig *ServerConfig) {
	return configuration.Server
}
