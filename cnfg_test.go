package cnfg

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommonConfig struct {
	CommonBool   bool
	CommonString string
	CommonInt    int
	CommonInt64  int64
}

type TestConfig struct {
	CommonConfig
	Bool   bool
	String string
	Int    int
	Int64  int64
}

func TestBasicConfig(tMain *testing.T) {

	makeConfig := func(name string) (cfgr Configurer, cfg *TestConfig, flags *flag.FlagSet, err error) {
		cfgr, err = NewConfig(name, "")
		flags = flag.NewFlagSet(name, flag.PanicOnError)
		cfg = &TestConfig{}
		return
	}

	setupOne := func(source Configurer, dest *TestConfig, flags *flag.FlagSet) {
		source.SetString(&(dest.String), "default", flags, "the-string", "test configuration string")
		source.SetInt(&(dest.Int), 5, flags, "the-int", "test configuration int")
		source.SetBool(&(dest.Bool), true, flags, "the-bool", "test configuration int")
		source.SetString(&(dest.CommonString), "default", flags, "common-string", "test configuration string")
		source.SetInt(&(dest.CommonInt), 5, flags, "common-int", "test configuration int")
		source.SetBool(&(dest.CommonBool), true, flags, "common-bool", "test configuration int")
	}

	tMain.Run("BasicConfig is a Configurer", func(t *testing.T) {
		config, _, _, err := makeConfig("init-test")
		assert.NotNil(t, config)
		assert.Nil(t, err)
	})

	// NEXT make a table to run this with env vars

	tMain.Run("config without file or ENV works handles structure pointers", func(t *testing.T) {
		source, dest, flags, err := makeConfig("just-defaults-test")
		assert.Nil(t, err)
		setupOne(source, dest, flags)
		args := make([]string, 0)
		flags.Parse(args)
		assert.Equal(t, "default", dest.String)
		assert.Equal(t, 5, dest.Int)
		assert.Equal(t, true, dest.Bool)
		assert.Equal(t, "default", dest.CommonString)
		assert.Equal(t, 5, dest.CommonInt)
		assert.Equal(t, true, dest.CommonBool)
	})
}
