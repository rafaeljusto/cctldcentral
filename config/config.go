package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const prefix = "cctldcentral"

// CCTLDCentral global project configuration.
var CCTLDCentral *cctldstatsConfig

// config contains all configuration parameters for running the statistic services.
type cctldstatsConfig struct {
	// Database stores all the database information to connect to the back-end.
	Database struct {
		// Name name of the database.
		Name string `envconfig:"name"`

		// Username user used to connect to the desired database.
		Username string `envconfig:"username"`

		// Password password used to connect to the desired database.
		Password string `envconfig:"password"`

		// Host address of the database
		Host string `envconfig:"host"`
	} `envconfig:"database"`

	Scheduler struct {
		Timeout time.Duration `envconfig:"timeout"`
		Retries int           `envconfig:"retries"`
	} `envconfig:"scheduler"`
}

// Load fill the global configuration variable using default values and environment variables.
func Load() error {
	CCTLDCentral = new(cctldstatsConfig)
	CCTLDCentral.Database.Name = "cctldcentral"
	CCTLDCentral.Database.Username = "cctldcentral"
	CCTLDCentral.Database.Host = "localhost:5432"
	CCTLDCentral.Scheduler.Timeout = 5 * time.Second
	CCTLDCentral.Scheduler.Retries = 3

	if err := envconfig.Process(prefix, CCTLDCentral); err != nil {
		return err
	}

	return nil
}
