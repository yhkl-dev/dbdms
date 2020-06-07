package system

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	Datasource *datasource
	Server     *serverConfig
	Token      *tokenConfig
	Gin        *ginConfig
}

type tokenConfig struct {
	Issuer     string `yaml:"issuer"`
	SignKey    string `yaml: "sign-key"`
	ActiveTime string `yaml: "active-time"`
	ExpireTime int64  `yaml:"expired-time"`
}

type ginConfig struct {
	RunMode                string `yaml:"run-mode"`
	HandleMethodNotAllowed bool   `yaml:"handle-method-not-allowed"`
	MaxMultipartMememory   int64  `yaml:"max-multipart-memory"`
}

type serverConfig struct {
	Addr           string        `yaml:"addr"`
	ReadTimeout    time.Duration `yaml:"read-timeout"`
	WriteTimeout   time.Duration `yaml:"write-timeout"`
	IdleTimeout    time.Duration `yaml:"idle-timeout"`
	MaxHeaderBytes int           `yaml:"max-header-bytes"`
}

type datasource struct {
	Driver        string `yaml:"driver"`
	URL           string `yaml:"url"`
	Username      string `yaml:"username"`
	Password      string `yaml: "password"`
	MaxOpenConns  int    `yaml:"max-open-conns"`
	ShowSQL       bool   `yaml:"show-sql"`
	SingularTable bool   `yaml:"singular-table"`
}

var configuration *config

// LoadDatasourceCOnfig load config from file path
func LoadDatasourceCOnfig(path string) error {
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

// GetDatasource get data source
func GetDatasource() (ds *datasource) {
	return configuration.Datasource
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
func GetServerConfig() (serverconfig *serverConfig) {
	return configuration.Server
}

// GetGinConfig get gin config
func GetGinConfig() (ginconfig *ginConfig) {
	return configuration.Gin
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

// GetTokenConfig get token config
func GetTokenConfig() (tokenconfig *tokenConfig) {
	return configuration.Token
}

// LoadConfig load config from file
func LoadConfig(path string) error {
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

// GetConfig return config
func GetConfig() *config {
	return configuration
}
