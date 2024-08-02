package yaml_reader

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Country struct {
	Name   string `yaml:"name"`
	Prefix string `yaml:"prefix"`
	Length int    `yaml:"length"`
}

type IBANRules struct {
	MinLength int `yaml:"min_length"`
	MaxLength int `yaml:"max_length"`
}

type Settings struct {
	IBAN_rules IBANRules `yaml:"iban_rules"`
	Countries  []Country `yaml:"countries"`
}

func ReadYAML(filename string) (Settings, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Settings{}, err
	}

	var settings Settings

	err = yaml.Unmarshal(data, &settings)

	if err != nil {
		return Settings{}, err
	}

	return settings, nil
}
