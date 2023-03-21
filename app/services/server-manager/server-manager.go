package server_manager

import (
	"os/exec"
	server_pid_service "simple-hosting/controller/app/services/server-pid-service"
	"strconv"
)

func Init() *ServerControllerContext {
	return &ServerControllerContext{}
}

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
	signal := "-15"
	if request.ForceInterrupt {
		signal = "-2"
	}
	pidService := server_pid_service.Init()
	pidStr, err := pidService.GetPid()
	if err != nil {
		return StopServerResponse{
			Success: false,
			Error:   err,
		}
	}
	cmd := exec.Command("kill", signal, pidStr)
	if err = cmd.Run(); err != nil {
		pidStr, _ = pidService.GetPid()
		cmd = exec.Command("kill", signal, pidStr)
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
