package book

import (
	"context"
	"db_lab_library/mysql/model"
	"db_lab_library/mysql/query"
	"db_lab_library/util"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
)

type InsertReq struct {
	Name        string   `json:"name"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	SupplierIds []int    `json:"supplier_ids"`
	Price       float64  `json:"price"`
	Keywords    []string `json:"keywords"`
	Catalogue   string   `json:"catalogue"`
	Cover       string   `json:"cover"`
	Inventory   int32    `json:"inventory"`
	Location    string   `json:"location"`
	SubbookIds  []int    `json:"subbook_ids"`
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

	authorsStr, err := sonic.MarshalString(req.Authors)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	keywordsStr, err := sonic.MarshalString(req.Keywords)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	subbookStr, err := sonic.MarshalString(req.SubbookIds)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	supplierStr, err := sonic.MarshalString(req.SupplierIds)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	book := &model.Book{
		Name:        req.Name,
		Authors:     authorsStr,
		Publisher:   req.Publisher,
		Price:       req.Price,
		Keywords:    keywordsStr,
		Catalogue:   req.Catalogue,
		Cover:       req.Cover,
		Inventory:   req.Inventory,
		SupplierIds: supplierStr,
		Location:    req.Location,
		SubbookIds:  subbookStr,
	}
	if err := query.Book.WithContext(ctx).Create(book); err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, InsertResp{ID: book.ID})
}

type GetReq struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name"`
	Publisher string   `json:"publisher"`
	Keyword   string   `json:"keyword"`
	Authors   []string `json:"authors"`
}

type GetResp struct {
	Books []*model.Book `json:"books"`
}

func Get(ctx context.Context, c *app.RequestContext) {
	var req GetReq
	if err := c.BindJSON(&req); err != nil {
		util.NewErrResp(c, err)
		return
	}

	books, err := query.Book.WithContext(ctx).GetByCondition(req.ID, req.Name, req.Publisher, req.Keyword, req.Authors)
	if err != nil {
		util.NewErrResp(c, err)
		return
	}

	util.NewOkResp(c, GetResp{Books: books})
}
