package pid_service_test

import (
	server_pid_service "simple-hosting/controller/app/services/server-pid-service"
	"testing"
)

var pid = "12345"

func TestSavePid(t *testing.T) {
	pidService := server_pid_service.Init()
	if err := pidService.SavePid(pid); err != nil {
		t.Error(err)
	}
}

func TestGetPid(t *testing.T) {
	pidService := server_pid_service.Init()
	actualPid, err := pidService.GetPid()
	if err != nil {
		t.Error(err)
	}

	if pid != actualPid {
		t.Errorf("expected pid %s is not equal to %s", pid, actualPid)
	}
}
