// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"db_lab_library/mysql/model"
)

func newDelivery(db *gorm.DB, opts ...gen.DOOption) delivery {
	_delivery := delivery{}

	_delivery.deliveryDo.UseDB(db, opts...)
	_delivery.deliveryDo.UseModel(&model.Delivery{})

	tableName := _delivery.deliveryDo.TableName()
	_delivery.ALL = field.NewAsterisk(tableName)
	_delivery.ID = field.NewInt32(tableName, "id")
	_delivery.OrderID = field.NewInt32(tableName, "order_id")
	_delivery.UserID = field.NewInt32(tableName, "user_id")
	_delivery.CreatedAt = field.NewTime(tableName, "created_at")
	_delivery.UpdatedAt = field.NewTime(tableName, "updated_at")

	_delivery.fillFieldMap()

	return _delivery
}

type delivery struct {
	deliveryDo

	ALL       field.Asterisk
	ID        field.Int32
	OrderID   field.Int32
	UserID    field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (d delivery) Table(newTableName string) *delivery {
	d.deliveryDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d delivery) As(alias string) *delivery {
	d.deliveryDo.DO = *(d.deliveryDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *delivery) updateTableName(table string) *delivery {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt32(table, "id")
	d.OrderID = field.NewInt32(table, "order_id")
	d.UserID = field.NewInt32(table, "user_id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")

	d.fillFieldMap()

	return d
}

func (d *delivery) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *delivery) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 5)
	d.fieldMap["id"] = d.ID
	d.fieldMap["order_id"] = d.OrderID
	d.fieldMap["user_id"] = d.UserID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
}

func (d delivery) clone(db *gorm.DB) delivery {
	d.deliveryDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d delivery) replaceDB(db *gorm.DB) delivery {
	d.deliveryDo.ReplaceDB(db)
	return d
}

type deliveryDo struct{ gen.DO }

type IDeliveryDo interface {
	gen.SubQuery
	Debug() IDeliveryDo
	WithContext(ctx context.Context) IDeliveryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeliveryDo
	WriteDB() IDeliveryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeliveryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeliveryDo
	Not(conds ...gen.Condition) IDeliveryDo
	Or(conds ...gen.Condition) IDeliveryDo
	Select(conds ...field.Expr) IDeliveryDo
	Where(conds ...gen.Condition) IDeliveryDo
	Order(conds ...field.Expr) IDeliveryDo
	Distinct(cols ...field.Expr) IDeliveryDo
	Omit(cols ...field.Expr) IDeliveryDo
	Join(table schema.Tabler, on ...field.Expr) IDeliveryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeliveryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeliveryDo
	Group(cols ...field.Expr) IDeliveryDo
	Having(conds ...gen.Condition) IDeliveryDo
	Limit(limit int) IDeliveryDo
	Offset(offset int) IDeliveryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeliveryDo
	Unscoped() IDeliveryDo
	Create(values ...*model.Delivery) error
	CreateInBatches(values []*model.Delivery, batchSize int) error
	Save(values ...*model.Delivery) error
	First() (*model.Delivery, error)
	Take() (*model.Delivery, error)
	Last() (*model.Delivery, error)
	Find() ([]*model.Delivery, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Delivery, err error)
	FindInBatches(result *[]*model.Delivery, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Delivery) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeliveryDo
	Assign(attrs ...field.AssignExpr) IDeliveryDo
	Joins(fields ...field.RelationField) IDeliveryDo
	Preload(fields ...field.RelationField) IDeliveryDo
	FirstOrInit() (*model.Delivery, error)
	FirstOrCreate() (*model.Delivery, error)
	FindByPage(offset int, limit int) (result []*model.Delivery, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeliveryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deliveryDo) Debug() IDeliveryDo {
	return d.withDO(d.DO.Debug())
}

func (d deliveryDo) WithContext(ctx context.Context) IDeliveryDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deliveryDo) ReadDB() IDeliveryDo {
	return d.Clauses(dbresolver.Read)
}

func (d deliveryDo) WriteDB() IDeliveryDo {
	return d.Clauses(dbresolver.Write)
}

func (d deliveryDo) Session(config *gorm.Session) IDeliveryDo {
	return d.withDO(d.DO.Session(config))
}

func (d deliveryDo) Clauses(conds ...clause.Expression) IDeliveryDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deliveryDo) Returning(value interface{}, columns ...string) IDeliveryDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deliveryDo) Not(conds ...gen.Condition) IDeliveryDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deliveryDo) Or(conds ...gen.Condition) IDeliveryDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deliveryDo) Select(conds ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deliveryDo) Where(conds ...gen.Condition) IDeliveryDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deliveryDo) Order(conds ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deliveryDo) Distinct(cols ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deliveryDo) Omit(cols ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deliveryDo) Join(table schema.Tabler, on ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deliveryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deliveryDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deliveryDo) Group(cols ...field.Expr) IDeliveryDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deliveryDo) Having(conds ...gen.Condition) IDeliveryDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deliveryDo) Limit(limit int) IDeliveryDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deliveryDo) Offset(offset int) IDeliveryDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deliveryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeliveryDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deliveryDo) Unscoped() IDeliveryDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deliveryDo) Create(values ...*model.Delivery) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deliveryDo) CreateInBatches(values []*model.Delivery, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deliveryDo) Save(values ...*model.Delivery) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deliveryDo) First() (*model.Delivery, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Delivery), nil
	}
}

func (d deliveryDo) Take() (*model.Delivery, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Delivery), nil
	}
}

func (d deliveryDo) Last() (*model.Delivery, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Delivery), nil
	}
}

func (d deliveryDo) Find() ([]*model.Delivery, error) {
	result, err := d.DO.Find()
	return result.([]*model.Delivery), err
}

func (d deliveryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Delivery, err error) {
	buf := make([]*model.Delivery, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deliveryDo) FindInBatches(result *[]*model.Delivery, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deliveryDo) Attrs(attrs ...field.AssignExpr) IDeliveryDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deliveryDo) Assign(attrs ...field.AssignExpr) IDeliveryDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deliveryDo) Joins(fields ...field.RelationField) IDeliveryDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deliveryDo) Preload(fields ...field.RelationField) IDeliveryDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deliveryDo) FirstOrInit() (*model.Delivery, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Delivery), nil
	}
}

func (d deliveryDo) FirstOrCreate() (*model.Delivery, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Delivery), nil
	}
}

func (d deliveryDo) FindByPage(offset int, limit int) (result []*model.Delivery, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deliveryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deliveryDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deliveryDo) Delete(models ...*model.Delivery) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deliveryDo) withDO(do gen.Dao) *deliveryDo {
	d.DO = *do.(*gen.DO)
	return d
}