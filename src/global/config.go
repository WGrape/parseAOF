package global

type AppConfig struct {
	Debug       bool `yaml:"debug"`
	MaxRoutines int  `yaml:"maxRoutines"`
}

func (appConfig *AppConfig) SetMaxRoutines(maxRoutines int) {
	appConfig.MaxRoutines = maxRoutines
}

func (appConfig *AppConfig) SetDebug(debug bool) {
	appConfig.Debug = debug
}

func (appConfig *AppConfig) GeMaxRoutines() int {
	return appConfig.MaxRoutines
}

func (appConfig *AppConfig) GetDebug() bool {
	return appConfig.Debug
}
