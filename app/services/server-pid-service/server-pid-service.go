package server_pid_service

import (
	"io/fs"
	"io/ioutil"
)

var pidFile = "server.pid"

func Init() *ServerPidServiceContext {
	return &ServerPidServiceContext{}
}

func (serviceContext *ServerPidServiceContext) SavePid(pid string) {
	err := ioutil.WriteFile(pidFile, []byte(pid), fs.ModeAppend)
	if err != nil {
		panic(err)
	}
}

func (serviceContext *ServerPidServiceContext) GetPid() string {
	pid, err := ioutil.ReadFile(pidFile)
	if err != nil {
		panic(err)
	}
	return string(pid)
}
