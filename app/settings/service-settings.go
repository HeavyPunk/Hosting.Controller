package settings

type ServiceSettings struct {
	EnvironmentVars map[string]string `yaml:"environment-variables"`
	App             struct {
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

			ServerFiles struct {
				S3 struct {
					Enabled  bool   `yaml:"enabled"`
					Endpoint string `yaml:"endpoint"`
				} `yaml:"s3"`
				WorkerPull struct {
					WorkerCount int `yaml:"worker-count"`
					QueueSize   int `yaml:"queue-size"`
				} `yaml:"worker-pool"`
			} `yaml:"server-files"`
		} `yaml:"services"`
	} `yaml:"app"`
	Server struct {
		StartupConfig string `yaml:"startup-config"`
	} `yaml:"server"`
}
