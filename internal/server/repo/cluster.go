package repo

import (
	_const "github.com/aaronchen2k/tester/internal/pkg/const"
	_commonUtils "github.com/aaronchen2k/tester/internal/pkg/libs/common"
	"github.com/aaronchen2k/tester/internal/server/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func NewClusterRepo() *ClusterRepo {
	return &ClusterRepo{}
}

type ClusterRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func (r *ClusterRepo) Query(keywords string, pageNo, pageSize int) (clusters []model.Cluster, total int64, err error) {
	query := r.DB.Select("*")
	if keywords != "" {
		query = query.Where("title LIKE ?", "%"+keywords+"%")
	}
	if pageNo > 0 {
		query = query.Offset((pageNo - 1) * pageSize).Limit(pageSize)
	}

	err = query.Find(&clusters).Error
	err = r.DB.Model(&model.Cluster{}).Count(&total).Error

	return
}

func (r *ClusterRepo) QueryByType(tp string) (clusters []model.Cluster, err error) {
	err = r.DB.Where("type=?", tp).Find(&clusters).Error

	return
}

func (r *ClusterRepo) Get(id uint) (cluster model.Cluster) {
	r.DB.Where("id=?", id).First(&cluster)
	return
}
func (r *ClusterRepo) GetByIdent(ident string) (cluster model.Cluster) {
	r.DB.Where("ident=?", ident).First(&cluster)
	return
}

func (r *ClusterRepo) QueryByImages(images []uint, ids []uint) (clusterId, backingImageId uint) {
	sql := `SELECT r.clusterId, r.backingImageId imageId 
			FROM BizClusterCapability_relation r 
		    INNER JOIN BizCluster cluster on r.clusterId = cluster.id 

	        WHERE cluster.status = 'active' 
			AND r.backingImageId IN (` +
		strings.Join(_commonUtils.UintToStrArr(images), ",") +
		`) AND cluster.id IN ("` +
		strings.Join(_commonUtils.UintToStrArr(ids), ",") +
		`) LIMIT 1`

	r.DB.Raw(sql).Scan(&ids)
	return
}

func (r *ClusterRepo) QueryIdle(cluster int) (ret []map[string]uint) {
	sql := `SELECT temp.clusterId, temp.vmCount 
			FROM (
				SELECT DISTINCT cluster.id clusterId, IFNULL(sub.num, 0) vmCount
				FROM BizCluster cluster
				LEFT JOIN
					(SELECT clusterId, COUNT(1) num
						FROM BizVm
						WHERE status != 'destroy' AND NOT deleted AND NOT disabled
						GROUP BY clusterId) sub
					ON cluster.id = sub.clusterId
			) temp
			WHERE temp.vmCount <= ` + strconv.Itoa(_const.MaxVmOnHost)

	r.DB.Raw(sql).Scan(&ret)
	return
}
