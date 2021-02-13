package handler

import (
	"encoding/json"
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_dateUtils "github.com/aaronchen2k/tester/internal/pkg/libs/date"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_stringUtils "github.com/aaronchen2k/tester/internal/pkg/libs/string"
	_utils "github.com/aaronchen2k/tester/internal/pkg/utils"
	"github.com/aaronchen2k/tester/internal/server/repo"
	"github.com/aaronchen2k/tester/internal/server/service"
	"github.com/kataras/iris/v12"
	"mime/multipart"
	"path"
	"time"
)

type BuildCtrl struct {
	Ctx          iris.Context
	BuildService *service.BuildService `inject:""`
	ResService   *service.ResService   `inject:""`

	BuildRepo *repo.BuildRepo `inject:""`
}

func NewBuildCtrl() *BuildCtrl {
	return &BuildCtrl{}
}

func (c *BuildCtrl) UpdateResult(ctx iris.Context) {
	dir := _fileUtils.AddSepIfNeeded(path.Join(_const.UploadDir, _dateUtils.DateStr(time.Now())))
	_fileUtils.MkDirIfNeeded(dir)

	ctx.UploadFormFiles(dir, c.BeforeResultSave)

	buildResultStr := ctx.PostValue("result")
	buildResult := _domain.RpcResult{}
	json.Unmarshal([]byte(buildResultStr), &buildResult)

	_, info, _ := ctx.FormFile("file")
	filePath := dir + info.Filename

	buildTo := c.BuildService.SaveResult(buildResult, filePath)
	c.ResService.DestroyByBuild(buildTo.ID)

	_, _ = ctx.JSON(_utils.ApiRes(200, "请求成功", nil))
}

func (c *BuildCtrl) BeforeResultSave(ctx iris.Context, file *multipart.FileHeader) bool {
	uuid := _stringUtils.NewUUID()
	file.Filename = "testResult-" + uuid + ".zip"

	return true
}
