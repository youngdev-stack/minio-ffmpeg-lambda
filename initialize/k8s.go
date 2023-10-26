package initialize

import (
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8sConfig() (conf *rest.Config) {
	var err error
	if global.GlobalConfig.Kubernetes.InCluster {
		conf, err = rest.InClusterConfig()
	} else if global.GlobalConfig.Kubernetes.KubeConfig != "" {
		conf, err = clientcmd.BuildConfigFromFlags("", global.GlobalConfig.Kubernetes.KubeConfig)

	} else {
		global.GlobalLog.Fatal("kubeconfig path is empty")
		return nil
	}
	if err != nil {
		global.GlobalLog.Fatal("failed to create Kubernetes config", zap.Error(err))
	}
	global.GlobalLog.Info("create Kubernetes config success")
	return conf
}

func NewDynamicClient(config *rest.Config) *kubernetes.Clientset {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.GlobalLog.Fatal("failed to create Kubernetes clientset: ", zap.Error(err))
		return nil
	}
	global.GlobalLog.Info("create Kubernetes clientset success")
	return clientset
}
