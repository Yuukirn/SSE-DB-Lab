package user

import (
	"context"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/cloudwego/hertz/pkg/app"
)

type InsertReq struct {
	Name         string  `json:"name"`
	Password     string  `json:"password"`
	Address      string  `json:"address"`
	Balance      float64 `json:"balance"`
	CreditRating int32   `json:"credit_rating"`
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

	user := &model.User{
		Name:         req.Name,
		PasswordHash: util.Hash(req.Password),
		Address:      req.Address,
		Balance:      req.Balance,
		CreditRating: req.CreditRating,
	}
	if err := query.User.WithContext(ctx).Create(user); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, "OK", InsertResp{ID: user.ID})
}

func Get(ctx context.Context, c *app.RequestContext) {
	// get by condition
}
