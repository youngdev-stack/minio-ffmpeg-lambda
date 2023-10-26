#!/bin/sh
set -e

printf "Starting Minio Lambda..."

exec ./minio-ffmpeg-lambda -c conf/config.yaml