package server_manager

type ServerControllerContext struct {
	ServerPid string
}

type RunServerRequest struct {
	RunCmd          string
	Args            []string
	WorkingDir      string
	EnvironmentVars []string
	SaveStdout      bool //TODO: implement
	SaveStderr      bool //TODO: implement
}

type RunServerResponse struct {
	Success bool
	Error   error
}

type StopServerRequest struct {
	ForceInterrupt bool
}

type StopServerResponse struct {
	Success bool
	Error   error
}

type GetFileFromServerRequest struct {
	RelativeFilepath string
}

type GetFileFromServerResponse struct {
	Success bool
	Error   error
	Content []byte
}

type PostFileToServerRequest struct {
	RelativeFilepath string
	Content          []byte
}

type PostFileToServerResponse struct {
	Success bool
}
