package cnfg

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	EnvironmentPrefix string
	File              string
	FileValues        map[string]string
}

func NewConfig(prefix, file string) (result *Config, err error) {
	result = &Config{EnvironmentPrefix: prefix, File: file, FileValues: make(map[string]string)}
	if file != "" {
		err = result.Load(file)
		if err != nil {
			return nil, err
		}
	}
	return
}

func MustConfig(prefix, file string) (result *Config) {
	result, err := NewConfig(prefix, file)
	if err != nil {
		panicf("error reading config file %s", file)
	}
	return
}

func (cfg *Config) Load(file string) (err error) {
	return errors.New("loading config files is not supported yet")
}

func (cfg *Config) SetEnvironmentPrefix(prefix string) {
	cfg.EnvironmentPrefix = prefix
}

func (cfg *Config) CheckEnvOrFile(key string) (result string, ok bool) {
	result, ok = os.LookupEnv(key)
	if ok {
		return
	}
	result, ok = cfg.FileValues[key]
	if ok {
		return
	}
	return
}

func (cfg *Config) setupDefault(defaultValue, cliName, description string) (result, fullDescription string, ok bool) {
	envname := cfg.EnvironmentPrefix + cliName
	result, ok = cfg.CheckEnvOrFile(envname)
	if ok {
		fullDescription = fmt.Sprintf("%s (overridden by %s in file or environment)", description, envname)
	} else {
		result = defaultValue
		fullDescription = fmt.Sprintf("%s (%s in file or environment)", description, envname)

	}
	return
}

func (cfg *Config) SetString(destination *string, defaultValue string, flags *flag.FlagSet, cliName, description string) {
	newDefault, newDescription, _ := cfg.setupDefault(defaultValue, cliName, description)
	if flags == nil {
		destination = &newDefault
	} else {
		flags.StringVar(destination, cliName, newDefault, newDescription)
	}
}

func (cfg *Config) SetInt(destination *int, defaultValue int, flags *flag.FlagSet, cliName, description string) {
	var err error
	value := defaultValue
	newDefault, newDescription, found := cfg.setupDefault(strconv.FormatInt(int64(defaultValue), 10), cliName, description)
	if found {
		if value, err = strconv.Atoi(newDefault); err != nil {
			panicf("failure trying to parse config value from environment or config file: '%s'", newDefault)
		}
	}
	if flags == nil {
		destination = &value
	} else {
		flags.IntVar(destination, cliName, value, newDescription)
	}
}

func (cfg *Config) SetBool(destination *bool, defaultValue bool, flags *flag.FlagSet, cliName, description string) {
	var err error
	value := defaultValue
	newDefault, newDescription, found := cfg.setupDefault(strconv.FormatBool(defaultValue), cliName, description)
	if found {
		if value, err = strconv.ParseBool(newDefault); err != nil {
			panicf("failure trying to parse config value from environment or config file: '%s'", newDefault)
		}
	}
	if flags == nil {
		destination = &value
	} else {
		flags.BoolVar(destination, cliName, value, newDescription)
	}
}

func panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}

func (cfg *Config) SetFloat(destination *float64, defaultValue float64, flags *flag.FlagSet, cliName, description string) {
	panic("setfloat isn't implemented yet")
}

func (cfg *Config) SetCSV(destination *[]string, defaultValue string, flags *flag.FlagSet, cliName, description string) {
	panic("setcsv isn't implemented yet")
}
