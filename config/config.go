package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Timer      Timer      `mapstructure:"timer" json:"timer" yaml:"timer"`
	Kubernetes Kubernetes `mapstructure:"kubernetes" json:"kubernetes" yaml:"kubernetes"`
	UseRedis   bool       `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
}
