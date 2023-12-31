package out_of_stock_record

import (
	"context"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/cloudwego/hertz/pkg/app"
)

type InsertReq struct {
	BookID int32 `json:"book_id"`
	Number int32 `json:"number"`
}

type InsertResp struct {
	ID int32 `json:"id"`
}

func Insert(ctx context.Context, c *app.RequestContext) {
	var req InsertReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	outOfStockRecord := &model.OutOfStockRecord{
		BookID:     req.BookID,
		Number: req.Number,
	}
	if err := query.OutOfStockRecord.WithContext(ctx).Create(outOfStockRecord); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, InsertResp{ID: outOfStockRecord.ID})
}
