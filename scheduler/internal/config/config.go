package config
package config

import "os"

type Config struct {
	Interval string
}

func Load() Config {
	interval := os.Getenv("SCHEDULER_INTERVAL")
	if interval == "" {
		interval = "30s"
	}

	return Config{Interval: interval}
}
