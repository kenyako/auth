package config

import "flag"

var ConfigPath string

func InitFlags() {
	flag.StringVar(&ConfigPath, "config-path", ".env", "path to config file")
}
