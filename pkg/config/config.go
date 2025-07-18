package config

type config struct {
	DataSource *dataSource `yaml:"dataSource" json:"dataSource"`
	Log        *logConf    `yaml:"log" json:"log"`
}

type dataSource struct {
	Type     string `yaml:"type" json:"type"`
	Env      string `yaml:"env" json:"env"`
	URL      string `yaml:"url" json:"url"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Schema   string `yaml:"schema" json:"schema"`
	Verbose  bool   `yaml:"verbose" json:"verbose"`
	Pem      string `yaml:"pem" json:"pem"`
	Param    string `yaml:"param" json:"param"`
}

type logConf struct {
	Level   string `yaml:"level" json:"level"`
	Path    string `yaml:"path" json:"path"`
	Size    int    `yaml:"size" json:"size"`
	Backups int    `yaml:"backups" json:"backups"`
	Age     int    `yaml:"age" json:"age"`
}
