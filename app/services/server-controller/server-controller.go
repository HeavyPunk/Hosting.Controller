package server_controller

import (
	"os/exec"
	server_pid_service "simple-hosting/controller/app/services/server-pid-service"
	"strconv"
)

func (serviceContext *ServerControllerContext) RunServer(request RunServerRequest) RunServerResponse {
	cmd := exec.Command(request.RunCmd, request.Args...)
	cmd.Dir = request.WorkingDir
	cmd.Env = request.EnvironmentVars
	if err := cmd.Start(); err != nil {
		return RunServerResponse{
			Success: false,
			Error:   err,
		}
	}
	pidService := server_pid_service.Init()
	pidStr := strconv.Itoa(cmd.Process.Pid)
	pidService.SavePid(pidStr)
	serviceContext.ServerPid = pidStr
	return RunServerResponse{
		Success: true,
	}
}

func (serviceContext *ServerControllerContext) StopServer(request StopServerRequest) StopServerResponse {
	pidStr := serviceContext.ServerPid
	pidService := server_pid_service.Init()
	cmd := exec.Command("kill", pidStr)
	if err := cmd.Run(); err != nil {
		pidStr = pidService.GetPid()
		cmd = exec.Command("kill", pidStr)
		if err = cmd.Run(); err != nil {
			return StopServerResponse{Success: false, Error: err}
		}
	}
	return StopServerResponse{Success: true}
}

func (serviceContext *ServerControllerContext) GetFileFromServer(request GetFileFromServerRequest) GetFileFromServerResponse {
	panic("Not implemented")
}

func (serviceContext *ServerControllerContext) PostFileToServer(request PostFileToServerRequest) PostFileToServerResponse {
	panic("Not implemented")
}
