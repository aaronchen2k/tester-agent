package _domain

type RpcReq struct {
	NodeIp   string
	NodePort int

	ApiPath   string
	ApiMethod string
	Data      interface{}
}
