package model

type Config struct {
	Settings []struct {
		Name           string `yaml:"name"`
		Network        string `yaml:"network"`
		ListenAddress  string `yaml:"listenAddress"`
		ListenPort     int    `yaml:"listenPort"`
		Ssl            bool   `yaml:"ssl"`
		BackendAddress []struct {
			Address          string `yaml:"address"`
			Port             int    `yaml:"port"`
			NodeExporterPort int    `yaml:"nodeExporterPort"`
			HealthCheck      struct {
				Path     string `yaml:"path"`
				Interval int    `yaml:"interval"`
				Timeout  int    `yaml:"timeout"`
			} `yaml:"healthCheck"`
		} `yaml:"backendAddress"`
	} `yaml:"settings"`
}
