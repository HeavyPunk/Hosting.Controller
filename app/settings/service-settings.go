package settings

type ServiceSettings struct {
	App struct {
		Port          uint   `yaml:"port"`
		Configuration string `yaml:"configuration"`
		Services      struct {
			ServerMessaging struct {
				Rcon struct {
					Enabled            bool   `yaml:"enabled"`
					ServerRconPort     uint   `yaml:"server-rcon-port"`
					ServerRconPassword string `yaml:"server-rcon-password"`
				} `yaml:"rcon"`
			} `yaml:"server-messaging"`

			ServerInfo struct {
				Query struct {
					Enabled         bool `yaml:"enabled"`
					ServerQueryPort uint `yaml:"server-query-port"`
				} `yaml:"query"`
			} `yaml:"server-info"`

			ServerLogging struct {
				Enabled       bool   `yaml:"enabled"`
				LoggingSystem string `yaml:"logging-system"`
				PageSize      int    `yaml:"page-size"`
				LogFile       string `yaml:"log-file"`
			} `yaml:"server-logging"`
		} `yaml:"services"`
	} `yaml:"app"`
	Server struct {
		StartupConfig string `yaml:"startup-config"`
	} `yaml:"server"`
}
