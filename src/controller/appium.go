package controller

import (
	"encoding/json"
	"github.com/aaronchen2k/openstc-common/src/domain"
	"github.com/aaronchen2k/openstc-common/src/libs/const"
	"github.com/aaronchen2k/openstc-common/src/libs/date"
	"github.com/aaronchen2k/openstc-common/src/libs/file"
	"github.com/aaronchen2k/openstc/src/service"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"mime/multipart"
	"time"
)

type AppiumController struct {
	Ctx     iris.Context
	Service *service.AppiumService `inject:""`
}

func NewAppiumController() *AppiumController {
	return &AppiumController{}
}
func (g *AppiumController) PostUpload() (result _domain.RpcResult) {
	dir := _const.UploadDir + _dateUtils.DateStr(time.Now())
	dir = _fileUtils.UpdateDir(dir)
	_fileUtils.MkDirIfNeeded(dir)

	g.Ctx.UploadFormFiles(dir, beforeResultSave)

	buildResultStr := g.Ctx.PostValue("result")
	buildResult := _domain.RpcResult{}
	json.Unmarshal([]byte(buildResultStr), &buildResult)

	_, info, _ := g.Ctx.FormFile("file")
	filePath := dir + info.Filename

	g.Service.SaveResult(buildResult, filePath)

	result.Success("success")
	return result
}

func beforeResultSave(context *context.Context, file *multipart.FileHeader) bool {
	uuid, _ := uuid.NewV4()
	file.Filename = "testResult-" + uuid.String() + ".zip"

	return true
}
