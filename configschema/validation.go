package configschema

import (
	"fmt"

	"github.com/msundalskliev/manifestlib-go/schema"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateInputs(configPath, manifestPath string) error {
	if err := validateYAMLAgainstSchema(configPath, schema.ConfigurationSchema(), "configuration"); err != nil {
		return err
	}
	if err := validateYAMLAgainstSchema(manifestPath, schema.ManifestSchema(), "manifest"); err != nil {
		return err
	}
	return nil
}

func validateYAMLAgainstSchema(docPath string, schemaData []byte, schemaName string) error {
	doc, err := loadRawMap(docPath)
	if err != nil {
		return err
	}
	schemaDoc, err := loadRawYAML(schemaData)
	if err != nil {
		return fmt.Errorf("failed to parse embedded %s schema: %w", schemaName, err)
	}

	result, err := gojsonschema.Validate(gojsonschema.NewGoLoader(schemaDoc), gojsonschema.NewGoLoader(doc))
	if err != nil {
		return fmt.Errorf("failed to validate %s against schema: %w", docPath, err)
	}
	if result.Valid() {
		return nil
	}

	msg := fmt.Sprintf("%s failed %s schema validation", docPath, schemaName)
	for _, desc := range result.Errors() {
		msg += fmt.Sprintf("\n- %s", desc)
	}
	return fmt.Errorf("%s", msg)
}
