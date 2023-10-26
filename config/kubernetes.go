package config

type Kubernetes struct {
	KubeConfig string `mapstructure:"kubeconfig" json:"kubeconfig" yaml:"kubeconfig"` // kubeconfig路径
	InCluster  bool   `mapstructure:"in-cluster" json:"in-cluster" yaml:"in-cluster"` // 是否使用集群内部配置
}
