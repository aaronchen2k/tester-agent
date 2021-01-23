package service

import (
	"bytes"
	"fmt"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/aaronchen2k/tester/internal/server/model"
	gateway "github.com/rpcx-ecosystem/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
	"io/ioutil"
	"net/http"
)

type RpcService struct{}

func NewRpcService() *RpcService {
	return &RpcService{}
}

func (s *RpcService) AppiumTest(build model.Build) (result _domain.RpcResult) {
	appiumTestTo := model.NewTestTo(build)
	appiumTestTo.AppiumPort = build.AppiumPort

	obj := interface{}(appiumTestTo)
	s.Request(build.NodeIp, build.NodePort, "appium", "AppiumTest", &obj)

	result.Success(fmt.Sprintf("success to send rpc build request %#v.", build))
	return
}

func (s *RpcService) SeleniumTest(build model.Build) (result _domain.RpcResult) {
	seleniumTestTo := model.NewTestTo(build)
	seleniumTestTo.SeleniumDriverType = build.Queue.BrowserType
	seleniumTestTo.SeleniumDriverVersion = build.Queue.BrowserVer

	obj := interface{}(seleniumTestTo)
	s.Request(build.NodeIp, build.NodePort, "selenium", "SeleniumTest", &obj)

	result.Success(fmt.Sprintf("success to send rpc build request %#v.", build))
	return
}

func (s *RpcService) CreateVm(req _domain.PveReq) (result _domain.RpcResult) {
	obj := interface{}(req)
	result = s.Request(req.NodeIp, req.NodePort, "vm", "Create", &obj)

	result.Success(fmt.Sprintf("success to create vm via rpc %#v.", req))
	return
}

func (s *RpcService) RestartVm(req _domain.PveReq) (result _domain.RpcResult) {
	obj := interface{}(req)
	result = s.Request(req.NodeIp, req.NodePort, "vm", "Restart", &obj)

	result.Success(fmt.Sprintf("success to start vm via rpc %#v.", req))
	return
}

func (s *RpcService) StartVm(req _domain.PveReq) (result _domain.RpcResult) {
	obj := interface{}(req)
	s.Request(req.NodeIp, req.NodePort, "vm", "Start", &obj)

	result.Success(fmt.Sprintf("success to start vm via rpc %#v.", req))
	return
}

func (s *RpcService) StopVm(req _domain.PveReq) (result _domain.RpcResult) {
	obj := interface{}(req)
	s.Request(req.NodeIp, req.NodePort, "vm", "Stop", &obj)

	result.Success(fmt.Sprintf("success to stop vm via rpc %#v.", req))
	return
}

func (s *RpcService) ConvertVmToTempl(req _domain.PveReq) (result _domain.RpcResult) {
	obj := interface{}(req)
	s.Request(req.NodeIp, req.NodePort, "vm", "ConvertToBackingImage", &obj)

	result.Success(fmt.Sprintf("success to convert vm to backingImage via rpc %#v.", req))
	return
}

func (s *RpcService) Request(ip string, port int, apiPath string, method string, param *interface{}) (rpcResult _domain.RpcResult) {
	cc := &codec.MsgpackCodec{}

	data, _ := cc.Encode(param)
	url := fmt.Sprintf("http://%s:%d/", ip, port)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		_logUtils.Errorf("Fail to create request: %#v", err)
		return
	}

	// 设置header
	h := req.Header
	h.Set(gateway.XServicePath, apiPath)
	h.Set(gateway.XServiceMethod, method)
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		_logUtils.Errorf("fail to call: %#v.", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_logUtils.Errorf("fail to read response: %#v.", err)
	}

	err = cc.Decode(body, &rpcResult)
	if err != nil {
		_logUtils.Errorf("fail to decode reply: %s.", err.Error())
	}

	msg := fmt.Sprintf("agent return %d-%s.", rpcResult.Code, rpcResult.Msg)
	_logUtils.Info(msg)
	return
}
