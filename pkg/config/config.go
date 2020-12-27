/*
 * telegram: @VasylNaumenko
 */

package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
)

type (
	// Cfg defines the properties of the application configuration
	Cfg struct {
		APIServer   Server      `yaml:"api-server"`
		RestVendors RestVendors `yaml:"rest-vendors"`
		Logger      Logger      `yaml:"logger"`
	}

	// Server defines API server configuration
	Server struct {
		HTTP HTTP `yaml:"http"`
	}

	// HTTP defines HTTP section of the API server configuration
	HTTP struct {
		ListenAddr string `yaml:"listen-address"`
	}

	// RestVendors defines REST Vendors API section configuration
	RestVendors struct {
		Osrm string `yaml:"osrm"`
	}

	// Logger defines logger section of the API server configuration
	Logger struct {
		OutputFilePath        string `yaml:"output-file-path"`
		DebugLevel            string `yaml:"debug-level"`
		LogFormat             string `yaml:"log-format"`
		IncludeCallerMethod   bool   `yaml:"include-caller-method"`
		RequestOutputFilePath string `yaml:"requests-log-output-file-path"`
	}
)

const (
	errMsgReadConfig = "Unable to read the '%s' configuration file"
)

// Init loads and validates all configuration data
func Init(cfgFilePath string) (*Cfg, error) {
	cfg := &Cfg{}
	if err := Load(cfg, cfgFilePath); err != nil {
		return nil, err
	}

	populateConfigDefaults(cfg)

	return cfg, nil
}

func populateConfigDefaults(cfg *Cfg) {
	for ref, defaultValue := range map[*string]string{
		&cfg.Logger.DebugLevel: defaultLoggerDebugLevel,
		&cfg.Logger.LogFormat:  defaultLogFormat,
	} {
		checkStringParam(ref, defaultValue)
	}
}

func checkStringParam(vCurrent *string, vDefault string) {
	if len(*vCurrent) == 0 {
		*vCurrent = vDefault
	}
}

// Reads configuration data from specified yaml file path
func Load(cfg interface{}, configPath string) error {
	if err := readConfigFile(configPath, cfg); err != nil {
		return err
	}

	return nil
}

func readConfigFile(name string, cfg interface{}) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return errors.Wrap(err, "read config file error")
	}

	conf, err := os.Open(name)
	if err != nil {
		return errors.WithMessagef(err, errMsgReadConfig, name)
	}

	data, err := ioutil.ReadAll(conf)
	if err != nil {
		return errors.WithMessagef(err, errMsgReadConfig, name)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return errors.WithMessagef(err, errMsgReadConfig, name)
	}

	return nil
}
