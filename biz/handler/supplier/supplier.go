package supplier

import (
	"context"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetSupplyInfoReq struct {
	Id int32 `json:"id"`
}

type GetSupplyInfoResp struct {
	Books []*model.Book `json:"books"`
}

func GetSupplyInfo(ctx context.Context, c *app.RequestContext) {
	var req GetSupplyInfoReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	books, err := query.Book.WithContext(ctx).GetBySupplierID(req.Id)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, GetSupplyInfoResp{Books: books})
}

type GetSupplierInfoReq struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type GetSupplierInfoResp struct {
	Suppliers []*model.Supplier `json:"supplier"`
}

func GetSupplierInfo(ctx context.Context, c *app.RequestContext) {
	var req GetSupplierInfoReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	supplier, err := query.Supplier.WithContext(ctx).GetByCondition(req.ID, req.Name)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, GetSupplierInfoResp{Suppliers: supplier})
}
