package settings

type ServiceSettings struct {
	App struct {
		Port uint `yaml:"port"`
	} `yaml:"app"`
}
