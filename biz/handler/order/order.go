package order

import (
	"context"
	"db_lab_library/consts"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
	"time"
)

type InsertReq struct {
	UserID          int32   `json:"user_id"`
	BookIDs         []int32 `json:"book_ids"`
	OrderAmounts    []int32 `json:"order_amounts"`
	DeliveryAddress string  `json:"delivery_address"`
	PaymentAmount   float64 `json:"payment_amount"`
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

	var orderMoney float64
	b := query.Book
	for i, bookID := range req.BookIDs {
		book, err := b.WithContext(ctx).Where(b.ID.Eq(bookID)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.NewErrResp(c, errors.New("book not found"))
			continue
		}
		if book.Inventory < req.OrderAmounts[i] {
			util.NewErrResp(c, errors.New("inventory not enough"))
			oosr := query.OutOfStockRecord
			record, err := oosr.WithContext(ctx).Where(oosr.BookID.Eq(bookID)).First()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newRecord := &model.OutOfStockRecord{
					BookID: bookID,
					Number: req.OrderAmounts[i] - book.Inventory,
				}
				if err := oosr.WithContext(ctx).Create(newRecord); err != nil {
					util.NewErrResp(c, err)
					return
				}
			} else {
				record.Number += req.OrderAmounts[i] - book.Inventory
				if _, err := oosr.WithContext(ctx).Updates(record); err != nil {
					util.NewErrResp(c, err)
					return
				}
			}
			return
		}

		orderMoney += book.Price * float64(req.OrderAmounts[i])
		book.Inventory -= req.OrderAmounts[i]
		if _, err := b.WithContext(ctx).Updates(book); err != nil {
			util.NewErrResp(c, err)
			return
		}
	}
	u := query.User
	user, err := u.WithContext(ctx).Where(u.ID.Eq(req.UserID)).First()
	if err != nil {
		util.NewErrResp(c, err)
		return
	}
	payment := countPayment(user.CreditRating, orderMoney)
	if payment > req.PaymentAmount {
		util.NewErrResp(c, errors.New("payment not enough"))
		return
	}
	user.Balance -= payment
	if _, err := u.WithContext(ctx).Updates(user); err != nil {
		util.NewErrResp(c, err)
		return
	}

	bookIDsStr, err := sonic.MarshalString(req.BookIDs)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}
	amountsStr, err := sonic.MarshalString(req.OrderAmounts)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}
	order := &model.Order{
		UserID:          req.UserID,
		BookIds:         bookIDsStr,
		OrderAmounts:    amountsStr,
		OrderDate:       time.Now(),
		OrderMoney:      orderMoney,
		DeliveryAddress: req.DeliveryAddress,
		DeliveryStatus:  consts.NotDelivered,
	}
	if err := query.Order.WithContext(ctx).Create(order); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, InsertResp{ID: order.ID})
}

func countPayment(creditRating int32, orderMoney float64) (payment float64) {
	switch creditRating {
	case 1:
		payment = orderMoney * 0.9
	case 2:
		payment = orderMoney * 0.85
	case 3:
		payment = orderMoney * 0.85
	case 4:
		payment = orderMoney * 0.8
	case 5:
		payment = orderMoney * 0.75
	}
	return
}
