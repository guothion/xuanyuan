package config

type logConf struct {
	Level    string `yaml:"level" json:"level"`
	RootDir  string `yaml:"rootdir" json:"rootdir"`
	Filename string `yaml:"filename" json:"filename"`
	Size     int    `yaml:"size" json:"size"`
	Backups  int    `yaml:"backups" json:"backups"`
	Age      int    `yaml:"age" json:"age"`
	Compress bool   `yaml:"compress" json:"compress"`
}
