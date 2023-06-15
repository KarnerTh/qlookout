package orchestration

import (
	"flag"
)

type CliFlags struct {
	ConfigPath string
}

func parseFlags() CliFlags {
	configPath := flag.String("config", "", "Path to the config file")
	flag.Parse()

	return CliFlags{
		ConfigPath: *configPath,
	}
}
