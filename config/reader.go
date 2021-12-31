package config

import (
	"espad_task/build/messages"
	"espad_task/pkg/derrors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func Parse(path string, cfg *Config) (err error) {

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYaml(path, cfg)
	default:
		return derrors.New(derrors.Invalid, messages.UnknownFileExtension)
	}

}

func parseYaml(path string, cfg *Config) (err error) {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); e == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
