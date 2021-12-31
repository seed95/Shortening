package config

type (
	Config struct {
		Application Application `yaml:"application"`
		Logger      Logger      `yaml:"logger"`
		Translator  Translator  `yaml:"translator"`
		Database    Database    `yaml:"database"`
		Server      Server      `yaml:"server"`
	}

	Application struct {
		Expire         string `yaml:"expire"`
		AliasMinLength int    `yaml:"alias_min_length"`
	}

	Logger struct {
		Logrus Logrus `yaml:"logrus"`
	}

	Logrus struct {
		Path         string `yaml:"internal_path"`
		Pattern      string `yaml:"filename_pattern"`
		RotationSize string `yaml:"max_size"`
		RotationTime string `yaml:"rotation_time"`
		MaxAge       string `yaml:"max_age"`
	}

	Translator struct {
		I18n I18n `yaml:"i18n"`
	}

	I18n struct {
		MessagePath string `yaml:"message_path"`
	}

	Database struct {
		Redis Redis `yaml:"redis"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DbIndex  int    `yaml:"db"`
	}

	Server struct {
		Rest Rest `yaml:"rest"`
		GRPC GRPC `yaml:"grpc"`
	}

	Rest struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
)
