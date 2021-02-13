package go_portainer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func NewPortainer(c *Config) Portainer {
	return Portainer{
		Config: c,
		Token:  "",
		ApiURL: fmt.Sprintf("%s://%s:%d%s", c.Schema, c.Host, c.Port, c.URL),
	}
}

func (p *Portainer) Auth() error {
	authData := make(map[string]string)
	authData["Username"] = p.Config.User
	authData["Password"] = p.Config.Password
	payload, err := json.Marshal(&authData)

	//res, err := http.Post(, "application/json", bytes.NewReader(payload))
	req, err := http.NewRequest("post", p.ApiURL+"/auth", bytes.NewReader(payload))
	if err != nil {
		log.Printf("http.NewRequest() error: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{Timeout: time.Second * 3}
	res, err := c.Do(req)
	if err != nil {
		_logUtils.Error(err.Error())
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("unauthorized")
	}

	jwtString, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		_logUtils.Error(err.Error())
		return err
	}
	jwtData := make(map[string]string)
	_ = json.Unmarshal(jwtString, &jwtData)
	p.Token = jwtData["jwt"]
	return err
}

func (p *Portainer) ListEndpoints() ([]Endpoint, error) {
	url := "/endpoints"
	res, err := p.makeRequest("GET", url, nil, nil)
	if err != nil {
		log.Printf("http.Do() error: %v\n", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, err
	}
	var endpoints []Endpoint
	err = json.Unmarshal(data, &endpoints)
	if err != nil {
		log.Printf("Endpoints unmarshaling error: %v\n", err)
		return nil, err
	}
	return endpoints, err
}

func (p *Portainer) ListImages(e uint) (images []Image, err error) {
	url := fmt.Sprintf("/endpoints/%d/docker/images/json", e)
	urlargs := make(map[string]string)
	urlargs["all"] = "1"
	res, err := p.makeRequest("GET", url, nil, urlargs)
	if err != nil {
		log.Printf("http.Do() error: %v\n", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, err
	}

	err = json.Unmarshal(data, &images)
	return
}

func (p *Portainer) ListContainers(e uint) ([]Container, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/json", e)
	urlargs := make(map[string]string)
	urlargs["all"] = "1"
	res, err := p.makeRequest("GET", url, nil, urlargs)
	if err != nil {
		log.Printf("http.Do() error: %v\n", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, err
	}
	var containers []Container
	err = json.Unmarshal(data, &containers)
	return containers, nil
}

//noinspection GoNilness
func (p *Portainer) CreateContainer(e uint, body map[string]interface{}) (id string, err error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/create", e)

	requestByte, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(requestByte)

	res, err := p.makeRequest("POST", url, bodyReader, nil)
	defer res.Body.Close()
	if err != nil {
		log.Printf("http.Do(%v) error: %v\n", res.Request.URL, err)
		return
	}

	bodyStr, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	json.Unmarshal(bodyStr, &result)

	switch res.StatusCode {
	case http.StatusNoContent:
		return
	case http.StatusInternalServerError:
		err = errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		err = errors.New(fmt.Sprintf("Not found: (%s)", url))
	case http.StatusNotModified:
		err = errors.New(fmt.Sprintf("Already started: (%s)", url))
	default:
		err = errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}

	return
}

//noinspection GoNilness
func (p *Portainer) StartContainer(e uint, id string) (int, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/start", e, id)
	res, err := p.makeRequest("POST", url, nil, nil)
	if err != nil {
		log.Printf("http.Do(%v) error: %v\n", res.Request.URL, err)
		return 0, err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return res.StatusCode, nil
	case http.StatusInternalServerError:
		return res.StatusCode, errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return res.StatusCode, errors.New(fmt.Sprintf("Not found: (%s)", url))
	case http.StatusNotModified:
		return res.StatusCode, errors.New(fmt.Sprintf("Already started: (%s)", url))
	default:
		return res.StatusCode, errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

//noinspection GoNilness
func (p *Portainer) StopContainer(e uint, id string) (int, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/stop", e, id)
	res, err := p.makeRequest("POST", url, nil, nil)
	if err != nil {
		log.Printf("http.Do(%v) error: %v\n", res.Request.URL, err)
		return 0, err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return res.StatusCode, nil
	case http.StatusInternalServerError:
		return res.StatusCode, errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return res.StatusCode, errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return res.StatusCode, errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

//noinspection GoNilness
func (p *Portainer) RemoveContainer(e uint, id string) (int, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/remove", e, id)
	res, err := p.makeRequest("DELETE", url, nil, nil)
	if err != nil {
		log.Printf("http.Do(%v) error: %v\n", res.Request.URL, err)
		return 0, err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return res.StatusCode, nil
	case http.StatusInternalServerError:
		return res.StatusCode, errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return res.StatusCode, errors.New(fmt.Sprintf("Not found: (%s)", url))
	case http.StatusNotModified:
		return res.StatusCode, errors.New(fmt.Sprintf("Already started: (%s)", url))
	default:
		return res.StatusCode, errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

func (p *Portainer) makeRequest(tp string, url string, body io.Reader, args map[string]string) (*http.Response, error) {
	urlargs := "?"
	for k, v := range args {
		urlargs += fmt.Sprintf("%s=%s", k, v)
	}
	if urlargs == "?" {
		urlargs = ""
	}

	req, err := http.NewRequest(tp, p.ApiURL+url+urlargs, body)
	if err != nil {
		log.Printf("http.NewRequest() error: %v\n", err)
	}
	req.Header.Add("Authorization", "Bearer "+p.Token)
	c := &http.Client{Timeout: time.Second * 3}
	return c.Do(req)
}
