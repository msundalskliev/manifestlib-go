package configschema

import "fmt"

func normalizeYAML(v interface{}) interface{} {
	switch value := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(value))
		for key, nested := range value {
			result[key] = normalizeYAML(nested)
		}
		return result
	case map[interface{}]interface{}:
		result := make(map[string]interface{}, len(value))
		for key, nested := range value {
			result[fmt.Sprintf("%v", key)] = normalizeYAML(nested)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(value))
		for i, item := range value {
			result[i] = normalizeYAML(item)
		}
		return result
	default:
		return value
	}
}
