package configschema

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadManifestRoot(path string) (*ManifestRoot, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var manifest ManifestRoot
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest root: %w", err)
	}
	return &manifest, nil
}

func LoadConfiguration(path string) (*ConfigurationFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg ConfigurationFile
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse configuration: %w", err)
	}
	return &cfg, nil
}

func LoadRawConfig(path string) (map[string]interface{}, error) {
	return loadRawMap(path)
}

func LoadRawManifest(path string) (map[string]interface{}, error) {
	return loadRawMap(path)
}

func loadRawMap(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return loadRawYAML(data)
}

func loadRawYAML(data []byte) (map[string]interface{}, error) {
	var raw interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	if raw == nil {
		return map[string]interface{}{}, nil
	}
	normalized, ok := normalizeYAML(raw).(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected YAML root object")
	}

	jsonBytes, err := json.Marshal(normalized)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}
