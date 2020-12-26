package controller

import (
	"github.com/aaronchen2k/openstc-common/src/libs/const"
	"github.com/aaronchen2k/openstc-common/src/libs/date"
	"github.com/aaronchen2k/openstc-common/src/libs/file"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"mime/multipart"
	"time"
)

type FileController struct {
	Ctx iris.Context
}

func NewFileController() *FileController {
	return &FileController{}
}
func (g *FileController) PostUpload() {
	dir := _const.UploadDir + _dateUtils.DateStr(time.Now())
	_fileUtils.MkDirIfNeeded(dir)

	g.Ctx.UploadFormFiles("./uploads", beforeFileSave)
}

func beforeFileSave(context *context.Context, file *multipart.FileHeader) bool {
	uuid, _ := uuid.NewV4()
	file.Filename = uuid.String() + "-" + file.Filename

	return true
}
