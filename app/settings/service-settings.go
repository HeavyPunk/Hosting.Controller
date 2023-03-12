package settings

type ServiceSettings struct {
	App struct {
		Port          uint   `yaml:"port"`
		Configuration string `yaml:"configuration"`
	} `yaml:"app"`
	Server struct {
		StartupConfig string `yaml:"startup-config"`
	} `yaml:"server"`
}
