package configschema

import "sort"

type ManifestRoot struct {
	Manifest struct {
		Metadata struct {
			Repo struct {
				URL    string `yaml:"url"`
				Branch string `yaml:"branch"`
				Tag    string `yaml:"tag"`
			} `yaml:"repo"`
		} `yaml:"metadata"`
		Includes map[string]string `yaml:"includes"`
	} `yaml:"manifest"`
}

func (m *ManifestRoot) IncludePath(tool string) string {
	if m == nil {
		return ""
	}
	return m.Manifest.Includes[tool]
}

type ToolBlock struct {
	Backend      map[string]string `yaml:"backend"`
	Includes     map[string]string `yaml:"includes"`
	IncludeOrder []string          `yaml:"include_order"`
}

type Structure struct {
	Root  string               `yaml:"root"`
	Tools map[string]ToolBlock `yaml:",inline"`
}

type ConfigurationFile struct {
	Configuration struct {
		Metadata struct {
			Structure Structure `yaml:"structure"`
		} `yaml:"metadata"`
	} `yaml:"configuration"`
}

func (c *ConfigurationFile) ToolIncludePaths(tool string) []string {
	toolBlock, ok := c.Configuration.Metadata.Structure.Tools[tool]
	if !ok {
		return nil
	}

	includes := toolBlock.Includes
	if len(includes) == 0 {
		return nil
	}

	orderLookup := map[string]int{}
	for idx, key := range toolBlock.IncludeOrder {
		orderLookup[key] = idx
	}

	keys := make([]string, 0, len(includes))
	for key := range includes {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		iOrder, iExists := orderLookup[keys[i]]
		jOrder, jExists := orderLookup[keys[j]]
		if iExists && jExists {
			return iOrder < jOrder
		}
		if iExists {
			return true
		}
		if jExists {
			return false
		}
		return keys[i] < keys[j]
	})

	paths := make([]string, 0, len(keys))
	for _, key := range keys {
		if includes[key] == "" {
			continue
		}
		paths = append(paths, includes[key])
	}
	return paths
}
