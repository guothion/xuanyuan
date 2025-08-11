package config

const (
	KeyHeaderServiceProfile = "X-GL-SERVICE-PROFILE"
	KeyHeaderAccessToken    = "X-GL-ACCESS"
	KeyHeaderAccount        = "X-GL-ACCOUNT"
	KeyRequestContext       = "X-GL-REQUEST-CONTEXT"
)

type Configuration struct {
	DataSource *dataSource `yaml:"dataSource" json:"dataSource"`
	Log        *logConf    `yaml:"log" json:"log"`
	App        App         `yaml:"app" json:"app"`
	Redis      *Redis      `yaml:"redis" json:"redis"`
}
