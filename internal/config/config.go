package config

type ConfigurationManifest struct {
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	URL string `env:"DB_URL"`
}
