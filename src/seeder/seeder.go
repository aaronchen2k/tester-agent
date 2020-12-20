package seeder

import (
	"fmt"
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/models"
	logger "github.com/sirupsen/logrus"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/azumads/faker"
	"github.com/jinzhu/configor"
	"gorm.io/gorm"
)

var Fake *faker.Faker

var Seeds = struct {
	Perms []struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayname"`
		Description string `json:"description"`
		Act         string `json:"act"`
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob(filepath.Join(common.CWD(), "seeder", "data", "*.yml"))
	if common.Config.Debug {
		fmt.Println(fmt.Sprintf("数据填充YML文件路径：%+v\n", filepaths))
	}
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		logger.Println(err)
	}
}

func Run() {
	AutoMigrates()

	fmt.Println(fmt.Sprintf("系统设置填充完成！！"))
	CreatePerms()
	fmt.Println(fmt.Sprintf("权限填充完成！！"))
	CreateAdminRole()
	fmt.Println(fmt.Sprintf("管理角色填充完成！！"))
	CreateAdminUser()
	fmt.Println(fmt.Sprintf("管理员填充完成！！"))
}

func AddPerm() {
	fmt.Println(fmt.Sprintf("开始填充权限！！"))
	CreatePerms()
	CreateAdminRole()
	CreateAdminUser()
	fmt.Println(fmt.Sprintf("权限填充完成！！"))
}

// CreatePerms 新建权限
func CreatePerms() {
	if common.Config.Debug {
		fmt.Println(fmt.Sprintf("填充权限：%+v\n", Seeds))
	}
	for _, m := range Seeds.Perms {
		s := &models.Search{
			Fields: []*models.Filed{
				{
					Key:       "name",
					Condition: "=",
					Value:     m.Name,
				}, {
					Key:       "act",
					Condition: "=",
					Value:     m.Act,
				},
			},
		}
		perm, err := models.GetPermission(s)
		if err == nil {
			if perm.ID == 0 {
				perm = &models.Permission{
					Model:       gorm.Model{CreatedAt: time.Now()},
					Name:        m.Name,
					DisplayName: m.DisplayName,
					Description: m.Description,
					Act:         m.Act,
				}
				if err := perm.CreatePermission(); err != nil {
					logger.Println(fmt.Sprintf("权限填充错误：%+v\n", err))
				}
			}
		}
	}
}

// CreateAdminRole 新建管理角色
func CreateAdminRole() {
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "name",
				Condition: "=",
				Value:     common.Config.Admin.RoleName,
			},
		},
	}
	role, err := models.GetRole(s)
	var permIds []uint
	ss := &models.Search{
		Limit:  -1,
		Offset: -1,
	}
	perms, _, err := models.GetAllPermissions(ss)
	if common.Config.Debug {
		if err != nil {
			fmt.Println(fmt.Sprintf("权限获取失败：%+v\n", err))
		}
	}

	for _, perm := range perms {
		permIds = append(permIds, perm.ID)
	}
	role.PermIds = permIds

	if err == nil {
		if role.ID == 0 {
			role = &models.Role{
				Name:        common.Config.Admin.RoleName,
				DisplayName: common.Config.Admin.RoleDisplayName,
				Description: common.Config.Admin.RoleDisplayName,
				Model:       gorm.Model{CreatedAt: time.Now()},
			}
			role.PermIds = permIds
			if err := role.CreateRole(); err != nil {
				logger.Println(fmt.Sprintf("管理角色填充错误：%+v\n", err))
			}
		} else {
			if err := models.UpdateRole(role.ID, role); err != nil {
				logger.Println(fmt.Sprintf("管理角色填充错误：%+v\n", err))
			}
		}
	}
	if common.Config.Debug {
		fmt.Println(fmt.Sprintf("填充角色数据：%+v\n", role))
		fmt.Println(fmt.Sprintf("填充角色权限：%+v\n", role.PermIds))
	}

}

// CreateAdminUser 新建管理员
func CreateAdminUser() {
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "username",
				Condition: "=",
				Value:     common.Config.Admin.UserName,
			},
		},
	}
	admin, err := models.GetUser(s)
	if err != nil {
		fmt.Println(fmt.Sprintf("Get admin error：%+v\n", err))
	}

	var roleIds []uint
	ss := &models.Search{
		Limit:  -1,
		Offset: -1,
	}
	roles, _, err := models.GetAllRoles(ss)
	if common.Config.Debug {
		if err != nil {
			fmt.Println(fmt.Sprintf("角色获取失败：%+v\n", err))
		}
	}

	for _, role := range roles {
		roleIds = append(roleIds, role.ID)
	}
	admin.RoleIds = roleIds

	if admin.ID == 0 {
		admin = &models.User{
			Username: common.Config.Admin.UserName,
			Name:     common.Config.Admin.Name,
			Password: common.Config.Admin.Password,
			Avatar:   "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIPbZRufW9zPiaGpfdXgU7icRL1licKEicYyOiace8QQsYVKvAgCrsJx1vggLAD2zJMeSXYcvMSkw9f4pw/132",
			Intro:    "超级弱鸡程序猿一枚！！！！",
			Model:    gorm.Model{CreatedAt: time.Now()},
		}
		admin.RoleIds = roleIds
		if err := admin.CreateUser(); err != nil {
			logger.Println(fmt.Sprintf("管理员填充错误：%+v\n", err))
		}
	} else {
		admin.Password = common.Config.Admin.Password
		if err := models.UpdateUserById(admin.ID, admin); err != nil {
			logger.Println(fmt.Sprintf("管理员填充错误：%+v\n", err))
		}
	}

	if common.Config.Debug {
		fmt.Println(fmt.Sprintf("管理员密码：%s\n", common.Config.Admin.Password))
		fmt.Println(fmt.Sprintf("填充管理员数据：%+v", admin))
	}
}

/*
	AutoMigrates 重置数据表
	libs.Db.DropTableIfExists 删除存在数据表
	libs.Db.AutoMigrate 重建数据表
*/
func AutoMigrates() {
	models.DropTables()
	models.Migrate()
}
