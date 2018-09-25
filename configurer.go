package cnfg

import "flag"

type Configurer interface {
	Load(file string) (err error)                          // load configuration from a file
	SetEnvironmentPrefix(prefix string)                    // env var prefix, like KINDRID_
	CheckEnvOrFile(key string) (result string, found bool) //return a value contained in ENV, then if blank, in a file, else ""
	SetBool(destination *bool, defaultValue bool, flags *flag.FlagSet, cliName, description string)
	SetString(destination *string, defaultValue string, flags *flag.FlagSet, cliName, description string)
	SetInt(destination *int, defaultValue int, flags *flag.FlagSet, cliName, description string)
	SetFloat(destination *float64, defaultValue float64, flags *flag.FlagSet, cliName, description string)
	SetCSV(destination *[]string, defaultValue string, flags *flag.FlagSet, cliName, description string)
}
