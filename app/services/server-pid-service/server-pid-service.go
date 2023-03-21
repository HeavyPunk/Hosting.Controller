package server_pid_service

import (
	"os"
)

var pidFile = "server.pid"

func Init() *ServerPidServiceContext {
	return &ServerPidServiceContext{}
}

func (serviceContext *ServerPidServiceContext) SavePid(pid string) error {
	err := os.WriteFile(pidFile, []byte(pid), 0777)
	return err
}

func (serviceContext *ServerPidServiceContext) GetPid() (string, error) {
	pid, err := os.ReadFile(pidFile)
	if err != nil {
		return "", err
	}
	return string(pid), nil
}
