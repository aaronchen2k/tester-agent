package domain

type VmReq struct {
	TemplId uint `json:"id"`

	Ip       string `json:"ip,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"-"`
	Password string `json:"-"`

	VmId string `json:"vmId"`
	Node string `json:"node"`
}
