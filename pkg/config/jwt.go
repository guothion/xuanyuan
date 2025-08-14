package config

type Jwt struct {
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" json:"blacklist_grace_period" yaml:"blacklist_grace_period"`
	RefreshGracePeriod      int64  `mapstructure:"refresh_grace_period" json:"refresh_grace_period" yaml:"refresh_grace_period"`
}
