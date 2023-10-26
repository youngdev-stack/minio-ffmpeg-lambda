package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境值
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 端口值
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由前缀

	//RunMode       string   `mapstructure:"run-mode" json:"run-mode" yaml:"run-mode"`                // 运行模式
}
