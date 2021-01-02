package repo

import (
	_const "github.com/aaronchen2k/openstc/internal/pkg/libs/const"
	"github.com/aaronchen2k/openstc/internal/server/model"
	"gorm.io/gorm"
)

func NewImageRepo() *ImageRepo {
	return &ImageRepo{}
}

type ImageRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *ImageRepo) Get(id int) (image model.Image) {
	r.DB.Where("id=?", id).First(&image)
	return
}

func (r *ImageRepo) QueryByOs(platform _const.OsPlatform, osType _const.OsType, osLang _const.OsLang) (ids []int) {
	var db = r.DB.Model(model.Image{}).Where("NOT disabled AND NOT deleted")
	if platform != "" {
		db.Where("osPlatform = ?", platform)
	}
	if osType != "" {
		db.Where("osType = ?", osType)
	}
	if platform != "" {
		db.Where("osLang = ?", osLang)
	}

	db.Order("id ASC, createdAt ASC").Find(&ids)

	return
}

func (r *ImageRepo) QueryByBrowser(browserType _const.BrowserType, version string) (ids []int) {
	sql := "SELECT r.backingImageId id " +
		"FROM BizBackingImageCapability_Relation r " +
		"LEFT JOIN BizBrowser browser ON browser.id = r.browserId " +
		"WHERE NOT browser.disabled AND NOT browser.deleted "
	if browserType != "" {
		sql += "AND browser.type = ? "
	}
	if version != "" {
		sql += "AND browser.version = ? "
	}

	sql += "ORDER BY id"
	r.DB.Raw(sql).Scan(&ids)

	return
}
