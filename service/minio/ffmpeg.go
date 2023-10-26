package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/xfrr/goffmpeg/transcoder"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"github.com/youngdev-stack/minio-ffmpeg-lambda/models/minio/request"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type FfmpegService struct {
	s3URL          string
	minioHost      string
	userIdentity   request.UserIdentity
	userRequestUrl string
	fileName       string
	filePrefix     string
	fileSuffix     string
	bucketName     string
}

func (f FfmpegService) ConvertVideo(event request.EventReq) error {

	f.s3URL = event.GetObjectContext.InputS3Url
	f.userRequestUrl = event.UserRequest.Url
	f.userIdentity = event.UserIdentity
	f.bucketName = strings.Split(event.UserRequest.Url, "/")[0]
	inputURL, err := url.Parse(f.s3URL)
	if err != nil {
		global.GlobalLog.Error("inputURL parse error", zap.Any("err: ", err))
		return err
	}

	f.minioHost = inputURL.Host

	u, err := url.Parse(f.minioHost + f.userRequestUrl)
	if err != nil {
		global.GlobalLog.Error("u parse error", zap.Any("err: ", err))
		return err
	}

	if f.fileName = u.Query().Get("lccFileName"); f.fileName == "" {
		return fmt.Errorf("lccFileName is empty")
	}

	f.filePrefix, f.fileSuffix = splitFileName(f.fileName)
	if err := f.downloadAndConvert(); err != nil {
		global.GlobalLog.Error("convert error", zap.Any("err: ", err))
		return err
	}

	if err := f.uploadMinio(); err != nil {
		global.GlobalLog.Error("upload error", zap.Any("err: ", err))
		return err
	}
	return nil
}

func (f FfmpegService) downloadAndConvert() error {
	// 下载文件
	r, err := http.Get(f.s3URL)
	if err != nil {
		global.GlobalLog.Error("download file error", zap.Any("err: ", err))
		return err
	}

	// 保存文件
	file, err := os.Create(f.fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer os.Remove(file.Name())
	defer file.Close()

	// 创建存储目录
	err = os.Mkdir(f.filePrefix, 0777)
	if err != nil {
		return err
	}
	defer os.RemoveAll(f.filePrefix)

	_, err = io.Copy(file, r.Body)
	if err != nil {
		global.GlobalLog.Error("copy file error", zap.Any("err: ", err))
		return err
	}

	if err := f.convert(); err != nil {
		global.GlobalLog.Error("convert error", zap.Any("err: ", err))
		return err
	}

	return nil

}

func (f FfmpegService) uploadMinio() error {
	s3Client, err := minio.New(f.minioHost, &minio.Options{
		Creds:  credentials.NewStaticV4(f.userIdentity.PrincipalId, f.userIdentity.AccessKeyId, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	filePath := "/home/youngdev-stack/minio-ffmpeg-lambda/" + f.filePrefix

	// Upload the files
	return filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		objectName := f.filePrefix + strings.TrimPrefix(path, filePath)
		contentType := "application/octet-stream"
		_, err = s3Client.FPutObject(context.Background(), f.bucketName, objectName, path, minio.PutObjectOptions{
			ContentType: contentType,
		})
		if err != nil {
			return err
		}
		fmt.Printf("Successfully uploaded %s\n", objectName)
		return nil
	})

}

func (f FfmpegService) convert() error {
	trans := new(transcoder.Transcoder)

	// 文件存储路径为 filePrefix/filePrefix+%d.ts
	err := trans.Initialize(f.fileName, f.filePrefix+"/"+f.filePrefix+".m3u8")
	if err != nil {
		return err
	}

	// 开始转码
	// TODO: 配置化
	trans.MediaFile().SetVideoCodec("libx264")
	trans.MediaFile().SetHlsSegmentDuration(4)
	trans.MediaFile().SetVideoBitRate("720k")
	trans.MediaFile().SetFrameRate(15)
	trans.MediaFile().SetAudioCodec("aac")
	trans.MediaFile().SetAudioBitRate("64k")
	trans.MediaFile().SetHlsSegmentDuration(2)
	trans.MediaFile().SetOutputFormat("hls")
	done := trans.Run(true)
	progress := trans.Output()
	for p := range progress {
		global.GlobalLog.Info("convert progress", zap.Any("progress: ", p))
	}

	for err := range done {
		if err != nil {
			global.GlobalLog.Error("convert error", zap.Any("err: ", err))
			return err
		}
	}
	return nil
}

// 按.切割文件前缀名和后缀名
func splitFileName(fileName string) (string, string) {
	var (
		fileNamePrefix string
		fileNameSuffix string
	)
	if fileName == "" {
		return fileNamePrefix, fileNameSuffix
	}
	fileNamePrefix = fileName[:strings.LastIndex(fileName, ".")]
	fileNameSuffix = fileName[strings.LastIndex(fileName, "."):]
	return fileNamePrefix, fileNameSuffix
}
