package config

type dataSource struct {
	Driver              string `yaml:"driver" json:"driver"`
	Env                 string `yaml:"env" json:"env"`
	URL                 string `yaml:"url" json:"url"`
	Host                string `yaml:"host" json:"host"`
	Port                int    `yaml:"port" json:"port"`
	Username            string `yaml:"username" json:"username"`
	Password            string `yaml:"password" json:"password"`
	Schema              string `yaml:"schema" json:"schema"`
	Verbose             bool   `yaml:"verbose" json:"verbose"`
	Pem                 string `yaml:"pem" json:"pem"`
	Param               string `yaml:"param" json:"param"`
	Charset             string `yaml:"charset" json:"charset"`
	MaxIdleConns        int    `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns        int    `yaml:"max_open_conns" json:"max_open_conns"`
	LogMode             string `yaml:"log_mode" json:"log_mode"`
	EnableFileLogWriter bool   `yaml:"enable_file_log_writer" json:"enable_file_log_writer"`
	LogFileName         string `yaml:"log_file_name" json:"log_file_name"`
}
