package purchase_order

import (
	"context"
	"db_lab_library/consts"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/cloudwego/hertz/pkg/app"
)

type InsertReq struct {
	OutOfStockRecordID int32 `gorm:"column:out_of_stock_record_id;not null" json:"out_of_stock_record_id"`
	Number             int32 `gorm:"column:number;not null" json:"number"`
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

	purchaseOrder := &model.PurchaseOrder{
		OutOfStockRecordID: req.OutOfStockRecordID,
		Number:             req.Number,
		Status:             consts.NotPurchased,
	}
	if err := query.PurchaseOrder.WithContext(ctx).Create(purchaseOrder); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, InsertResp{ID: purchaseOrder.ID})
}

type UpdateStatusReq struct {
	ID                 int32 `json:"id"`
	OutOfStockRecordID int32 `json:"out_of_stock_record_id"`
}

func UpdateStatus(ctx context.Context, c *app.RequestContext) {
	var req UpdateStatusReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	po := query.PurchaseOrder
	o := query.OutOfStockRecord
	b := query.Book

	results, err := po.WithContext(ctx).GetByCondition(req.ID, req.OutOfStockRecordID)
	if err != nil || len(results) == 0 {
		util.NewErrResp(c, err)
		return
	}
	purchaseOrder := results[0]

	outOfStockRecord, err := o.WithContext(ctx).Where(o.ID.Eq(purchaseOrder.OutOfStockRecordID)).First()
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	book, err := b.WithContext(ctx).Where(b.ID.Eq(outOfStockRecord.BookID)).First()
	if err != nil {
		util.NewErrResp(c, err)
		return
	}
	book.Inventory += purchaseOrder.Number
	if _, err := b.WithContext(ctx).Updates(book); err != nil {
		util.NewErrResp(c, err)
		return
	}

	if _, err := o.WithContext(ctx).Delete(outOfStockRecord); err != nil {
		util.NewErrResp(c, err)
		return
	}

	purchaseOrder.Status = consts.Purchased
	if _, err := po.Updates(purchaseOrder); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, nil)
}
