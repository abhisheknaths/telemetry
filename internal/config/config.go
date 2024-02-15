package config

import (
	"github.com/kelseyhightower/envconfig"
)

func LoadDBConfig() (DBConfig, error) {
	d := DBConfig{}
	err := envconfig.Process("", &d)
	if err != nil {
		return d, err
	}
	return d, err
}

func LoadInstrumentationConfig() (InstrumentationConfig, error) {
	i := InstrumentationConfig{}
	err := envconfig.Process("", &i)
	if err != nil {
		return i, err
	}
	return i, err
}

func LoadUserDetailConfig() (UserDetailConfig, error) {
	ud := UserDetailConfig{}
	err := envconfig.Process("", &ud)
	if err != nil {
		return ud, err
	}
	return ud, err
}
