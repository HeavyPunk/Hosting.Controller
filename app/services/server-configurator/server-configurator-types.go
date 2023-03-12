package server_configurator

type ServerConfiguration struct {
	WorkingDir       string   `yaml:"working-dir"`
	Params           []string `yaml:"params"`
	EnvVars          []string `yaml:"env-vars"`
	PrelaunchScripts []string `yaml:"prelaunch-scripts"`
	RunCmd           struct {
		Cmd  string   `yaml:"cmd"`
		Args []string `yaml:"args"`
	} `yaml:"run-cmd"`
}
