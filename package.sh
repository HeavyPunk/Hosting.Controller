set -eu

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=1
# export CC_FOR_TARGET=gcc-aarch64-linux-gnu
# export CC=aarch64-linux-gnu-gcc
BUILD_PATH=./publish
BIN_NAME=controller
ARTEFACT=$BUILD_PATH/$BIN_NAME
rm -rf $BUILD_PATH
go build -o=$ARTEFACT

cp prod-settings.yml $BUILD_PATH/settings.yml
cp prod-startup-config.yml $BUILD_PATH/startup-config.yml
