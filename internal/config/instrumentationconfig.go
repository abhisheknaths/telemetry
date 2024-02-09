package config

type InstrumentationConfig struct {
	ExporterEndpoint string `envconfig:"EXPORTER_ENDPOINT"`
	ExporterPath     string `envconfig:"EXPORTER_PATH"`
}
