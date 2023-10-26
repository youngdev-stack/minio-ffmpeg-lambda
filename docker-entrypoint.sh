#!/bin/sh
set -e

printf "Starting Minio Lambda..."

exec ./minio-ffmpeg-lambda -c config/config.yaml