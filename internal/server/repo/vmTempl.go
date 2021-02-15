package repo

import (
	"fmt"
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/model/base"
	"gorm.io/gorm"
	"strings"
)

func NewVmTemplRepo() *VmTemplRepo {
	return &VmTemplRepo{}
}

type VmTemplRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *VmTemplRepo) Get(id uint) (templ model.VmTempl) {
	r.DB.Where("id=?", id).First(&templ)
	return
}
func (r *VmTemplRepo) GetByIdent(ident, computer, cluster string) (templ model.VmTempl) {
	r.DB.Model(&templ).Where("Ident=? AND computer=? AND cluster=?", ident, computer, cluster).First(&templ)
	return
}

func (r *VmTemplRepo) GetByEnv(env base.TestEnv) (templ model.VmTempl) {
	conditions := make([]string, 0)
	if templ.OsPlatform != "" {
		conditions = append(conditions, fmt.Sprintf("templ.os_platform=%s", templ.OsPlatform))
	}
	if templ.OsType != "" {
		conditions = append(conditions, fmt.Sprintf("templ.os_type=%s", templ.OsType))
	}
	if templ.OsLang != "" {
		conditions = append(conditions, fmt.Sprintf("templ.os_lang=%s", templ.OsLang))
	}
	if templ.OsVer != "" {
		conditions = append(conditions, fmt.Sprintf("templ.os_ver=%s", templ.OsVer))
	}
	if templ.OsBits != "" {
		conditions = append(conditions, fmt.Sprintf("templ.os_bits=%s", templ.OsBits))
	}

	sql := fmt.Sprintf(`SELECT templ.id, templ.name, templ.vm_id,
		computer.id computerId, computer.ident computerIdent, 
		computer.cluster computerCluster, computer.inst_count

	FROM biz_vm_templ templ
	LEFT JOIN biz_computer computer ON templ.computer = computer.ident

	WHERE %s
	ORDER BY computer.inst_count
	LIMIT 1`, strings.Join(conditions, "AND"))

	r.DB.Raw(sql).Scan(&templ)

	condition := r.convertEnvToVmTempl(env)
	r.DB.Where(&condition).First(&templ)
	return
}

func (r *VmTemplRepo) Create(templ *model.VmTempl) {
	r.DB.Model(&templ).Create(templ)
	return
}

func (r *VmTemplRepo) Update(templ *model.VmTempl) (err error) {
	r.DB.Model(&templ).Updates(templ)
	return
}
func (r *VmTemplRepo) UpdateAllSameName(templ *model.VmTempl) (err error) {
	templ.ID = 0
	r.DB.Model(&model.VmTempl{}).Where("name = ?", templ.Name).Updates(templ)
	return
}

func (r *VmTemplRepo) convertEnvToVmTempl(env base.TestEnv) (templ model.VmTempl) {
	if env.OsPlatform != "" {
		templ.OsPlatform = env.OsPlatform
	}
	if env.OsType != "" {
		templ.OsType = env.OsType
	}
	if env.OsLevel != "" {
		templ.OsLevel = env.OsLevel
	}
	if env.OsLang != "" {
		templ.OsLang = env.OsLang
	}

	if env.OsVer != "" {
		templ.OsVer = env.OsVer
	}
	if env.OsBuild != "" {
		templ.OsBuild = env.OsBuild
	}
	if env.OsBits != "" {
		templ.OsBits = env.OsBits
	}

	if env.BrowserType != "" {
		templ.BrowserType = env.BrowserType
	}
	if env.BrowserVer != "" {
		templ.BrowserVer = env.BrowserVer
	}
	if env.BrowserLang != "" {
		templ.BrowserLang = env.BrowserLang
	}

	return
}
