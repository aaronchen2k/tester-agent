package repo

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/casbin"
	"github.com/aaronchen2k/openstc/src/libs/common"
	db2 "github.com/aaronchen2k/openstc/src/libs/db"
	"github.com/aaronchen2k/openstc/src/models"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type BaseRepo struct {
}

func NewBaseRepo() *BaseRepo {
	return &BaseRepo{}
}

// GetAll 批量查询
func (r *BaseRepo) GetAll(model interface{}, s *models.Search) *gorm.DB {
	db := db2.Db.Model(model)
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", orderBy, sort))

	db.Scopes(r.FoundByWhere(s.Fields), r.Relation(s.Relations))

	return db
}

// Found 查询条件
func (r *BaseRepo) Found(s *models.Search) *gorm.DB {
	return db2.Db.Scopes(r.Relation(s.Relations), r.FoundByWhere(s.Fields))
}

// IsNotFound 判断是否是查询不存在错误
func (r *BaseRepo) IsNotFound(err error) bool {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
		color.Yellow("查询数据不存在")
		return true
	}
	return false
}

// Update 更新
func (r *BaseRepo) Update(v, d interface{}, id uint) error {
	if err := db2.Db.Model(v).Where("id = ?", id).Updates(d).Error; err != nil {
		color.Red(fmt.Sprintf("Update %+v to %+v\n", v, d))
		return err
	}
	return nil
}

// GetRolesForUser 获取角色
func (r *BaseRepo) GetRolesForUser(uid uint) []string {
	uids, err := casbinUtils.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}

// Relation 加载关联关系
func (r *BaseRepo) Relation(relates []*models.Relate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(relates) > 0 {
			for _, re := range relates {
				if len(re.Value) > 0 {
					if re.Func != nil {
						db = db.Preload(re.Value, re.Func)
					} else {
						db = db.Preload(re.Value)
					}
				}
				color.Yellow(fmt.Sprintf("Preoad %s", re))
			}
		}
		return db
	}
}

// FoundByWhere 查询条件
func (r *BaseRepo) FoundByWhere(fields []*models.Filed) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) > 0 {
			for _, field := range fields {
				if field != nil {
					if field.Condition == "" {
						field.Condition = "="
					}
					if value, ok := field.Value.(int); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(uint); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]int); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else {
						color.Red(fmt.Sprintf("未知数据类型：%+v", field.Value))
					}
				}
			}
		}
		return db
	}
}

// GetRelations 转换前端获取关联关系为 []*Relate
func (r *BaseRepo) GetRelations(relation string, fs map[string]interface{}) []*models.Relate {
	var relates []*models.Relate
	if len(relation) > 0 {
		arr := strings.Split(relation, ";")
		for _, item := range arr {
			relate := &models.Relate{
				Value: item,
			}
			// 增加关联过滤
			for key, f := range fs {
				if key == item {
					relate.Func = f
				}
			}
			relates = append(relates, relate)
		}

	}
	color.Yellow(fmt.Sprintf("relation :%s , relates:%+v", relation, relates))
	return relates
}

// GetSearch 转换前端查询关系为 *Filed
func (r *BaseRepo) GetSearch(key, search string) *models.Filed {
	if len(search) > 0 {
		if strings.Contains(search, ":") {
			searches := strings.Split(search, ":")
			if len(searches) == 2 {
				value := searches[0]
				if strings.ToLower(searches[1]) == "like" {
					value = fmt.Sprintf("%%%s%%", searches[0])
				}

				return &models.Filed{
					Condition: searches[1],
					Key:       key,
					Value:     value,
				}

			} else if len(searches) == 1 {
				return &models.Filed{
					Condition: "=",
					Key:       key,
					Value:     searches[0],
				}
			}
		} else {
			return &models.Filed{
				Condition: "=",
				Key:       key,
				Value:     search,
			}
		}
	}
	return nil
}

// Paginate 分页
func (r *BaseRepo) Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

// GetPermissionsForUser 获取角色权限
func (r *BaseRepo) GetPermissionsForUser(uid uint) [][]string {
	return casbinUtils.Enforcer.GetPermissionsForUser(strconv.FormatUint(uint64(uid), 10))
}

// DropTables 删除数据表
func (r *BaseRepo) DropTables() {
	_ = db2.Db.Migrator().DropTable(
		common.Config.DB.Prefix+"users",
		common.Config.DB.Prefix+"roles",
		common.Config.DB.Prefix+"permissions",
		common.Config.DB.Prefix+"articles",
		common.Config.DB.Prefix+"configs",
		common.Config.DB.Prefix+"tags",
		common.Config.DB.Prefix+"types",
		common.Config.DB.Prefix+"article_tags",
		"casbin_rule")
}

// Migrate 迁移数据表
func (r *BaseRepo) Migrate() {
	err := db2.Db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&gormadapter.CasbinRule{},
	)

	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}
}
