package vm_service

import (
	"github.com/docker/docker/client"
)

func Init() *VmServiceContext {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return &VmServiceContext{
		client: cli,
	}
}

func (hypContext *VmServiceContext) CreateVm(request VmCreateRequest) VmCreateResponse {
	cli := hypContext.client
	if cli == nil {
		panic("Hypervisor client is nil")
	}

}

func (hypContext *VmServiceContext) StopVm(request VmStopRequest) VmStopResponse {

}

func (hypContext *VmServiceContext) SuspendVm(request VmSuspendRequest) VmSuspendResponse {

}
