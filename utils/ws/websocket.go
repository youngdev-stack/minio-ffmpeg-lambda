package ws

import (
	"github.com/youngdev-stack/minio-ffmpeg-lambda/global"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"net"
)

func WriterCopy(src net.Conn, dst *websocket.Conn) {
	defer src.Close()
	defer dst.Close()

	// 从源 net.Conn 中读取数据，并写入到目标 WebSocket 连接中
	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err != io.EOF {
				global.BaileysLog.Info("The client disconnected actively, reason: ", zap.Error(err))
			}
			break
		}
		w, err := dst.NextWriter(websocket.BinaryMessage)
		if err != nil {
			global.BaileysLog.Error("Error getting next writer:", zap.Error(err))
			break
		}
		_, err = w.Write(buf[:n])
		if err != nil {
			global.BaileysLog.Error("Error writing to destination connection: ", zap.Error(err))
			break
		}
		err = w.Close()
		if err != nil {
			global.BaileysLog.Error("Error closing writer: ", zap.Error(err))
			break
		}
	}
}

func ReaderCopy(src *websocket.Conn, dest io.Writer, done chan struct{}) {
	for {
		_, message, err := src.ReadMessage()
		if err != nil {
			global.BaileysLog.Info("The client disconnected actively, reason: ", zap.Error(err))
			close(done)
			return
		}
		if string(message) == "exit" {
			close(done)
			return
		}
		_, err = dest.Write(message)
		if err != nil {
			global.BaileysLog.Error("Error writing message to container: ", zap.Error(err))
			close(done)
			return
		}
	}
}
