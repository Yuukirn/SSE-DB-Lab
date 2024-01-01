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

	util.NewOkResp(c, InsertResp{ID: user.ID})
}

type GetReq struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type GetResp struct {
	UserInfo   *model.User       `json:"user_info"`
	Orders     []*model.Order    `json:"orders"`
	Deliveries []*model.Delivery `json:"deliveries"`
}

func Get(ctx context.Context, c *app.RequestContext) {
	var req GetReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	users, err := query.User.WithContext(ctx).GetByCondition(req.ID, req.Name)
	if err != nil || len(users) == 0 {
		util.NewErrResp(c, err)
		return
	}

	user := users[0]
	orders, err := query.Order.WithContext(ctx).Where(query.Order.UserID.Eq(user.ID)).Find()
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	var deliveries []*model.Delivery
	d := query.Delivery
	for _, order := range orders {
		delivery, err := d.WithContext(ctx).Where(d.OrderID.Eq(order.ID)).First()
		if err != nil {
			util.NewErrResp(c, err)
			return
		}
		deliveries = append(deliveries, delivery)
	}

	util.NewOkResp(c, GetResp{
		UserInfo:   user,
		Orders:     orders,
		Deliveries: deliveries,
	})
}

type UpdateBasicInfoReq struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func UpdateBasicInfo(ctx context.Context, c *app.RequestContext) {
	var req UpdateBasicInfoReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	user := &model.User{
		Name:         req.Name,
		PasswordHash: util.Hash(req.Password),
		Address:      req.Address,
	}

	u := query.User
	info, err := u.WithContext(ctx).Where(u.ID.Eq(req.ID)).Updates(user)
	if err != nil || info.Error != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, nil)
}

type UpdateCreditRatingReq struct {
	ID           int32 `json:"id"`
	CreditRating int32 `json:"credit_rating"`
}

func UpdateCreditRating(ctx context.Context, c *app.RequestContext) {
	var req UpdateCreditRatingReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	user := &model.User{
		CreditRating: req.CreditRating,
	}

	u := query.User
	info, err := u.WithContext(ctx).Where(u.ID.Eq(req.ID)).Updates(&user)
	if err != nil || info.Error != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, nil)
}
