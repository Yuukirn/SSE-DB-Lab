package delivery

import (
	"context"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/cloudwego/hertz/pkg/app"
)

type InsertReq struct {
	OrderID int32 `json:"order_id"`
	UserID  int32 `json:"user_id"`
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

	delivery := &model.Delivery{
		OrderID: req.OrderID,
		UserID:  req.UserID,
	}

	if err := query.Delivery.WithContext(ctx).Create(delivery); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, InsertResp{ID: delivery.ID})
}
