package book

import (
	"context"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/cloudwego/hertz/pkg/app"
)

type InsertReq struct {
	Name         string  `json:"name"`
	Authors      string  `json:"authors"`
	PublisherIds string  `json:"publisher_ids"`
	Price        float64 `json:"price"`
	Keywords     string  `json:"keywords"`
	Catalogue    string  `json:"catalogue"`
	Cover        string  `json:"cover"`
	Inventory    int32   `json:"inventory"`
	Location     string  `json:"location"`
	SubbookIds   string  `json:"subbook_ids"`
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

	book := &model.Book{
		Name:         req.Name,
		Authors:      req.Authors,
		Price:        req.Price,
		Keywords:     req.Keywords,
		Catalogue:    req.Catalogue,
		Cover:        req.Cover,
		Inventory:    req.Inventory,
		PublisherIds: req.PublisherIds,
		Location:     req.Location,
		SubbookIds:   req.SubbookIds,
	}
	if err := query.Book.WithContext(ctx).Create(book); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, "OK", InsertResp{ID: book.ID})
}

func Get(ctx context.Context, c *app.RequestContext) {
	
}
