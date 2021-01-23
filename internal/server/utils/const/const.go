package serverConst

type VmPlatform string

const (
	Kvm VmPlatform = "kvm"
	Pve VmPlatform = "pve"
)

type ContainerPlatform string

const (
	Docker    ContainerPlatform = "docker"
	Portainer ContainerPlatform = "portainer"
)
