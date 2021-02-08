package serverConst

type VmPlatform string

const (
	AppType = "server"

	PageSize            = 15
	Kvm      VmPlatform = "kvm"
	Pve      VmPlatform = "pve"

	Docker    ContainerPlatform = "docker"
	Portainer ContainerPlatform = "portainer"
)

type ContainerPlatform string
