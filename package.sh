OS=linux
ARCH=amd64
BUILD_PATH=./publish
BIN_NAME=controller
ARTEFACT=$BUILD_PATH/$BIN_NAME
rm -rf $BUILD_PATH
env GOOS=$OS GOARCH=$ARCH go build -o=$ARTEFACT

cp prod-settings.yml $BUILD_PATH/settings.yml
cp prod-startup-config.yml $BUILD_PATH/startup-config.yml
