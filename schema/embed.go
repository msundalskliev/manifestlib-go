package schema

import _ "embed"

var (
	//go:embed configuration.schema.yaml
	configurationSchema []byte

	//go:embed manifest.schema.yaml
	manifestSchema []byte
)

func ConfigurationSchema() []byte {
	return configurationSchema
}

func ManifestSchema() []byte {
	return manifestSchema
}
