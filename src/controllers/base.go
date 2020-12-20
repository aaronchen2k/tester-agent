package controllers

import (
	"github.com/aaronchen2k/openstc/src/libs/common"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/kataras/iris/v12"
)

func GetCommonListSearch(ctx iris.Context) *models.Search {
	offset := common.ParseInt(ctx.FormValue("page"), 1)
	limit := common.ParseInt(ctx.FormValue("limit"), 20)
	orderBy := ctx.FormValue("orderBy")
	sort := ctx.FormValue("sort")

	relation := ctx.FormValue("relation")
	return &models.Search{
		Sort:      sort,
		Offset:    offset,
		Limit:     limit,
		OrderBy:   orderBy,
		Relations: models.GetRelations(relation, nil),
	}
}
