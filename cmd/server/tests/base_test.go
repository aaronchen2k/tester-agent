// +build test api access perm role user expire

package tests

import (
	"flag"
	server2 "github.com/aaronchen2k/tester/cmd/server/server"
	"github.com/aaronchen2k/tester/cmd/testerserver/tests/mock"
	"github.com/aaronchen2k/tester/internal/server/cfg"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/bxcodec/faker/v3"
	"github.com/iris-contrib/httpexpect/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"net/http"
	"os"
	"testing"
)

var (
	app   *iris.Application
	token string
)

//单元测试基境
func TestMain(m *testing.M) {
	flag.Parse()
	serverConf.Init("")
	s := server2.NewServer(nil) // 初始化app
	s.NewApp(nil, nil, nil, nil)
	app = s.App
	service.Run()

	exitCode := m.Run()

	model.DropTables() // 删除测试数据表，保持测试环境
	os.Exit(exitCode)
}

func getHttpexpect(t *testing.T) *httpexpect.Expect {
	return httptest.New(t, app, httptest.Configuration{Debug: true, URL: "http://app.irisadminapi.com/v1/admin/"})
}

// 单元测试 login 方法
func login(t *testing.T, Object interface{}, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	e.POST("login").WithJSON(Object).
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Code, Msg)

	return
}

// 单元测试 create 方法
func create(t *testing.T, url string, Object interface{}, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	ob := e.POST(url).WithHeader("Authorization", "Bearer "+GetOauthToken(e)).WithJSON(Object).
		Expect().Status(StatusCode).JSON().Object()
	ob.Value("code").Equal(Code)
	ob.Value("message").Equal(Msg)

	return
}

// 单元测试 update 方法
func update(t *testing.T, url string, Object interface{}, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	ob := e.PUT(url).WithHeader("Authorization", "Bearer "+GetOauthToken(e)).WithJSON(Object).
		Expect().Status(StatusCode).JSON().Object()
	ob.Value("code").Equal(Code)
	ob.Value("message").Equal(Msg)

	return
}

// 单元测试 getOne 方法
func getOne(t *testing.T, url string, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	e.GET(url).WithHeader("Authorization", "Bearer "+GetOauthToken(e)).
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Code, Msg)
	return
}

// 单元测试 getOnAuth 方法
func getOnAuth(t *testing.T, url string, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	e.GET(url).
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Code, Msg)
	return
}

// 单元测试 bImport 方法
func bImport(t *testing.T, url string, StatusCode int, Code int, Msg string, _ map[string]interface{}) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	e.POST(url).WithHeader("Authorization", "Bearer "+GetOauthToken(e)).
		WithMultipart().
		WithFile("file", "permissions.xlsx").
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Code, Msg)

	return
}

// 单元测试 getMore 方法
func getMore(t *testing.T, url string, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	e.GET(url).WithHeader("Authorization", "Bearer "+GetOauthToken(e)).
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Code, Msg)

	return
}

// 单元测试 delete 方法
func delete(t *testing.T, url string, StatusCode int, Code int, Msg string) (e *httpexpect.Expect) {
	e = getHttpexpect(t)
	e.DELETE(url).WithHeader("Authorization", "Bearer "+GetOauthToken(e)).
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Code, Msg)
	return
}

func CreatePermission() (*model.Permission, error) {
	m := mock.Permission{}
	err := faker.FakeData(&m)
	if err != nil {
		return nil, err
	}
	perm := &model.Permission{
		Name:        m.Name,
		DisplayName: m.DisplayName,
		Description: m.Description,
		Act:         m.Act,
	}
	err = perm.CreatePermission()
	if err != nil {
		return perm, err
	}

	return perm, nil

}

func CreateRole() (*model.Role, error) {
	m := mock.Role{}
	err := faker.FakeData(&m)
	if err != nil {
		return nil, err
	}
	role := &model.Role{
		Name:        m.Name,
		DisplayName: m.DisplayName,
		Description: m.Description,
	}
	err = role.CreateRole()
	if err != nil {
		return role, err
	}

	return role, nil
}

func CreateUser() (*model.User, error) {
	r, err := CreateRole()
	if err != nil {
		return nil, err
	}
	m := mock.User{}
	err = faker.FakeData(&m)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username: m.Username,
		Password: m.Password,
		Name:     m.Name,
		RoleIds:  []uint{r.ID},
	}
	err = user.CreateUser()
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetOauthToken(e *httpexpect.Expect) string {
	if len(token) > 0 {
		return token
	}

	oj := map[string]string{
		"username": serverConf.Config.Admin.UserName,
		"password": serverConf.Config.Admin.Pwd,
	}
	r := e.POST("login").WithJSON(oj).
		Expect().
		Status(http.StatusOK).JSON().Object()

	token = r.Value("data").Object().Value("token").String().Raw()

	return token
}
