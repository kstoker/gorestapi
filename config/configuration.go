package config

import ()

type Configuration struct {
	Router RouterConfiguration
	Database DatabaseConfiguration
	Create CreateConfiguration
	Security SecurityConfiguration
}
