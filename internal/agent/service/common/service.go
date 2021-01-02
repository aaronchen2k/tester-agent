package commonService

import (
	"github.com/aaronchen2k/openstc/internal/agent/cfg"
	_domain "github.com/aaronchen2k/openstc/internal/pkg/domain"
	_fileUtils "github.com/aaronchen2k/openstc/internal/pkg/libs/file"
	uuid "github.com/satori/go.uuid"
	"os"
)

func SetBuildWorkDir(build *_domain.BuildTo) {
	build.WorkDir = agentConf.Inst.WorkDir + uuid.NewV4().String() + string(os.PathSeparator)
	_fileUtils.MkDirIfNeeded(build.WorkDir)
}
