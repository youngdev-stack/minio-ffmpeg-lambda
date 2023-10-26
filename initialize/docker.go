package initialize

import (
	"context"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"github.com/docker/docker/client"
	"go.uber.org/zap"
)

func DockerClient() *client.Client {
	ctx := context.Background()
	if cli, err := client.NewClientWithOpts(client.FromEnv); err != nil {
		global.GlobalLog.Error("Initialize the docker client error, %s", zap.Any(" error:", err))
		return nil
	} else {
		global.GlobalLog.Info("Initialize the docker client success")
		cli.NegotiateAPIVersion(ctx)
		return cli
	}

}
