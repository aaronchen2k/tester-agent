package service

import (
	"github.com/aaronchen2k/tester/internal/server/model"
	"github.com/aaronchen2k/tester/internal/server/repo"
)

type ClusterService struct {
	ClusterRepo *repo.ClusterRepo `inject:""`
}

func NewClusterService() *ClusterService {
	return &ClusterService{}
}

func (s *ClusterService) ListByType(tp string) (clusters []model.Cluster) {
	clusters, _ = s.ClusterRepo.QueryByType(tp)

	return
}

func (s *ClusterService) ListAll(keywords string, pageNo, pageSize int) (hosts []model.Cluster, total int64) {
	hosts, total, _ = s.ClusterRepo.Query(keywords, pageNo, pageSize)

	return
}
