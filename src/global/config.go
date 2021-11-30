package global

type AppConfig struct {
	Debug bool `yaml:"debug"`
}

func (appConfig *AppConfig) SetDebug(debug bool) {
	appConfig.Debug = debug
}

func (appConfig *AppConfig) GetDebug() bool {
	return appConfig.Debug
}
